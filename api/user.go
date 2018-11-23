package api

import (
	"context"
	"net/http"

	"github.com/go-chi/render"

	"github.com/go-chi/jwtauth"
	"github.com/pkg/errors"

	"github.com/smarthut/smarthut/model"
)

func (api *API) userCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get JWT token from context
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			handleError(err, w, r)
			return
		}

		user, err := model.GetUser(api.db, claims["sub"].(string))
		if err != nil {
			handleError(errors.Wrapf(err, "unable to find user with name %s", claims["sub"]), w, r)
			return
		}
		ctx := context.WithValue(r.Context(), userKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (api *API) adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			handleError(err, w, r)
			return
		}

		if !claims["admin"].(bool) || claims["role"] != "admin" {
			handleError(errors.New("admins inly area"), w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (api *API) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := model.AllUsers(api.db)
	if err != nil {
		handleError(err, w, r)
		return
	}
	render.JSON(w, r, users)
}

func (api *API) createUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented yet"))
}

func (api *API) getUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(userKey).(*model.JSONBucket)
	render.JSON(w, r, user)
}

func (api *API) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented yet"))
}

func (api *API) deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not implemented yet"))
}
