package user

import (
	"net/http"

	dbhandlers "github.com/mdiazp/gm/server/db/handlers"
	"github.com/mdiazp/gm/server/db/models"

	"github.com/mdiazp/gm/server/api"
)

func readUserFilter(c api.Base, w http.ResponseWriter, r *http.Request) *dbhandlers.UserFilter {
	f := dbhandlers.UserFilter{}

	f.UsernamePrefix = c.GetQString(w, r, "usernamePrefix", false)
	f.Provider = c.GetQString(w, r, "provider", false)
	f.Rol = c.GetQString(w, r, "rol", false)
	f.NameSubstr = c.GetQString(w, r, "nameSubstr", false)
	f.Enabled = c.GetQBool(w, r, "enabled", false)
	f.GroupWhichAdmin = c.GetQInt(w, r, "groupWhichAdmin", false)
	return &f
}

func verificateUserExistence(c api.Base, w http.ResponseWriter, id uint) {
	e := c.DB().RetrieveUserByID(id, &models.User{})
	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
}
