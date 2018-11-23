package api

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
)

// ErrorResponse struct contains error
type ErrorResponse struct {
	Error string `json:"error,omitempty"`
}

func handleError(err error, w http.ResponseWriter, r *http.Request) {
	log.Println(err)
	render.JSON(w, r, ErrorResponse{Error: err.Error()})
}
