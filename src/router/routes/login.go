package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute = Route{
	URI: "/auth/login",
	Method: http.MethodPost,
	Function: controllers.Login,
	RequiresAuthentication: false,
}