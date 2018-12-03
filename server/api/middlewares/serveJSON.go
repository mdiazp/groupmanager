package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/mdiazp/groupmanager/server/api"
)

//ServeJSON ...
func ServeJSON(b api.Base) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					if err, ok := rec.(api.Error); ok {
						tmp := err.Error()
						b.ContextWriteResponse(
							r, &api.Response{Status: err.Status, Body: &tmp},
						)
					} else {
						tmp := http.StatusText(500)
						b.ContextWriteResponse(
							r, &api.Response{Status: 500, Body: &tmp},
						)
					}
				}

				response := b.ContextReadResponse(r)
				status := response.Status
				body, e := json.Marshal(response.Body)
				if e != nil {
					status = 500
					body = []byte(http.StatusText(500))
				}

				w.WriteHeader(status)
				w.Write(body)
			}()

			next.ServeHTTP(w, r)
		})
	}
}
