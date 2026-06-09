package routes

import (
	"Api-Aula1-golang/controller"
	"net/http"
)

var booksRoutes = []Route{
	{
		URI:    "/books/search",
		Method: http.MethodPost,
		Func:   controller.HandleSearch,
		Auth:   false,
	},
	{
		URI:    "/library",
		Method: http.MethodPost,
		Func:   controller.AddBook,
		Auth:   true,
	},
	{
		URI:    "/library",
		Method: http.MethodGet,
		Func:   controller.GetUserBooks,
		Auth:   true,
	},
	{
		URI:    "/library/{bookID}",
		Method: http.MethodGet,
		Func:   controller.GetBook,
		Auth:   true,
	},
	{
		URI:    "/library/{bookID}",
		Method: http.MethodPut,
		Func:   controller.UpdateBook,
		Auth:   true,
	},
	{
		URI:    "/library/{bookID}",
		Method: http.MethodDelete,
		Func:   controller.DeleteBook,
		Auth:   true,
	},
}
