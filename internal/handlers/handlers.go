package handlers

import (
	"net/http"

	"github.com/ShebnSp/endecrypror/internal/services"
)

// type handler struct {}

// func(h *handler) ServeHTTP (w http.ResponseWriter, r *http.Client) error {

// 	return nil
// }

func EncodeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {	
	case "POST":
		services.EncodeImage(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		services.DecodeImage(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

