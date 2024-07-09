package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

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

	repositories := repositories.InitRepositores(db)
	services := services.InitServices(
		&repositories.UserRepository,
		tokenAuth,
		&repositories.EventRepository,
		&repositories.TagRepository)
	handlers := handlers.InitHandlers(
		&services.UserService,
		&services.EventService,
		&services.TagService,
	)

	go cleaner(&services.EventService, 60*time.Minute)
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

	r.Post("/auth/login", handlers.AuthHandler.Login)

	// Public
	r.Route("/users", func(r chi.Router) {
		r.With(jwtAuthMiddlewares...).Get("/{id}", handlers.UserHandler.GetUserByID)
		r.With(jwtAuthMiddlewares...).Get("/me", handlers.UserHandler.GetUserProfile)
		r.With(jwtAuthMiddlewares...).Post("/interests", handlers.UserHandler.CreateInterest)
		r.With(jwtAuthMiddlewares...).Delete("/interests", handlers.UserHandler.DeleteInterests)
		r.Post("/", handlers.UserHandler.CreateUser)
	})

	r.Route("/events", func(r chi.Router) {
		r.With(jwtAuthMiddlewares...).Delete("/{id}", handlers.EventHandler.DeleteEvent)
		r.With(jwtAuthMiddlewares...).Get("/{id}", handlers.EventHandler.GetEventByID)
		r.With(jwtAuthMiddlewares...).Post("/", handlers.EventHandler.CreateEvent)
		r.Get("/", handlers.EventHandler.GetEvents)
	})

	r.Route("/tags", func(r chi.Router) {
		r.Get("/", handlers.TagsHandler.GetTags)
	})

	// Start HTTP server
	log.Println("Starting server on: ", os.Getenv("SERVER_ADRESS"))
	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_ADRESS"), r))
}

func cleaner(service *services.EventService, duration time.Duration) {
	for {
		time.Sleep(duration)
		err := service.DeleteExpiredEvents(context.Background())
		if err != nil {
			log.Println(err)
		}
	}

}
