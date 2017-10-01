package handler

import (
	"fmt"
	"net/http"
	"strconv"

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
		render.JSON(w, r, Error{err.Error()})
		return
	}

	render.JSON(w, r, d)
}

// ListSockets ...
func ListSockets(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "devicename")

	d, err := model.GetDevice(name)
	if err != nil {
		render.JSON(w, r, Error{err.Error()})
		return
	}

	render.JSON(w, r, d.Sockets)
}

// GetSocket ...
func GetSocket(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "devicename")
	d, err := model.GetDevice(name)

	if err != nil {
		render.JSON(w, r, Error{err.Error()})
		return
	}

	if err != nil {
		render.JSON(w, r, Error{err.Error()})
		return
	}

	num := chi.URLParam(r, "num")
	i, err := strconv.Atoi(num)
	if err != nil {
		render.JSON(w, r, Error{err.Error()})
	}

	if i < 0 || i >= len(d.Sockets) {
		render.JSON(w, r, Error{fmt.Sprintf("smarthome: there is no socket#%d in %s", i, name)})
	}
	render.JSON(w, r, d.Sockets[i])
}
