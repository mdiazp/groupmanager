package routes

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
	"github.com/mdiazp/gm/server/api/controllers/btu"
	"github.com/mdiazp/gm/server/api/controllers/group"
	"github.com/mdiazp/gm/server/api/controllers/group/groupadmin"
	"github.com/mdiazp/gm/server/api/controllers/group/groupaduser"
	"github.com/mdiazp/gm/server/api/controllers/session"
	"github.com/mdiazp/gm/server/api/controllers/user"
	"github.com/mdiazp/gm/server/api/controllers/user/usergroup"
	"github.com/mdiazp/gm/server/api/middlewares"
)

// Router ...
func Router(base api.Base) http.Handler {
	// Middlewares
	logger := middlewares.Logger(base)

	ctrs := []controllers.BaseController{
		session.NewLoginController(base),
		session.NewLogoutController(base),
		session.NewProvidersController(base),

		user.NewCreateController(base),
		user.NewRetrieveController(base),
		user.NewRetrieveListController(base),
		user.NewCountController(base),
		user.NewUpdateController(base),
		user.NewDeleteController(base),

		usergroup.NewRetrieveListController(base),
		usergroup.NewCountController(base),

		group.NewCreateController(base),
		group.NewRetrieveController(base),
		group.NewRetrieveListController(base),
		group.NewCountController(base),
		group.NewUpdateController(base),
		group.NewDeleteController(base),

		groupadmin.NewCreateController(base),
		groupadmin.NewRetrieveListController(base),
		groupadmin.NewCountController(base),
		groupadmin.NewDeleteController(base),

		groupaduser.NewCreateController(base),
		groupaduser.NewRetrieveListController(base),
		groupaduser.NewCountController(base),
		groupaduser.NewDeleteController(base),

		btu.NewRetrieveBTUController(base),
	}

	router := mux.NewRouter()

	/*
		router.
			PathPrefix("/swagger/").
			Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir(base.PublicFolderPath()))))
	*/

	router.
		Host("swagger.api." + base.GetHost()).
		Handler(http.FileServer(http.Dir(base.PublicFolderPath() + "/swagger")))

	router.
		Host(base.GetHost()).
		Handler(http.FileServer(http.Dir(base.PublicFolderPath() + "/dist")))

		/*
			router.
				Host("swagger." + base.GetHost()).
				Handler(http.FileServer(http.Dir(base.PublicFolderPath())))
		*/
	/*
		PathPrefix("/swagger/").
		Handler(
			http.StripPrefix(
				"/swagger/",
				http.FileServer(http.Dir(base.PublicFolderPath())),
			),
		)
	*/
	// router.Use(logger)
	api := router.Host("api." + base.GetHost()).Subrouter()
	api.Use(logger)
	for _, ctr := range ctrs {
		var h http.Handler = ctr
		if ctr.GetAccess() != "" {
			h = middlewares.CheckAccessControl(base, ctr)
			h = middlewares.MustAuth(base, h)
		}
		api.Handle(ctr.GetRoute(), h).Methods(ctr.GetMethods()...)
	}

	h := cors.AllowAll().Handler(router)

	return h
}
