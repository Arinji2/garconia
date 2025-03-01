package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"

	"github.com/arinji2/garconia/routes"
	"github.com/arinji2/garconia/sqlite"
)

func main() {
	environ := os.Environ()
	for _, e := range environ {
		fmt.Println(e)
	}
	isProduction := os.Getenv("ENVIRONMENT") == "PRODUCTION"
	err := godotenv.Load()
	if err != nil && !isProduction {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Server Started on Port 8080")
	db, err := sqlite.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB Found and Ready")
	db.Close()
	r := chi.NewRouter()
	frontendURL := os.Getenv("FRONTEND_URL")
	fmt.Println("Allowing Origins: ", frontendURL)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{frontendURL},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Post("/add", routes.AddEmailRoute)
		r.Post("/verify", routes.VerifyEmailRoute)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Landing Backend: Health Check"))
		render.Status(r, 200)
	})

	if err := http.ListenAndServe(":8080", r); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %s\n", err)
	}
}
