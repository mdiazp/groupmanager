package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/mdiazp/gm/server/api"
)

//Logger ...
func Logger(b api.Base) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ww := &responseWriter{ResponseWriter: w}
			tBegin := time.Now()

			defer func() {
				tEnd := time.Now()
				username := "not logued"
				status := ww.status
				e := ""

				if rec := recover(); rec != nil {
					if err, ok := rec.(api.Error); ok {
						e = fmt.Sprintf("%s -> %s", err.Location, err.Error())
					} else if err, ok := rec.(error); ok {
						e = fmt.Sprintf("%s -> %s", api.WAI(1), err.Error())
						b.WR(w, 500, err.Error())
					} else {
						msg := "Panic returns unknowed value"
						e = fmt.Sprintf("%s -> %s", api.WAI(1), msg)
						b.WR(w, 500, msg)
					}
				}

				tms := tEnd.Sub(tBegin).Nanoseconds() / (1000 * 1000)

				author := b.ContextReadAuthor(w, r, false)
				if author != nil {
					username = author.Username
				}

				log := fmt.Sprintf("RemoteAddr: %s | user: %s | URL: %s | "+
					"Method: %s | User-Agent: %s | Duration: %d ms | Response Status: %d",
					r.RemoteAddr, username, r.URL.String(),
					r.Method, r.Header.Get("User-Agent"), tms, status)

				if e == "" {
					b.Logger().Println(log)
				} else {
					b.Logger().Printf(log+"\n\t%s", e)
				}
			}()

			next.ServeHTTP(ww, r)
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
