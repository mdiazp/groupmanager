package groupaduser

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
	"github.com/mdiazp/gm/server/api/controllers/group"
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
	return "/group/{id}/adusers/{aduser}"
}

func (c *createController) GetMethods() []string {
	return []string{"PUT"}
}

// GetAccess ...
func (c *createController) GetAccess() controllers.Permission {
	return controllers.PermissionCreateGroupADUser
}

// ServeHTTP ...
func (c *createController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	group.CheckGroupAdminAccessControl(c, w, r)

	groupID := c.GetPInt(w, r, "id")
	aduser := c.GetPString(w, r, "aduser")

	provider := api.GetUsersProvider(c, api.UserProviderAD)
	if provider == nil {
		c.WE(w, fmt.Errorf("Unknowed Provider: %s", api.UserProviderAD), 500)
	}

	ur, e := provider.GetUserRecords(aduser)
	if e != nil {
		c.WE(w, fmt.Errorf("ADUser not found"), 404)
	}

	o := models.GroupADUser{
		ADUser:  ur.Username,
		ADName:  ur.Name,
		GroupID: (uint)(groupID),
	}

	c.Validate(w, o)

	e = c.DB().CreateGroupADUser(&o)
	if e != nil && strings.Contains(e.Error(),
		fmt.Sprintf(`"%s"`, models.ConstrainstGroupADUserUniqueIndex)) {
		c.WE(w, fmt.Errorf("ADUser already exists in group"), 400)
	}
	c.WE(w, e, 500)
	c.WR(w, 201, o)
}
