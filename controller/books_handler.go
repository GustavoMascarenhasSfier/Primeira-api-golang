package controller

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"Api-Aula1-golang/responses"
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
