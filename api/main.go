package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/joho/godotenv"

	"github.com/arinji2/garconia/routes"
	"github.com/arinji2/garconia/sqlite"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Server Started")
	db, err := sqlite.NewConnection()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB Found and Ready")
	db.Close()
	r := chi.NewRouter()

	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Post("/add", routes.AddEmailRoute)
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Landing Backend: Health Check"))
		render.Status(r, 200)
	})

	if err := http.ListenAndServe(":8080", r); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %s\n", err)
	}
}
