package group

import (
	"net/http"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
	"github.com/mdiazp/gm/server/db/models"
)

// UpdateController ...
type UpdateController interface {
	controllers.BaseController
}

// NewUpdateController ...
func NewUpdateController(base api.Base) UpdateController {
	return &updateController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type updateController struct {
	api.Base
}

func (c *updateController) GetRoute() string {
	return "/group/{id}"
}

func (c *updateController) GetMethods() []string {
	return []string{"PATCH"}
}

// GetAccess ...
func (c *updateController) GetAccess() controllers.Permission {
	return controllers.PermissionUpdateGroup
}

// ServeHTTP ...
func (c *updateController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := (uint)(c.GetPInt(w, r, "id"))
	var group models.Group
	c.ReadJSON(w, r, &group)
	group.ID = id

	verificateGroupExistence(c, w, id)

	c.Validate(w, group)

	e := c.DB().UpdateGroup(id, &group)
	c.WE(w, e, 500)
	c.WR(w, 200, group)
}
