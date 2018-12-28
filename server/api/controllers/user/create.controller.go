package user

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
	return "/user"
}

func (c *createController) GetMethods() []string {
	return []string{"POST"}
}

// GetAccess ...
func (c *createController) GetAccess() controllers.Permission {
	return controllers.PermissionCreateUser
}

// ServeHTTP ...
func (c *createController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var u ToPost
	c.ReadJSON(w, r, &u)

	provider := c.GetUsersProvider((api.UserProvider)(u.Provider))
	if provider == nil {
		c.WE(w, fmt.Errorf("Unknowed Provider: %s", u.Provider), 404)
	}

	ur, e := provider.GetUserRecords(u.Username)
	c.WE(w, e, 404)

	user := models.User{
		ID:       0,
		Provider: u.Provider,
		Username: ur.Username,
		Name:     ur.Name,
		Rol:      (string)(controllers.RolUser),
		Enabled:  true,
	}
	c.Validate(w, user)

	e = c.DB().CreateUser(&user)
	if e != nil && strings.Contains(e.Error(), `"uix_system_user_username"`) {
		c.WE(w, fmt.Errorf("User with same username already exists"), 400)
	}
	c.WE(w, e, 500)
	c.WR(w, 200, user)
}

// ToPost ...
type ToPost struct {
	Username string
	Provider string
}
