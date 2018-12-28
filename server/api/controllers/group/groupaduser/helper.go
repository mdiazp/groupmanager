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
		GroupID: &tmp,
	}
	return f
}
