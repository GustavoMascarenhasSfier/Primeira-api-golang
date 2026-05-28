package controller

import (
	"Api-Aula1-golang/models"
	"Api-Aula1-golang/persistency"
	"Api-Aula1-golang/repository"
	"Api-Aula1-golang/responses"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var newUser models.User
	if err = json.Unmarshal(bodyRequest, &newUser); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	log.Println(newUser)

	if err = newUser.Prepare("create"); err != nil {
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
	id, err := repo.Create(newUser)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	newUser.ID = id
	responses.JSON(w, http.StatusCreated, newUser)
}

// Busca todos os usuários
func FetchUsers(w http.ResponseWriter, r *http.Request) {
	db, err := persistency.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepo(db)
	users, err := repo.FindAll()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// Busca um usuário por ID
func FetchUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseInt(params["userID"], 10, 64)
	if err != nil {
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
	user, err := repo.FindByID(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {}
func DeleteUser(w http.ResponseWriter, r *http.Request) {}
