package api

import (
	"net/http"

	"github.com/smarthut/smarthut/store"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/rs/cors"

	"github.com/smarthut/smarthut/conf"
)

var tokenAuth *jwtauth.JWTAuth

// API is the main REST API
type API struct {
	handler http.Handler
	db      *store.DB
	config  *conf.Configuration
	version string
}

// NewAPI instantiates a new REST API
func NewAPI(config *conf.Configuration, db *store.DB, version string) *API {
	api := &API{
		config:  config,
		db:      db,
		version: version,
	}

	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

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

	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler)

	// Protected routes
	r.Group(func(r chi.Router) {
		// Seek, verify and validate JWT tokens
		// r.Use(jwtauth.Verifier(tokenAuth))
		// r.Use(jwtauth.Authenticator)

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
				r.Route("/{device_id:[0-9]+}", func(r chi.Router) {
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
				r.With(api.deviceCtx).Get("/{device_name}", api.getDevice)
			})
		})
	})

	// Public routes
	// r.Group(func(r chi.Router) {
	// 	r.Post("/token", authenticateHandler)
	// })

	api.handler = r
	return api
}

// Start starts API at address
func (a *API) Start(addr string) {
	http.ListenAndServe(addr, a.handler)
}

// func newRouter() http.Handler {
// 	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

// 	r := chi.NewRouter()

// 	cors := cors.New(cors.Options{
// 		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
// 		AllowedOrigins:   []string{"*"},
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
// 		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
// 		ExposedHeaders:   []string{"Link"},
// 		AllowCredentials: true,
// 		MaxAge:           300,
// 	})

// 	r.Use(middleware.RequestID)
// 	r.Use(middleware.Recoverer)
// 	r.Use(cors.Handler)

// 	// Protected routes
// 	r.Group(func(r chi.Router) {
// 		// Seek, verify and validate JWT tokens
// 		// r.Use(jwtauth.Verifier(tokenAuth))
// 		// r.Use(jwtauth.Authenticator)

// 		// APIv1 routes
// 		r.Route("/api/v1", func(r chi.Router) {
// 			// User routes
// 			r.Route("/user", func(r chi.Router) {
// 				r.Get("/", handler.ListUsers)
// 				r.Get("/{username}", handler.GetUser)
// 			})

// 			// Device router
// 			r.Route("/device", func(r chi.Router) {
// 				r.Get("/", handler.ListDevices)
// 				// TODO: add POST   to CREATE device
// 				// TODO: add PUT    to UPDATE device
// 				// TODO: add DELETE to REMOVE device
// 				r.Route("/{devicename}", func(r chi.Router) {
// 					r.Get("/", handler.GetDevice)
// 					r.Post("/socket", handler.SetSocket)
// 				})
// 			})
// 		})
// 	})

// 	// Public routes
// 	r.Group(func(r chi.Router) {
// 		r.Post("/token", authenticateHandler)
// 	})

// 	return r
// }

// func authenticateHandler(w http.ResponseWriter, r *http.Request) {
// 	username := r.FormValue("username")
// 	password := r.FormValue("password")

// 	u, err := model.GetUser(username)
// 	if err != nil {
// 		log.Println(err)
// 	} else {
// 		if err := u.Validate(password); err != nil {
// 			fmt.Println("Password not match")
// 		} else {
// 			_, tokenString, _ := tokenAuth.Encode(jwtauth.Claims{
// 				"user_id": username,
// 			})
// 			http.SetCookie(w, &http.Cookie{
// 				Name:     "jwt",
// 				Value:    tokenString,
// 				Domain:   r.URL.Host,
// 				Expires:  time.Now().Add(1 * time.Hour),
// 				Secure:   false,
// 				HttpOnly: true,
// 			})
// 			w.Write([]byte(tokenString))
// 		}
// 	}
// }
