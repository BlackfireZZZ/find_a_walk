package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v4/pgxpool"

	"find_a_walk/internal/handlers"
	"find_a_walk/internal/repositories"
	"find_a_walk/internal/services"
)

func main() {
	// Подключение к базе данных
	db, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание репозитория, сервиса и обработчиков
	repo := repositories.NeweventRepository(db)
	service := services.NewDefaulteventService(repo)
	handler := handlers.NeweventHandler(service)

	// Настройка маршрутизатора
	r := chi.NewRouter()
	r.Get("/events/{id}", handler.GeteventByID)
	r.Post("/events", handler.Createevent)

	// Запуск HTTP сервера
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
