package rest

import (
	"blog-api/rest/req"
	"blog-api/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type Server struct {
	server *http.Server
	router chi.Router

	Service *service.Service
}

type Config struct {
	Addr string
	Port string

	Service *service.Service
}

func NewServer(config *Config) *Server {
	return &Server{
		server: &http.Server{
			Addr:    config.Addr + ":" + config.Port,
			Handler: http.NotFoundHandler(),
		},
		router: chi.NewRouter(),

		Service: config.Service,
	}
}

func (s *Server) RunServer() error {
	return s.server.ListenAndServe()
}

func (s *Server) SetupRouter() {
	s.setupCors()

	s.router.Route("/api", func(r chi.Router) {
		r.Method("POST", "/signup", req.NewHandler(s.Service.User.Signup))
		r.Method("POST", "/signin", req.NewHandler(s.Service.User.Signin))
	})
	s.server.Handler = s.router
}

func (s *Server) setupCors() {
	s.router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}).Handler)
}
