package routes

import (
	"net/http"
	"sendit-api/api/controllers"
)

var AuthRoutes = []Route{
	Route{
		URI:     "/auth/login",
		Method:  http.MethodPost,
		Handler: controllers.LoginUser,
	},
	Route{
		URI:     "/auth/signup",
		Method:  http.MethodPost,
		Handler: controllers.RegisterUser,
	},
}
