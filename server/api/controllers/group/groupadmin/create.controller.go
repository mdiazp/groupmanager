package groupadmin

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
	return "/group/{id}/admins/{userID}"
}

func (c *createController) GetMethods() []string {
	return []string{"PUT"}
}

// GetAccess ...
func (c *createController) GetAccess() controllers.Permission {
	return controllers.PermissionCreateGroupAdmin
}

// ServeHTTP ...
func (c *createController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	groupID := c.GetPInt(w, r, "id")
	userID := c.GetPInt(w, r, "userID")

	o := models.GroupAdmin{
		GroupID: (uint)(groupID),
		UserID:  (uint)(userID),
	}

	c.Validate(w, o)

	e := c.DB().CreateGroupAdmin(&o)
	if e != nil && strings.Contains(e.Error(), `"idx_system_user_group"`) {
		c.WE(w, fmt.Errorf("Admin already exists"), 400)
	}
	c.WE(w, e, 500)
	c.WR(w, 201, o)
}
