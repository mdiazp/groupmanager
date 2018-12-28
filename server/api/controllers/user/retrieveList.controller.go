package user

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
	return "/users"
}

func (c *retrieveListController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *retrieveListController) GetAccess() controllers.Permission {
	return controllers.PermissionRetrieveUser
}

// ServeHTTP ...
func (c *retrieveListController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f := readUserFilter(c.Base, w, r)
	ob := c.GetQOrderBy(w, r)
	p := c.GetQPaginator(w, r)

	var users []models.User
	e := c.DB().RetrieveUserList(&users, f, ob, p)

	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
	c.WR(w, 200, users)
}
