package main

import (
	"fmt"
	"log"

	"github.com/smarthut/smarthut/api"
	"github.com/smarthut/smarthut/conf"
	"github.com/smarthut/smarthut/store"
)

var (
	version = "master"
	commit  = "none"
	date    = "unknown"
)

func main() {
	config, err := conf.Load("")
	if err != nil {
		fmt.Println(err)
	}

	storage, err := store.NewStore("data/")
	if err != nil {
		fmt.Println(err)
	}
	defer storage.Close()

	api := api.NewAPI(config, storage, version)

	l := fmt.Sprintf("%s:%d", config.API.Host, config.API.Port)
	log.Printf("Starting SmartHut Service %s on %s\n", version, l)
	api.Start(l)

	// model.InitializeDevices()
	// http.ListenAndServe(":8080", router.New())
}
