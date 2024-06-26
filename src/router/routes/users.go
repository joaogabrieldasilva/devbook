package routes

import (
	"api/src/controllers"
	"net/http"
)


var userRoutes = []Route{
	{
		URI: "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		Function: controllers.GetUsers,
		RequiresAuthentication: true,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodGet,
		Function: controllers.GetUserById,
		RequiresAuthentication: true,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		RequiresAuthentication: true,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		RequiresAuthentication: true,
	},
	{
		URI: "/users/{userId}/follow",
		Method: http.MethodPost,
		Function: controllers.FollowUser,
		RequiresAuthentication: true,
	},
	{
		URI: "/users/{userId}/unfollow",
		Method: http.MethodPost,
		Function: controllers.UnfollowUser,
		RequiresAuthentication: true,
	},
	{
		URI: "/users/{userId}/followers",
		Method: http.MethodGet,
		Function: controllers.GetFollowers,
		RequiresAuthentication: true,
	},
	{
		URI: "/users/{userId}/following",
		Method: http.MethodGet,
		Function: controllers.GetFollowing,
		RequiresAuthentication: true,
	},
	{
		URI: "/users/update-password",
		Method: http.MethodPost,
		Function: controllers.UpdatePassword,
		RequiresAuthentication: true,
	},
}