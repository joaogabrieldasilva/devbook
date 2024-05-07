package routes

import (
	"api/src/controllers"
	"net/http"
)


var postsRoutes = []Route{
	{
		URI: "/posts",
		Method: http.MethodPost,
		Function: controllers.CreatePost,
		RequiresAuthentication: true,
	},
}