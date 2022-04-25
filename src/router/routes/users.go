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
}
