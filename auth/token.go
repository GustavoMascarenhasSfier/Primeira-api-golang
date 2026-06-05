package auth

import (
	"Api-Aula1-golang/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix() //tempo onde expira o token do usuario
	permissions["userId"] = userID

	//jwt e um dos metodos para assinaturas de tokens
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	println(token)
	return token.SignedString(config.SecretKey)

}

// validacao do recebimento do token
func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, retriveAuthKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid token")

}

// valida se o token esta dentro do objeto do corpo da requisicao
// extrai apenas o token para a validacao
func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func retriveAuthKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Token signing method unexpected! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
