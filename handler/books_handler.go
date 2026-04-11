package handler

import (
	"fmt"
	"net/http"
)

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Search funcionando 🚀")
}
