package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/rs/cors"

	"github.com/smarthut/smarthut"
	"github.com/smarthut/smarthut/store"
)

var tokenAuth *jwtauth.JWTAuth

// API is the main REST API
type API struct {
	handler   http.Handler
	db        *store.DB
	config    *smarthut.Configuration
	tokenAuth *jwtauth.JWTAuth
	version   string
}

// NewAPI instantiates a new REST API
func NewAPI(config *smarthut.Configuration, db *store.DB, version string) *API {
	api := &API{
		config:    config,
		db:        db,
		tokenAuth: jwtauth.New("HS256", []byte(config.JWT.Secret), nil),
		version:   version,
	}

	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		// AllowedOrigins:   []string{config.API.Host},
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	// disable log colors
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{
		Logger:  log.New(os.Stdout, "", log.LstdFlags),
		NoColor: true,
	})

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler)

	// Public routes
	r.Group(func(r chi.Router) {
		// Returns the JWT token
		r.Post("/auth", api.authenticate)

		// TODO: Remove API calls from public routes
		// BODY: All API calls must provide JWT token

		// APIv1 routes
		r.Route("/api/v1", func(r chi.Router) {
			// User routes
			r.Route("/user", func(r chi.Router) {
				r.Get("/", api.listUsers)
				r.Post("/", api.createUser)
				r.Route("/{login}", func(r chi.Router) {
					r.Get("/", api.getUser)
					r.Put("/", api.updateUser)
					r.Delete("/", api.deleteUser)
				})
			})
			// Device router
			r.Route("/device", func(r chi.Router) {
				r.Get("/", api.listDevices)
				r.Post("/", api.createDevice)
				r.Route("/{device}", func(r chi.Router) {
					r.Use(api.deviceCtx)
					r.Get("/", api.getDevice)
					r.Put("/", api.updateDevice)
					r.Delete("/", api.deleteDevice)
					// Socket operations
					r.Route("/socket", func(r chi.Router) {
						r.Get("/", api.getSocket)
						r.Post("/", api.setSocket)
					})
				})
			})
		})
	})

	// Protected routes
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(api.tokenAuth))
		r.Use(jwtauth.Authenticator)

		// This endpoint is used to test a JWT token
		r.Get("/token", func(w http.ResponseWriter, r *http.Request) {
			_, claims, err := jwtauth.FromContext(r.Context())
			if err != nil {
				handleError(err, w, r)
				return
			}
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["sub"])))
		})

		// APIv2 routes
		r.Route("/api/v2", func(r chi.Router) {
			// r.Use(api.userCtx)

			// User routes
			r.Get("/users", api.listUsers)
			r.Route("/user", func(r chi.Router) {
				r.Get("/", api.listUsers)
				r.Post("/", api.createUser)
				r.Route("/{login}", func(r chi.Router) {
					r.Get("/", api.getUser)
					r.Put("/", api.updateUser)
					r.Delete("/", api.deleteUser)
				})
			})

			// Device routes
			r.Route("/device", func(r chi.Router) {
				r.Get("/", api.listDevices)
				r.Post("/", api.createDevice)
				r.Route("/{device}", func(r chi.Router) {
					r.Use(api.deviceCtx)
					r.Get("/", api.getDevice)
					r.Put("/", api.updateDevice)
					r.Delete("/", api.deleteDevice)
					r.Get("/history", api.getDeviceHistory)
					// Socket operations
					r.Route("/socket", func(r chi.Router) {
						r.Get("/", api.getSocket)
						r.Post("/", api.setSocket)
					})
				})
			})

			// JSON bucket routes
			r.Get("/buckets", api.listBuckets)
			r.Route("/bucket", func(r chi.Router) {
				r.Get("/", api.listBuckets)
				r.Post("/", api.createBucket)
				r.Route("/{bucket}", func(r chi.Router) {
					r.Use(api.bucketCtx)
					r.Get("/", api.getBucket)
					r.Put("/", api.updateBucket)
					r.Delete("/", api.deleteBucket)
				})
			})

			// Channel
			r.Route("/channel", func(r chi.Router) {
				// Return the list of all channels
				r.Get("/", api.listThings)
				r.Route("/{channel}", func(r chi.Router) {
					r.Get("/", api.getThing)
					r.Get("/value", api.getThing)
				})
			})
		})
	})

	api.handler = r
	return api
}

// Start starts API at address
func (a *API) Start(addr string) {
	http.ListenAndServe(addr, a.handler)
}
