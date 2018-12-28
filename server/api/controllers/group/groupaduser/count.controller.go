package groupaduser

import (
	"net/http"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
	"github.com/mdiazp/gm/server/api/controllers/group"
	dbhandlers "github.com/mdiazp/gm/server/db/handlers"
)

// CountController ...
type CountController interface {
	controllers.BaseController
}

// NewCountController ...
func NewCountController(base api.Base) CountController {
	return &countController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type countController struct {
	api.Base
}

func (c *countController) GetRoute() string {
	return "/group/{id}/aduserscount"
}

func (c *countController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *countController) GetAccess() controllers.Permission {
	return controllers.PermissionRetrieveGroupADUser
}

// ServeHTTP ...
func (c *countController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	group.CheckGroupAdminAccessControl(c, w, r)

	f := readGroupADUserFilter(c.Base, w, r)

	count, e := c.DB().CountGroupADUsers(f)

	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
	c.WR(w, 200, count)
}
