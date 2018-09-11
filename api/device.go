package api

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/pkg/errors"

	"github.com/smarthut/smarthut/model"
)

// DeviceResponse holds device data
type DeviceResponse struct {
	*model.Device
}

func (api *API) deviceCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		device := new(model.Device)
		if deviceName := chi.URLParam(r, "device_name"); deviceName != "" {
			if err := api.db.One("Name", deviceName, device); err != nil {
				handleError(errors.Wrapf(err, "unable to find device with name %s", deviceName), w, r)
				return
			}
		}
		ctx := context.WithValue(r.Context(), deviceKey, device)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

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
	device := r.Context().Value(deviceKey).(*model.Device)
	render.JSON(w, r, DeviceResponse{Device: device})
}

func (api *API) updateDevice(w http.ResponseWriter, r *http.Request) {
	device := r.Context().Value(deviceKey).(*model.Device)

	var params model.Device
	if err := render.DecodeJSON(r.Body, &params); err != nil {
		handleError(err, w, r)
		return
	}

	params.ID = device.ID
	params.UpdatedAt = time.Now()
	if err := api.db.Update(&params); err != nil {
		handleError(err, w, r)
		return
	}

	render.JSON(w, r, struct{}{})
}

func (api *API) deleteDevice(w http.ResponseWriter, r *http.Request) {
	device := r.Context().Value(deviceKey).(*model.Device)
	if err := api.db.DeleteStruct(&device); err != nil {
		handleError(err, w, r)
		return
	}
	render.JSON(w, r, struct{}{})
}

func (api *API) getSocket(w http.ResponseWriter, r *http.Request) {
	device := r.Context().Value(deviceKey).(*model.Device)

	resp, err := http.Get(device.Host + "/api/v1")
	if err != nil {
		handleError(err, w, r)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		handleError(err, w, r)
		return
	}

	w.Write(body)
}

func (api *API) setSocket(w http.ResponseWriter, r *http.Request) {
	device := r.Context().Value(deviceKey).(*model.Device)

	resp, err := http.Post(device.Host+"/api/v1/socket", "application/json; charset=utf-8", r.Body)
	if err != nil {
		handleError(err, w, r)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		handleError(err, w, r)
		return
	}

	w.Write(body)
}
