package api

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/smarthut/smarthut/model"
)

func (api *API) listDevices(w http.ResponseWriter, r *http.Request) {
	var devices []model.Device
	if err := api.db.All(&devices); err != nil {
		handleError(err, w, r)
		return
	}
	render.JSON(w, r, devices)
}

func (api *API) createDevice(w http.ResponseWriter, r *http.Request) {
	var d model.Device
	if err := render.DecodeJSON(r.Body, &d); err != nil {
		handleError(err, w, r)
		return
	}
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
	if err := api.db.Save(&d); err != nil {
		handleError(err, w, r)
		return
	}
	render.JSON(w, r, struct{}{})
}

func (api *API) getDevice(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "device_name")

	var d model.Device
	if err := api.db.One("Name", name, &d); err != nil {
		handleError(err, w, r)
		return
	}

	render.JSON(w, r, d)
}

func (api *API) updateDevice(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "device_name")

	var d model.Device
	if err := api.db.One("Name", name, &d); err != nil {
		handleError(err, w, r)
		return
	}

	var params model.Device
	if err := render.DecodeJSON(r.Body, &params); err != nil {
		handleError(err, w, r)
		return
	}

	params.ID = d.ID
	params.UpdatedAt = time.Now()
	if err := api.db.Update(&params); err != nil {
		handleError(err, w, r)
		return
	}

	render.JSON(w, r, struct{}{})
}

func (api *API) deleteDevice(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "device_name")

	var d model.Device
	if err := api.db.One("Name", name, &d); err != nil {
		handleError(err, w, r)
		return
	}

	if err := api.db.DeleteStruct(&d); err != nil {
		handleError(err, w, r)
		return
	}

	render.JSON(w, r, struct{}{})
}

func (api *API) getSocket(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "device_name")
	var d model.Device
	if err := api.db.One("Name", name, &d); err != nil {
		handleError(err, w, r)
		return
	}

	resp, err := http.Get(d.Host + "/api/v1")
	if err != nil {
		handleError(err, w, r)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		handleError(err, w, r)
		return
	}

	w.Write(body)
}

func (api *API) setSocket(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "device_name")
	var d model.Device
	if err := api.db.One("Name", name, &d); err != nil {
		handleError(err, w, r)
		return
	}

	resp, err := http.Post(d.Host+"/api/v1/socket", "application/json; charset=utf-8", r.Body)
	if err != nil {
		handleError(err, w, r)
		return
	}
	resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		handleError(err, w, r)
		return
	}

	w.Write(body)
}
