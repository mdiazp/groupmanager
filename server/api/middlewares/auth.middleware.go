package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
	dbhandlers "github.com/mdiazp/gm/server/db/handlers"
	"github.com/mdiazp/gm/server/db/models"
)

// AuthHeader ...
const AuthHeader = "AuthToken"

// MustAuth ...
func MustAuth(base api.Base, next http.Handler) http.Handler {
	return &Auth{
		next: next,
		Base: base,
	}
}

// Auth ...
type Auth struct {
	next http.Handler
	api.Base
}

func (c *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get(AuthHeader)
	claims, e := c.JWTHandler().GetClaims(tokenString)
	c.WE(w, e, 401)

	if !claims.VerifyExpiresAt(time.Now().Unix(), false) {
		c.WE(w, fmt.Errorf("Token expired"), 401)
	}

	var user models.User
	if claims.Provider == (string)(api.UserProviderAPIClient) {
		user = models.User{
			ID:       0,
			Provider: claims.Provider,
			Username: claims.Username,
			Name:     claims.Username,
			Rol:      (string)(controllers.RolReadOnly),
			Enabled:  true,
		}
	} else if claims.Provider != (string)(api.UserProviderRoot) {
		e = c.DB().RetrieveUserByUsername(claims.Username, &user)
		if e != nil {
			if e == dbhandlers.ErrRecordNotFound {
				c.WE(w, fmt.Errorf("User Not Found"), 401)
			}
			c.WE(w, e, 500)
		}
	} else {
		user = models.User{
			ID:       0,
			Provider: claims.Provider,
			Username: claims.Username,
			Name:     "Root",
			Rol:      (string)(controllers.RolAdmin),
			Enabled:  true,
		}
	}

	// Check Enabled property
	if !user.Enabled {
		c.WE(w, fmt.Errorf("User is not enabled"), 401)
	}

	c.ContextWriteAuthor(r, &user)

	c.next.ServeHTTP(w, r)
}
