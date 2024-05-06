package routes

import (
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

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}