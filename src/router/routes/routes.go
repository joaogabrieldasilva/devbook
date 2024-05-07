package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)


type Route struct {
	URI string
	Method string
	Function func(w http.ResponseWriter, r *http.Request)
	RequiresAuthentication bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, postsRoutes...)

	for _, route := range routes {

		if (route.RequiresAuthentication) {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}

	}

	return r
}