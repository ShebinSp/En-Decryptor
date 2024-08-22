package routes

import (
	"github.com/ShebnSp/endecrypror/internal/handlers"
	"net/http"
)


func RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/encrypt", handlers.EncodeHandler)
	mux.HandleFunc("/decrypt", handlers.DecodeHandler)

	return mux
}