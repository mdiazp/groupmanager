package user

import (
	"net/http"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
	dbhandlers "github.com/mdiazp/gm/server/db/handlers"
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
	return "/user/{id}"
}

func (c *updateController) GetMethods() []string {
	return []string{"PATCH"}
}

// GetAccess ...
func (c *updateController) GetAccess() controllers.Permission {
	return controllers.PermissionUpdateUser
}

// ServeHTTP ...
func (c *updateController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := (uint)(c.GetPInt(w, r, "id"))
	var eu EditUser
	c.ReadJSON(w, r, &eu)

	var user models.User
	e := c.DB().RetrieveUserByID(id, &user)
	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)

	user.Rol = (string)(eu.Rol)
	user.Enabled = eu.Enabled

	c.Validate(w, user)

	e = c.DB().UpdateUser(id, &user)
	c.WE(w, e, 500)
	c.WR(w, 200, user)
}

// EditUser ...
type EditUser struct {
	Rol     controllers.Rol
	Enabled bool
}
