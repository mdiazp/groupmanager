package group

import (
	"net/http"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
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
	return "/groups"
}

func (c *retrieveListController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *retrieveListController) GetAccess() controllers.Permission {
	return controllers.PermissionRetrieveGroup
}

// ServeHTTP ...
func (c *retrieveListController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f := readGroupFilter(c.Base, w, r)
	ob := c.GetQOrderBy(w, r)
	p := c.GetQPaginator(w, r)

	var groups []models.Group
	e := c.DB().RetrieveGroupList(&groups, f, ob, p)

	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
	c.WR(w, 200, groups)
}
