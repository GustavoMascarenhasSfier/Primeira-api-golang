package main

import (
	"Api-Aula1-golang/router"
	"log"
	"net/http"
)

func main() {
	r := router.New()
	const addr = ":8080"
	log.Printf("Servidor rodando em http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, r)) //mata nossa aplicacao caso haja algum erro, como a porta ja estar em uso
}
