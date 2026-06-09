package controller

import (
	"Api-Aula1-golang/auth"
	"Api-Aula1-golang/models"
	"Api-Aula1-golang/persistency"
	"Api-Aula1-golang/repository"
	"Api-Aula1-golang/responses"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	b, _ := io.ReadAll(r.Body)
	query := strings.TrimSpace(string(b))

	googleBooksKey := r.URL.Query().Get("key")
	googleURL := "https://www.googleapis.com/books/v1/volumes?q=" + url.QueryEscape(query) + "&key=" + googleBooksKey

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	gReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, googleURL, nil)

	resp, err := http.DefaultClient.Do(gReq)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println("📨 RESPOSTA GOOGLE SIZE:", len(body))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var book models.Book
	if err = json.Unmarshal(body, &book); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	book.UserID = userID

	if err = book.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := persistency.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewBookRepo(db)
	book.ID, err = repo.Create(book)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, book)
}

func GetUserBooks(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	db, err := persistency.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewBookRepo(db)
	books, err := repo.FindByUser(userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if books == nil {
		books = []models.Book{}
	}

	responses.JSON(w, http.StatusOK, books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	bookID, err := strconv.ParseInt(params["bookID"], 10, 64)
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

	repo := repository.NewBookRepo(db)
	book, err := repo.FindByID(bookID, userID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if book.ID == 0 {
		responses.Err(w, http.StatusNotFound, fmt.Errorf("livro não encontrado"))
		return
	}

	responses.JSON(w, http.StatusOK, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	bookID, err := strconv.ParseInt(params["bookID"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var book models.Book
	if err = json.Unmarshal(body, &book); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	book.ID = bookID
	book.UserID = userID

	if err = book.Prepare(); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	db, err := persistency.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewBookRepo(db)
	if err = repo.Update(book); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.ExtractUserID(r)
	if err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	params := mux.Vars(r)
	bookID, err := strconv.ParseInt(params["bookID"], 10, 64)
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

	repo := repository.NewBookRepo(db)
	if err = repo.Delete(bookID, userID); err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
