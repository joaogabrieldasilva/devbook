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
	{
		URI: "/posts",
		Method: http.MethodGet,
		Function: controllers.GetPosts,
		RequiresAuthentication: true,
	},
	{
		URI: "/posts/{postId}",
		Method: http.MethodGet,
		Function: controllers.GetPostById,
		RequiresAuthentication: true,
	},
	{
		URI: "/posts/{postId}",
		Method: http.MethodPut,
		Function: controllers.UpdatePost,
		RequiresAuthentication: true,
	},
	{
		URI: "/posts/{postId}",
		Method: http.MethodDelete,
		Function: controllers.DeletePost,
		RequiresAuthentication: true,
	},
	{
		URI: "/users/{userId}/posts",
		Method: http.MethodGet,
		Function: controllers.GetPostsByUser,
		RequiresAuthentication: true,
	},
	
}