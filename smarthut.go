package main

import (
	"net/http"

	"github.com/smarthut/smarthut/model"
	"github.com/smarthut/smarthut/router"
)

func main() {
	model.InitializeDevices()
	http.ListenAndServe(":8080", router.New())
}
