package session

import (
	"net/http"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
)

// LogoutController ...
type LogoutController interface {
	controllers.BaseController
}

// NewLogoutController ...
func NewLogoutController(base api.Base) LogoutController {
	return &logoutController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type logoutController struct {
	api.Base
}

func (c *logoutController) GetRoute() string {
	return "/session"
}

func (c *logoutController) GetMethods() []string {
	return []string{"DELETE"}
}

// GetAccess ...
func (c *logoutController) GetAccess() controllers.Permission {
	return controllers.PermissionSession
}

// ServeHTTP ...
func (c *logoutController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.WR(w, 200, "Session has been closed")
}
