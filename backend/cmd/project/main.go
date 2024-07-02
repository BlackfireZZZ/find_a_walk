package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"

	"find_a_walk/internal/handlers"
	"find_a_walk/internal/repositories"
	"find_a_walk/internal/services"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}

func main() {
	// Подключение к базе данных
	db, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание репозитория, сервиса и обработчиков
	UserRepo := repositories.NewUserRepository(db)
	EventRepo := repositories.NewEventRepository(db)
	UserService := services.NewDefaultUserService(UserRepo)
	EventService := services.NewDefaultEventService(EventRepo)
	UserHandler := handlers.NewUserHandler(UserService)
	EventHandler := handlers.NewEventHandler(EventService)

	// Настройка маршрутизатора
	r := chi.NewRouter()
	UserRouter := chi.NewRouter()
	EventRouter := chi.NewRouter()

	r.Mount("/users", UserRouter)
	UserRouter.Get("{id}", UserHandler.GetUserByID)
	UserRouter.Post("", UserHandler.CreateUser)

	r.Mount("/events", EventRouter)
	EventRouter.Get("{id}", EventHandler.GetEventByID)
	EventRouter.Post("", EventHandler.CreateEvent)

	// Запуск HTTP сервера
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
