package routes

import (
	"Api-Aula1-golang/handler"
	"net/http"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) {
	r.HandleFunc("/books/search", handler.HandleSearch).Methods(http.MethodGet)
}
