package routes

import (
	"Api-Aula1-golang/controller"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Func:   controller.CreateUser,
		Auth:   false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Func:   controller.FetchUsers,
		Auth:   true,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodGet,
		Func:   controller.FetchUser,
		Auth:   true,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodPut,
		Func:   controller.UpdateUser,
		Auth:   true,
	},
	{
		URI:    "/users/{userID}",
		Method: http.MethodDelete,
		Func:   controller.DeleteUser,
		Auth:   true,
	},
}
