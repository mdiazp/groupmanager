package groupadmin

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
	return "/group/{id}/admins/{userID}"
}

func (c *deleteController) GetMethods() []string {
	return []string{"DELETE"}
}

// GetAccess ...
func (c *deleteController) GetAccess() controllers.Permission {
	return controllers.PermissionCreateGroupAdmin
}

// ServeHTTP ...
func (c *deleteController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	groupID := c.GetPInt(w, r, "id")
	userID := c.GetPInt(w, r, "userID")

	e := c.DB().DeleteGroupAdmin((uint)(groupID), (uint)(userID))

	c.WE(w, e, 500)
	c.WR(w, 200, "Admin was deleted")
}
