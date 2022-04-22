package routes

import (
	"net/http"
	"social-app/src/controllers"
)

var loginRoute = []Route{
	{
		Uri:            "/",
		Method:         http.MethodGet,
		Function:       controllers.LoadLoginScreen,
		IsAuthRequired: false,
	},
	{
		Uri:            "/login",
		Method:         http.MethodGet,
		Function:       controllers.LoadLoginScreen,
		IsAuthRequired: false,
	},
}
