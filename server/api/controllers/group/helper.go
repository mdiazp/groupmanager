package group

import (
	"net/http"

	dbhandlers "github.com/mdiazp/gm/server/db/handlers"
	"github.com/mdiazp/gm/server/db/models"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
)

func readGroupFilter(c api.Base, w http.ResponseWriter, r *http.Request) *dbhandlers.GroupFilter {
	f := dbhandlers.GroupFilter{}
	f.NameSubstr = c.GetQString(w, r, "nameSubstr", false)
	f.Actived = c.GetQBool(w, r, "actived", false)
	f.ADUser = c.GetQString(w, r, "aduser", false)

	user := c.ContextReadAuthor(w, r)
	if (controllers.Rol)(user.Rol) == controllers.RolUser {
		f.AdminID = &user.ID
		tmp := true
		f.Actived = &tmp
	} else {
		tmp := c.GetQInt(w, r, "adminID", false)
		if tmp != nil {
			tmp2 := (uint)(*tmp)
			f.AdminID = &tmp2
		}

		f.Actived = c.GetQBool(w, r, "actived", false)
	}

	return &f
}

// CheckGroupAdminAccessControl ...
func CheckGroupAdminAccessControl(c api.Base, w http.ResponseWriter, r *http.Request) {
	user := c.ContextReadAuthor(w, r)
	userRol := (controllers.Rol)(user.Rol)
	if userRol != controllers.RolUser {
		return
	}

	groupID := (uint)(c.GetPInt(w, r, "id"))
	e := c.DB().RetrieveGroupAdmin(groupID, user.ID)
	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 403)
	}
	c.WE(w, e, 500)

	// Check active status
	var group models.Group
	e = c.DB().RetrieveGroupByID(groupID, &group)
	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)

	if !group.Actived {
		c.WE(w, dbhandlers.ErrRecordNotFound, 404)
	}
}

func verificateGroupExistence(c api.Base, w http.ResponseWriter, id uint) {
	e := c.DB().RetrieveGroupByID(id, &models.Group{})
	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
}
