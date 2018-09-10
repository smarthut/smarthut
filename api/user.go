package api

import (
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth"

	"github.com/asdine/storm/q"
	"github.com/go-chi/render"

	"github.com/smarthut/smarthut/model"
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

func (api *API) authenticate(w http.ResponseWriter, r *http.Request) {
	var cred model.Credentials
	if err := render.DecodeJSON(r.Body, &cred); err != nil {
		handleError(err, w, r)
		return
	}

	// Try to find with Username or Email provided in the Login field
	var user model.User
	if err := api.db.Select(q.Or(q.StrictEq("Username", cred.Login), q.StrictEq("Email", cred.Login))).First(&user); err != nil {
		handleError(err, w, r)
		return
	}

	// Validate password provided
	if !user.Authenticate(cred.Password) {
		handleError(errors.New("password is not valid"), w, r)
		return
	}

	// Generate a new JWT token
	_, tokenString, err := api.tokenAuth.Encode(jwtauth.Claims{"user_id": user.Username})
	if err != nil {
		handleError(err, w, r)
		return
	}

	w.Write([]byte(tokenString))
}
