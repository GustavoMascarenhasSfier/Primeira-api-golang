package routes

import (
	"Api-Aula1-golang/controller"
	"net/http"
)

var booksRoutes = []Route{

	{
		URI:    "/books",
		Method: http.MethodGet,
		Func:   controller.HandleSearch,
	},
}
