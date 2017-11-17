package router

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/smarthut/smarthut/handler"
	"github.com/smarthut/smarthut/model"
)

var tokenAuth *jwtauth.JwtAuth

// New initializes routes
func New() http.Handler {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		// r.Use(jwtauth.Verifier(tokenAuth))
		// r.Use(jwtauth.Authenticator)

		// APIv2 routes
		r.Route("/api/v2", func(r chi.Router) {
			// User routes
			r.Route("/user", func(r chi.Router) {
				r.Get("/", handler.ListUsers)
				r.Get("/{username}", handler.GetUser)
			})

			// Device router
			r.Route("/device", func(r chi.Router) {
				r.Get("/", handler.ListDevices)
				r.Get("/{devicename}", handler.GetDevice)

				r.Route("/{devicename}/socket", func(r chi.Router) {
					r.Get("/", handler.ListSockets)
					r.Get("/{num}", handler.GetSocket)
					r.Post("/{num}", handler.SetSocket)
				})
			})
		})
	})

	// Public routes
	r.Group(func(r chi.Router) {
		r.Post("/token", authenticateHandler)
	})

	return r
}

func authenticateHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	u, err := model.GetUser(username)
	if err != nil {
		log.Println(err)
	} else {
		if err := u.Validate(password); err != nil {
			fmt.Println("Password not match")
		} else {
			_, tokenString, _ := tokenAuth.Encode(jwtauth.Claims{
				"user_id": username,
			})
			http.SetCookie(w, &http.Cookie{
				Name:     "jwt",
				Value:    tokenString,
				Domain:   r.URL.Host,
				Expires:  time.Now().Add(1 * time.Hour),
				Secure:   false,
				HttpOnly: true,
			})
			w.Write([]byte(tokenString))
		}
	}
}
