package api

import (
	"context"
	"net/http"

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

		user := new(model.User)
		if err := api.db.One("Username", claims["sub"], &user); err != nil {
			handleError(errors.Wrapf(err, "unable to find user with name %s", claims["user"]), w, r)
			return
		}
		ctx := context.WithValue(r.Context(), userKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

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
