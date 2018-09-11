package api

import (
	"errors"
	"net/http"

	"github.com/asdine/storm/q"
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
	claims := jwtauth.Claims{
		"sub":  user.Username,
		"role": user.Role,
	}
	claims.SetIssuedNow()
	// TODO: implement refresh token
	// claims.SetExpiryIn(api.config.JWT.Exp)
	_, tokenString, err := api.tokenAuth.Encode(claims)
	if err != nil {
		handleError(err, w, r)
		return
	}

	w.Write([]byte(tokenString))
}
