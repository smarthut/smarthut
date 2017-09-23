package main

import (
	"net/http"

	"github.com/leonidboykov/smarthut/router"
)

func main() {
	http.ListenAndServe(":8080", router.Initialize())
}
