package user

import (
	"net/http"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
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
	return "/userscount"
}

func (c *countController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *countController) GetAccess() controllers.Permission {
	return controllers.PermissionRetrieveUser
}

// ServeHTTP ...
func (c *countController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f := readUserFilter(c.Base, w, r)

	count, e := c.DB().CountUsers(f)

	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
	c.WR(w, 200, count)
}
