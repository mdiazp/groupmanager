package group

import (
	"net/http"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
)

// DeleteController ...
type DeleteController interface {
	controllers.BaseController
}

// NewDeleteController ...
func NewDeleteController(base api.Base) DeleteController {
	return &deleteController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type deleteController struct {
	api.Base
}

func (c *deleteController) GetRoute() string {
	return "/group/{id}"
}

func (c *deleteController) GetMethods() []string {
	return []string{"DELETE"}
}

// GetAccess ...
func (c *deleteController) GetAccess() controllers.Permission {
	return controllers.PermissionDeleteGroup
}

// ServeHTTP ...
func (c *deleteController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := (uint)(c.GetPInt(w, r, "id"))

	verificateGroupExistence(c, w, id)

	e := c.DB().DeleteGroup(id)
	c.WE(w, e, 500)
	c.WR(w, 204, "Group was deleted successfully")
}
