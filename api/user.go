package api

import (
	"net/http"
)

// func (api *API) userCtx(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		user := new(model.User)
// 		if
// 		// device := new(model.Device)
// 		// if deviceName := chi.URLParam(r, "device_name"); deviceName != "" {
// 		// 	if err := api.db.One("Name", deviceName, device); err != nil {
// 		// 		handleError(errors.Wrapf(err, "unable to find device with name %s", deviceName), w, r)
// 		// 		return
// 		// 	}
// 		// }
// 		// ctx := context.WithValue(r.Context(), deviceKey, device)
// 		// next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func (api *API) listUsers(w http.ResponseWriter, r *http.Request) {

}

func (api *API) createUser(w http.ResponseWriter, r *http.Request) {

}

func (api *API) getUser(w http.ResponseWriter, r *http.Request) {

}

func (api *API) updateUser(w http.ResponseWriter, r *http.Request) {

}

func (api *API) deleteUser(w http.ResponseWriter, r *http.Request) {

}
