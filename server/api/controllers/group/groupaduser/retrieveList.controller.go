package groupaduser

import (
	"net/http"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
	"github.com/mdiazp/gm/server/api/controllers/group"
	dbhandlers "github.com/mdiazp/gm/server/db/handlers"
	"github.com/mdiazp/gm/server/db/models"
)

// RetrieveListController ...
type RetrieveListController interface {
	controllers.BaseController
}

// NewRetrieveListController ...
func NewRetrieveListController(base api.Base) RetrieveListController {
	return &retrieveListController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type retrieveListController struct {
	api.Base
}

func (c *retrieveListController) GetRoute() string {
	return "/group/{id}/adusers"
}

func (c *retrieveListController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *retrieveListController) GetAccess() controllers.Permission {
	return controllers.PermissionRetrieveGroupADUser
}

// ServeHTTP ...
func (c *retrieveListController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	group.CheckGroupAdminAccessControl(c, w, r)

	f := readGroupADUserFilter(c, w, r)
	ob := c.GetQOrderBy(w, r)
	p := c.GetQPaginator(w, r)

	l := make([]models.GroupADUser, 0)
	e := c.DB().RetrieveGroupADUserList(f, ob, p, &l)

	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
	c.WR(w, 200, l)
}
