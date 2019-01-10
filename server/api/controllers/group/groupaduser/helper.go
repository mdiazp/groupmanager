package groupaduser

import (
	"net/http"

	dbhandlers "github.com/mdiazp/gm/server/db/handlers"

	"github.com/mdiazp/gm/server/api"
)

func readGroupADUserFilter(c api.Base, w http.ResponseWriter, r *http.Request) *dbhandlers.GroupADUserFilter {

	groupID := c.GetPInt(w, r, "id")
	tmp := (uint)(groupID)
	f := &dbhandlers.GroupADUserFilter{
		GroupID:      &tmp,
		ADUserPrefix: c.GetQString(w, r, "adUserPrefix"),
	}
	return f
}

func verificateGroupADUserExistence(c api.Base, w http.ResponseWriter,
	groupID uint, aduser string) {
	e := c.DB().RetrieveGroupADUser(groupID, aduser)
	if e == dbhandlers.ErrRecordNotFound {
		c.WE(w, e, 404)
	}
	c.WE(w, e, 500)
}
