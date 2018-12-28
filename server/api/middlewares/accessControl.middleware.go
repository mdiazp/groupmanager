package middlewares

import (
	"fmt"
	"net/http"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
)

// CheckAccessControl ...
func CheckAccessControl(base api.Base, next controllers.BaseController) http.Handler {
	return &AccessControl{
		next: next,
		Base: base,
	}
}

// AccessControl ...
type AccessControl struct {
	next controllers.BaseController
	api.Base
}

func (c *AccessControl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	preq := c.next.GetAccess()
	user := c.ContextReadAuthor(w, r)
	ps := controllers.GetPermissions(controllers.Rol(user.Rol))
	ok := false
	for _, p := range ps {
		if p == preq {
			ok = true
			break
		}
	}
	if !ok {
		c.WE(w, fmt.Errorf(http.StatusText(403)), 403)
	}
	c.next.ServeHTTP(w, r)
}
