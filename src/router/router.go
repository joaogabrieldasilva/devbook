package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate will return a router
func Generate() *mux.Router {

	return routes.Configure(mux.NewRouter())
}