package main

import (
	"Api-Aula1-golang/config"
	"Api-Aula1-golang/router"
	"log"
	"net/http"
)

func main() {

	config.LoadEnv()

	r := router.New()
	//const addr = ":8080"
	log.Printf("Servidor rodando em http://localhost%s", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, r)) //mata nossa aplicacao caso haja algum erro, como a porta ja estar em uso
}
