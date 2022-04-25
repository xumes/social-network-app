package routes

import (
	"net/http"
	"social-app/src/controllers"
)

var userRoutes = []Route{
	{
		Uri:            "/signup",
		Method:         http.MethodGet,
		Function:       controllers.LoadSignUp,
		IsAuthRequired: false,
	},
	{
		Uri:            "/users",
		Method:         http.MethodPost,
		Function:       controllers.CreateUser,
		IsAuthRequired: false,
	},
}
