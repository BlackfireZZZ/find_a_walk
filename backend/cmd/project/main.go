package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
		log.Println(err)
	}
	defer db.Close()

	tokenAuth := jwtauth.New(os.Getenv("TOKEN_ALG"), []byte(os.Getenv("SECRET_TOKEN")), nil)

	mainRepositories := repositories.InitRepositores(db)
	mainServices := services.InitServices(
		&mainRepositories.UserRepository,
		tokenAuth,
		&mainRepositories.EventRepository,
		&mainRepositories.TagRepository)
	mainHandlers := handlers.InitHandlers(
		&mainServices.UserService,
		&mainServices.EventService,
		&mainServices.TagService,
	)

	go cleaner(&mainServices.EventService, 60*time.Minute)
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

	r.Post("/auth/login", mainHandlers.AuthHandler.Login)

	// Public
	r.Route("/users", func(r chi.Router) {
		r.With(jwtAuthMiddlewares...).Get("/{id}", mainHandlers.UserHandler.GetUserByID)
		r.With(jwtAuthMiddlewares...).Get("/me", mainHandlers.UserHandler.GetUserProfile)
		r.With(jwtAuthMiddlewares...).Post("/interests", mainHandlers.UserHandler.CreateInterest)
		r.With(jwtAuthMiddlewares...).Delete("/interests", mainHandlers.UserHandler.DeleteInterests)
		r.Post("/", mainHandlers.UserHandler.CreateUser)
	})

	r.Route("/events", func(r chi.Router) {
		r.With(jwtAuthMiddlewares...).Delete("/{id}", mainHandlers.EventHandler.DeleteEvent)
		r.With(jwtAuthMiddlewares...).Get("/{id}", mainHandlers.EventHandler.GetEventByID)
		r.With(jwtAuthMiddlewares...).Post("/", mainHandlers.EventHandler.CreateEvent)
		r.With(jwtAuthMiddlewares...).Get("/me", mainHandlers.EventHandler.GetMyEvents)
		r.With(jwtAuthMiddlewares...).Post("/{id}/member", mainHandlers.EventHandler.CreateEventMember)
		r.With(jwtAuthMiddlewares...).Delete("/{id}/member", mainHandlers.EventHandler.DeleteEventMember)
		r.With(jwtAuthMiddlewares...).Get("/{id}/members", mainHandlers.EventHandler.GetEventMembers)
		r.With(jwtAuthMiddlewares...).Get("/members/me", mainHandlers.EventHandler.GetMyEventMembers)

		r.Get("/", mainHandlers.EventHandler.GetEvents)
	})

	r.Route("/tags", func(r chi.Router) {
		r.Get("/", mainHandlers.TagsHandler.GetTags)
	})

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
