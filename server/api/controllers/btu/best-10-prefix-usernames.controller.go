package btu

import (
	"fmt"
	"net/http"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
)

// RetrieveBTUController ...
type RetrieveBTUController interface {
	controllers.BaseController
}

// NewRetrieveBTUController ...
func NewRetrieveBTUController(base api.Base) RetrieveBTUController {
	return &retrieveBTUController{
		Base: base,
	}
}

//////////////////////////////////////////////////////////////////////////////////////

type retrieveBTUController struct {
	api.Base
}

func (c *retrieveBTUController) GetRoute() string {
	return "/btu/{usernamePrefix}"
}

func (c *retrieveBTUController) GetMethods() []string {
	return []string{"GET"}
}

// GetAccess ...
func (c *retrieveBTUController) GetAccess() controllers.Permission {
	return controllers.PermissionRetrieveBTU
}

// ServeHTTP ...
func (c *retrieveBTUController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	prefix := c.GetPString(w, r, "usernamePrefix")

	provider := c.GetUsersProvider((api.UserProviderAD))
	if provider == nil {
		c.WE(w, fmt.Errorf("Unknowed Provider: %s", api.UserProviderAD), 404)
	}

	users, e := provider.GetFirst10BestUsernamePrefixMatchs(prefix)

	c.WE(w, e, 500)
	c.WR(w, 200, users)
}
