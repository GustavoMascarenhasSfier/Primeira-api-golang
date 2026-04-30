package controller

import (
	"Api-Aula1-golang/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Simulando um "banco de dados" em memória
var users []models.User
var nextID int64 = 1

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User // variável para armazenar 1 usuário

	// 1️- Decodificar o JSON recebido no body da requisição
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	// 2️- Validar e formatar os dados (regra de negócio do model)
	if err := user.Prepare("create"); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 3️- Gerar ID automático
	user.ID = nextID
	nextID++

	// 4️- Salvar o usuário no "banco" (slice em memória)
	users = append(users, user)

	// não retornar a senha
	user.Senha = ""

	// 5️- Retornar resposta em JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// READ USERS
func FetchUser(w http.ResponseWriter, r *http.Request) {

	// lista auxiliar sem senha
	var safeUsers []models.User

	// 1️- Percorrer todos os usuários
	for _, u := range users {
		u.Senha = "" // remover senha
		safeUsers = append(safeUsers, u)
	}

	// 2️- Retornar lista
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(safeUsers)
}

// UPDATE USER
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	// ✔️ pegar da URL /users/{userID}
	vars := mux.Vars(r)
	idParam := vars["userID"]

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	var updated models.User

	err = json.NewDecoder(r.Body).Decode(&updated)
	if err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	for i, u := range users {
		if u.ID == id {

			updated.ID = u.ID

			if err := updated.Prepare("update"); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			users[i] = updated

			updated.Senha = ""

			json.NewEncoder(w).Encode(updated)
			return
		}
	}

	http.Error(w, "usuário não encontrado", http.StatusNotFound)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idParam := vars["userID"]

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "usuário não encontrado", http.StatusNotFound)
}
