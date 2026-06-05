package routes

import (
	"Api-Aula1-golang/controller"
	"net/http"
)

var loginRoutes = []Route{
	{
		URI:    "/login",
		Method: http.MethodPost,
		Func:   controller.Login,
		Auth:   false,
	},
}
