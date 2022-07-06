package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/spinales/ASIA-backend/token"
	"github.com/spinales/ASIA-backend/util"
	"gorm.io/gorm"
)

type Server struct {
	store      *gorm.DB
	config     *util.Config
	tokenMaker token.Maker
	router     *chi.Mux
	service    *Service
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store *gorm.DB, config *util.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		tokenMaker: tokenMaker,
		service:    NewService(store),
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	router.Use(cors.AllowAll().Handler)

	router.Post("/login", server.login)
	router.Post("/register", server.register)

	router.Route("/api", func(r chi.Router) {
		r.Use(server.authMiddleware)
		r.Route("/course", func(r chi.Router) {
			r.Get("/", server.GetCoursesHandler)
			r.Post("/", server.AddCourseHandler)
			r.Route("/{id}", func(r chi.Router) {
				r.Delete("/", server.DeleteCourseHandler)
				r.Put("/", server.UpdateCourseHandler)
				r.Get("/", server.GetCourseByIDHandler)
			})
			r.Get("/{code}", server.GetCourseByCodeHandler)
			r.Get("/{name}", server.GetCourseByNameHandler)
		})
		r.Route("/user", func(r chi.Router) {
			r.Get("/", server.GetUsersHandler)
			r.Post("/", server.AddUserHandler)
			r.Route("/{id}", func(r chi.Router) {
				r.Delete("/", server.DeleteUserHandler)
				r.Put("/", server.UpdateUserHandler)
				r.Get("/", server.GetUserByIDHandler)
			})
			r.Get("/{tuition}", server.GetUserByTuitionHandler)
			r.Get("/{firstname}", server.GetUserByFirstnameHandler)
		})
		r.Get("/ranking", server.GetRankingHandler)
		r.Route("/student", func(r chi.Router) {
			r.Get("/", server.GetStudentsHandler)
			r.Post("/", server.AddStudentHandler)
			r.Route("/{id}", func(r chi.Router) {
				r.Put("/", server.UpdateStudentHandler)
				r.Delete("/", server.DeleteStudentHandler)
				r.Get("/", server.GetStudentByIDHandler)
			})
		})
	})

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	log.Println("Server running in http://localhost" + address)
	return http.ListenAndServe(address, server.router)
}
