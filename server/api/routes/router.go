package routes

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"

	"github.com/mdiazp/gm/server/api"
	"github.com/mdiazp/gm/server/api/controllers"
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
	}

	router := mux.NewRouter()

	router.
		PathPrefix("/swagger/").
		Handler(http.StripPrefix("/swagger/", http.FileServer(http.Dir(base.PublicFolderPath()))))

	router.Use(logger)

	for _, ctr := range ctrs {
		var h http.Handler = ctr
		if ctr.GetAccess() != "" {
			h = middlewares.CheckAccessControl(base, ctr)
			h = middlewares.MustAuth(base, h)
		}
		router.Handle(ctr.GetRoute(), h).Methods(ctr.GetMethods()...)
	}

	h := cors.AllowAll().Handler(router)

	return h
}
