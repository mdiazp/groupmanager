package routes

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"

	"github.com/mdiazp/groupmanager/server/api"
	"github.com/mdiazp/groupmanager/server/api/controllers"
	"github.com/mdiazp/groupmanager/server/api/middlewares"
)

// Router ...
func Router(base api.Base) http.Handler {
	// Middlewares
	auth := middlewares.Auth(base)
	logger := middlewares.Logger(base)
	serveJSON := middlewares.ServeJSON(base)

	// Controllers
	accountCtr := controllers.NewAccountController(base)

	router := mux.NewRouter()
	router.Use(logger, serveJSON, auth)

	// AccountController routes
	router.HandleFunc("/login", accountCtr.Login).Methods(http.MethodGet)
	router.HandleFunc("/profile", accountCtr.Profile).Methods(http.MethodGet)
	router.HandleFunc("/logout", accountCtr.Logout).Methods(http.MethodGet)

	h := cors.AllowAll().Handler(router)

	return h
}
