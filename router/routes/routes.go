package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
}

func Register(r *mux.Router) {
	//r.HandleFunc("/books/search", handler.HandleSearch).Methods(http.MethodGet)

	var routes []Route
	routes = append(routes, booksRoutes...)
	routes = append(routes, userRoutes...)

	//remontar essa lista

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}

}
