package usergroup

import (
	"net/http"

	dbhandlers "github.com/mdiazp/gm/server/db/handlers"

	"github.com/mdiazp/gm/server/api"
)

func readGroupAdminFilter(c api.Base, w http.ResponseWriter, r *http.Request) *dbhandlers.GroupAdminFilter {
	userID := c.GetPInt(w, r, "id")
	tmp := (uint)(userID)
	f := &dbhandlers.GroupAdminFilter{
		UserID: &tmp,
	}
	return f
}
