package handler

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/smarthut/smarthut/model"
)

// ListDevices all devices
func ListDevices(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, model.ListDevices())
}

// GetDevice ...
func GetDevice(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "devicename")

	d, err := model.GetDevice(name)
	if err != nil {
		render.JSON(w, r, model.Message{Msg: err.Error()})
		log.Println(err)
		return
	}

	render.JSON(w, r, d)
}

// SetSocket ...
func SetSocket(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "devicename")
	d, err := model.GetDevice(name)

	if err != nil {
		render.JSON(w, r, model.Message{Msg: err.Error()})
		log.Println(err)
		return
	}

	resp, err := http.Post(d.Host+"/api/v1/socket", "application/json; charset=utf-8", r.Body)
	if err != nil {
		render.JSON(w, r, model.Message{Msg: err.Error()})
		log.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		render.JSON(w, r, model.Message{Msg: err.Error()})
		log.Println(err)
		return
	}

	render.JSON(w, r, string(body))
}
