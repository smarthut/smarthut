package main

import (
	"net/http"

	"github.com/smarthut/smarthut/router"
)

func main() {
	http.ListenAndServe(":8080", router.Initialize())
}
