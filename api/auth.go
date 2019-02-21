package api

import (
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/go-chi/render"

	"github.com/smarthut/smarthut/model"
)

func (api *API) authenticate(w http.ResponseWriter, r *http.Request) {
	var cred model.Credentials
	if err := render.DecodeJSON(r.Body, &cred); err != nil {
		handleError(err, w, r)
		return
	}

	// Try to find with Username or Email provided in the Login field
	user, err := model.GetUser(api.db, cred.Login)
	if err != nil {
		handleError(err, w, r)
		return
	}

	// Validate password provided
	if !user.Authenticate(cred.Password) {
		handleError(errors.New("password is not valid"), w, r)
		return
	}

	// Generate a new JWT token
	claims := jwtauth.Claims{
		"sub":   user.Username,
		"admin": user.Admin,
		"role":  user.Role,
	}
	claims.SetIssuedNow()
	// TODO: Implement refresh token
	// BODY: At the moment JWT won't expire at all
	// claims.SetExpiryIn(api.config.JWT.Exp)
	_, tokenString, err := api.tokenAuth.Encode(claims)
	if err != nil {
		handleError(err, w, r)
		return
	}

	w.Write([]byte(tokenString))
}
