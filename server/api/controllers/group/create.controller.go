package group

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
	"github.com/mdiazp/gm/server/db/models"
)

// CreateController ...
type CreateController interface {
	controllers.BaseController
}

// NewCreateController ...
func NewCreateController(base api.Base) CreateController {
	return &createController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type createController struct {
	api.Base
}

func (c *createController) GetRoute() string {
	return "/group"
}

func (c *createController) GetMethods() []string {
	return []string{"POST"}
}

// GetAccess ...
func (c *createController) GetAccess() controllers.Permission {
	return controllers.PermissionCreateGroup
}

// ServeHTTP ...
func (c *createController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var group models.Group
	c.ReadJSON(w, r, &group)
	group.ID = 0

	c.Validate(w, group)

	e := c.DB().CreateGroup(&group)
	if e != nil && strings.Contains(e.Error(), `"uix_system_group_name"`) {
		c.WE(w, fmt.Errorf("Group with same name already exists"), 400)
	}
	c.WE(w, e, 500)
	c.WR(w, 200, group)
}
