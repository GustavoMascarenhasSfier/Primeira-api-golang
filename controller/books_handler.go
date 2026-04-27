package controller // controle / empacota e comunica

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// HandleSearch chama nossa API Google
func HandleSearch(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // corpo da requisição

	fmt.Println("🚀 CHEGOU NO HANDLER")

	// lê a string enviada
	// "_" usado para ignorar o erro na tratativa
	b, _ := io.ReadAll(r.Body)
	fmt.Println("📦 BODY RAW:", string(b))

	// faz tratamento da query
	query := strings.TrimSpace(string(b)) // remove espaços desnecessários
	fmt.Println("🔎 QUERY:", query)

	googleURL := "https://www.googleapis.com/books/v1/volumes?q=" + url.QueryEscape(query)
	fmt.Println("🌐 URL FINAL:", googleURL)

	// contexto com timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// prepara a requisição
	gReq, _ := http.NewRequestWithContext(ctx, http.MethodGet, googleURL, nil)

	resp, err := http.DefaultClient.Do(gReq)
	if err != nil {
		fmt.Println("❌ ERRO REQUEST GOOGLE:", err)
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer resp.Body.Close()

	fmt.Println("📡 STATUS GOOGLE:", resp.Status)

	// lê resposta da API
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("❌ ERRO READ BODY:", err)
		writeJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("📨 RESPOSTA GOOGLE SIZE:", len(body))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	_, _ = w.Write(body)
}

// writeJSON facilita o retorno de erros em JSON
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
