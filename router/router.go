package router

import (
	"Api-Aula1-golang/router/routes"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	rotas := mux.NewRouter()
	routes.Register(rotas)
	return rotas
}
