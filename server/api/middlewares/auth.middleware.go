package middlewares

import (
	"net/http"

	"github.com/mdiazp/groupmanager/server/api"
)

// Auth ...
func Auth(base api.Base, exceptions ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}
