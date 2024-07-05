package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
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

	// Connect dependencies
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewDefaultUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	eventRepo := repositories.NewEventRepository(db)
	eventService := services.NewDefaultEventService(eventRepo)
	eventHandler := handlers.NewEventHandler(eventService)
	tagRepo := repositories.NewTagRepository(db)
	tagService := services.NewDefaultTagService(tagRepo)
	tagHandler := handlers.NewTagsHandler(tagService)

	go cleaner(eventService, 5*time.Minute)
	// Setting routes
	r := chi.NewRouter()
	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
		middleware.RequestID,
		middleware.Recoverer,
	)
	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", userHandler.GetUserByID)
		r.Post("/", userHandler.CreateUser)
	})

	r.Route("/events", func(r chi.Router) {
		r.Get("/{id}", eventHandler.GetEventByID)
		r.Get("/", eventHandler.GetEvents)
		r.Post("/", eventHandler.CreateEvent)
	})

	r.Route("/tags", func(r chi.Router) {
		r.Get("/", tagHandler.GetTags)
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
