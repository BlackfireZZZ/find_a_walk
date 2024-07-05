package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"

	"github.com/go-chi/render"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"

	"find_a_walk/internal/handlers"
	"find_a_walk/internal/repositories"
	"find_a_walk/internal/services"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	// Connect to DB
	db, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tokenAuth := jwtauth.New(os.Getenv("TOKEN_ALG"), []byte(os.Getenv("SECRET_TOKEN")), nil)

	// Connect dependencies
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewDefaultUserService(userRepo, tokenAuth)
	userHandler := handlers.NewUserHandler(userService)
	eventRepo := repositories.NewEventRepository(db)
	eventService := services.NewDefaultEventService(eventRepo)
	eventHandler := handlers.NewEventHandler(eventService)
	tagRepo := repositories.NewTagRepository(db)
	tagService := services.NewDefaultTagService(tagRepo)
	tagHandler := handlers.NewTagsHandler(tagService)

	authHandler := handlers.NewAuthHandler(userService)

	// Setting routes
	r := chi.NewRouter()

	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.RequestID,
		middleware.Recoverer,
	)
	r.Mount("/api/v1", r)

	jwtAuthMiddlewares := []func(http.Handler) http.Handler{
		jwtauth.Verifier(tokenAuth),
		jwtauth.Authenticator(tokenAuth),
	}

	r.Post("/auth/login", authHandler.Login)

	// Public
	r.Route("/users", func(r chi.Router) {
		r.With(jwtAuthMiddlewares...).Get("/{id}", userHandler.GetUserByID)
		r.Post("/", userHandler.CreateUser)
	})

	r.Route("/events", func(r chi.Router) {
		r.With(jwtAuthMiddlewares...).Get("/{id}", eventHandler.GetEventByID)
		r.Get("/", eventHandler.GetEvents)
		r.With(jwtAuthMiddlewares...).Post("/", eventHandler.CreateEvent)
	})

	r.Route("/tags", func(r chi.Router) {
		r.Get("/", tagHandler.GetTags)
	})

	// Start HTTP server
	log.Println("Starting server on: ", os.Getenv("SERVER_ADRESS"))
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_ADRESS"), r))
}
