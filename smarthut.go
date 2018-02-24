package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/smarthut/smarthut/model"
	"github.com/smarthut/smarthut/router"
)

func main() {
	// .env file is used for the local development
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
	}

	model.InitializeDevices()
	http.ListenAndServe(":8080", router.New())
}
