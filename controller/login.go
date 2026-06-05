package controller

import (
	"Api-Aula1-golang/auth"
	"Api-Aula1-golang/models"
	"Api-Aula1-golang/persistency"
	"Api-Aula1-golang/repository"
	"Api-Aula1-golang/responses"
	"Api-Aula1-golang/security"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err := json.Unmarshal(reqBody, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := persistency.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepo(db)
	userSalvoemDB, err := repo.FetchByEmail(user.Email) // minúsculo
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.ValidatePassword(userSalvoemDB.Senha, user.Senha); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.GenerateToken(uint64(userSalvoemDB.ID)) // auth importado + cast
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	w.Write([]byte(token))
}
