package user

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
	return "/user/{id}"
}

func (c *deleteController) GetMethods() []string {
	return []string{"DELETE"}
}

// GetAccess ...
func (c *deleteController) GetAccess() controllers.Permission {
	return controllers.PermissionDeleteUser
}

// ServeHTTP ...
func (c *deleteController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := (uint)(c.GetPInt(w, r, "id"))

	e := c.DB().DeleteUser(id)
	c.WE(w, e, 500)
	c.WR(w, 204, "User was deleted successfully")
}
