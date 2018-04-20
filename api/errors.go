package api

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
)

// Message struct contains error
type Message struct {
	Error error `json:"error"`
}

func handleError(err error, w http.ResponseWriter, r *http.Request) {
	log.Println(err)
	render.JSON(w, r, Message{Error: err})
}
