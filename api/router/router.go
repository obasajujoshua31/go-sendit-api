package router

import (
	"sendit-api/api/router/routes"

	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) *mux.Router {
	r.StrictSlash(true)
	routes := routes.LoadRoutes()

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Handler).Methods(route.Method)
	}
	return r
}
