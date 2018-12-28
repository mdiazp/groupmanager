package groupadmin

import (
	"net/http"

	dbhandlers "github.com/mdiazp/gm/server/db/handlers"

	"github.com/mdiazp/gm/server/api"
)

func readGroupAdminFilter(c api.Base, w http.ResponseWriter, r *http.Request) *dbhandlers.GroupAdminFilter {
	groupID := c.GetPInt(w, r, "id")
	tmp := (uint)(groupID)
	f := &dbhandlers.GroupAdminFilter{
		GroupID: &tmp,
	}
	return f
}