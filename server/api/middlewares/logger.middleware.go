package middlewares

import (
	"net/http"
	"time"

	"github.com/mdiazp/groupmanager/server/api"
)

//Logger ...
func Logger(b api.Base) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := &responseWriter{ResponseWriter: w}

			tBegin := time.Now()
			next.ServeHTTP(ww, r)
			tEnd := time.Now()

			username := "not logued"
			status := ww.status

			tms := tEnd.Sub(tBegin).Nanoseconds() / 1000000

			author := b.ContextReadAuthor(r)
			if author != nil {
				username = author.Username
			}

			b.Logger().Printf("%s | %s | %s | %s | %s | %d ms | %d \n",
				r.RemoteAddr, username, r.URL.String(), r.Method, r.Header.Get("User-Agent"), tms, status)
		})
	}
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (w *responseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}
