package api

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/pkg/errors"

	"github.com/smarthut/smarthut/model"
)

func (api *API) deviceCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		deviceName := chi.URLParam(r, "device")
		device, err := model.GetDevice(api.db, deviceName)
		if err != nil {
			handleError(errors.Wrapf(err, "unable to find device with name %s", deviceName), w, r)
			return
		}
		ctx := context.WithValue(r.Context(), deviceKey, device)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) listDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := model.AllDevices(api.db)
	if err != nil {
		handleError(err, w, r)
		return
	}
	render.JSON(w, r, devices)
}

func (api *API) createDevice(w http.ResponseWriter, r *http.Request) {
	var data model.Device
	if err := render.DecodeJSON(r.Body, &data); err != nil {
		handleError(err, w, r)
		return
	}
	device, err := model.NewDevice(data.Name, data.Host, data.Title)
	if err != nil {
		handleError(err, w, r)
		return
	}
	if err := api.db.Save(device); err != nil {
		handleError(err, w, r)
		return
	}
	render.JSON(w, r, device)
}

func (api *API) getDevice(w http.ResponseWriter, r *http.Request) {
	device := r.Context().Value(deviceKey).(*model.Device)
	render.JSON(w, r, device)
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
	if err := device.Delete(api.db); err != nil {
		handleError(err, w, r)
		return
	}
	render.JSON(w, r, struct{}{})
}

func (api *API) getDeviceHistory(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented yet"))
}

func (api *API) getSocket(w http.ResponseWriter, r *http.Request) {
	device := r.Context().Value(deviceKey).(*model.Device)
	resp, err := http.Get(device.Host + "/api/v1")
	if err != nil {
		handleError(err, w, r)
		return
	}
	defer resp.Body.Close()
	if _, err := io.Copy(w, resp.Body); err != nil {
		handleError(err, w, r)
		return
	}
}

func (api *API) setSocket(w http.ResponseWriter, r *http.Request) {
	device := r.Context().Value(deviceKey).(*model.Device)
	resp, err := http.Post(device.Host+"/api/v1/socket", "application/json; charset=utf-8", r.Body)
	if err != nil {
		handleError(err, w, r)
		return
	}
	defer resp.Body.Close()
	if _, err := io.Copy(w, resp.Body); err != nil {
		handleError(err, w, r)
		return
	}
}
