package main

import (
	"context"
	"log"
	"net/http"
	"os"

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

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewDefaultUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	// EventRepo := repositories.NewEventRepository(db)
	// EventService := services.NewDefaultEventService(EventRepo)
	// EventHandler := handlers.NewEventHandler(EventService)

	r := chi.NewRouter()
	r.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.RedirectSlashes,
	)
	r.Mount("/api/v1", r)

	r.Route("/users", func(r chi.Router) {
		r.Get("/{id}", userHandler.GetUserByID)
		r.Post("/", userHandler.CreateUser)
	})

	// r.Mount("/events", EventRouter)
	// EventRouter.Get("{id}", EventHandler.GetEventByID)
	// EventRouter.Post("", EventHandler.CreateEvent)

	// Запуск HTTP сервера
	serverPort := os.Getenv("SERVER_PORT")
	log.Println("Starting server on: ", serverPort)
	log.Fatal(http.ListenAndServe("localhost:"+serverPort, r))
}
