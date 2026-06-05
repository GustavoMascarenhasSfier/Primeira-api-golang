package main

import (
	"Api-Aula1-golang/config"
	"Api-Aula1-golang/router"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
)

func init() {
	chave := make([]byte, 64)

	if _, err := rand.Read(chave); err != nil {
		log.Fatal(err)
	}
	stringBase64 := base64.StdEncoding.EncodeToString(chave)
	fmt.Println(stringBase64)
}

func main() {

	config.LoadEnv()

	r := router.New()
	//const addr = ":8080"
	log.Printf("Servidor rodando em http://localhost%s", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, r)) //mata nossa aplicacao caso haja algum erro, como a porta ja estar em uso
}
