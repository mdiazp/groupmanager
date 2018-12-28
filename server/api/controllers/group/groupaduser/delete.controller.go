package groupaduser

import (
	"net/http"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
	"github.com/mdiazp/gm/server/api/controllers/group"
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
	return "/group/{id}/adusers/{aduser}"
}

func (c *deleteController) GetMethods() []string {
	return []string{"DELETE"}
}

// GetAccess ...
func (c *deleteController) GetAccess() controllers.Permission {
	return controllers.PermissionCreateGroupADUser
}

// ServeHTTP ...
func (c *deleteController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	group.CheckGroupAdminAccessControl(c, w, r)

	groupID := c.GetPInt(w, r, "id")
	aduser := c.GetPString(w, r, "aduser")

	e := c.DB().DeleteGroupADUser((uint)(groupID), aduser)

	c.WE(w, e, 500)
	c.WR(w, 201, "ADUser was deleted")
}
