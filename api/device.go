package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/smarthut/smarthut/model"
)

func (api *API) listDevices(w http.ResponseWriter, r *http.Request) {
	var devices []model.Device
	if err := api.db.All(&devices); err != nil {
		log.Println(err)
		render.JSON(w, r, model.Message{Error: err.Error()})
		return
	}
	render.JSON(w, r, devices)
}

func (api *API) createDevice(w http.ResponseWriter, r *http.Request) {
	var d model.Device
	if err := render.DecodeJSON(r.Body, &d); err != nil {
		log.Println(err)
		return
	}
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
	if err := api.db.Save(&d); err != nil {
		log.Println(err)
		render.JSON(w, r, model.Message{Error: err.Error()})
		return
	}
	render.JSON(w, r, struct{}{})
}

func (api *API) getDevice(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "device_name")

	var d model.Device
	if err := api.db.One("Name", name, &d); err != nil {
		log.Println(err)
		render.JSON(w, r, model.Message{Error: err.Error()})
		return
	}

	render.JSON(w, r, d)
}

func (api *API) updateDevice(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "device_name")

	var d model.Device
	if err := api.db.One("Name", name, &d); err != nil {
		log.Println(err)
		render.JSON(w, r, model.Message{Error: err.Error()})
		return
	}

	var d2 model.Device
	if err := render.DecodeJSON(r.Body, &d2); err != nil {
		log.Println(err)
		render.JSON(w, r, model.Message{Error: err.Error()})
		return
	}
	d2.ID = d.ID
	d2.UpdatedAt = time.Now()

	if err := api.db.Update(&d2); err != nil {
		log.Println(err)
		render.JSON(w, r, model.Message{Error: err.Error()})
		return
	}

	render.JSON(w, r, struct{}{})
}

func (api *API) deleteDevice(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "device_name")

	var d model.Device
	if err := api.db.One("Name", name, &d); err != nil {
		log.Println(err)
		render.JSON(w, r, model.Message{Error: err.Error()})
		return
	}

	if err := api.db.DeleteStruct(&d); err != nil {
		log.Println(err)
		render.JSON(w, r, model.Message{Error: err.Error()})
		return
	}

	render.JSON(w, r, struct{}{})
}

func (api *API) getSocket(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "device_name")
	var d model.Device
	if err := api.db.One("Name", name, &d); err != nil {
		log.Println(err)
		render.JSON(w, r, model.Message{Error: err.Error()})
		return
	}

	resp, err := http.Get(d.Host + "/api/v1")
	if err != nil {
		log.Println(err)
		render.JSON(w, r, model.Message{Error: err.Error()})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		render.JSON(w, r, model.Message{Error: err.Error()})
		return
	}

	w.Write(body)
}

func (api *API) setSocket(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "device_name")
	var d model.Device
	if err := api.db.One("Name", name, &d); err != nil {
		log.Println(err)
		render.JSON(w, r, model.Message{Error: err.Error()})
		return
	}

	resp, err := http.Post(d.Host+"/api/v1/socket", "application/json; charset=utf-8", r.Body)
	if err != nil {
		log.Println(err)
		render.JSON(w, r, model.Message{Error: err.Error()})
		return
	}
	resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		render.JSON(w, r, model.Message{Error: err.Error()})
		return
	}

	w.Write(body)
}
