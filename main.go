package main

import (
	"com.nguyenonline/formipro/internal/api"
	"com.nguyenonline/formipro/job"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"log"
	"net/http"
)

const PORT = ":22222"

func main() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Letter api
	r.Post("/api/letter/create", api.Letter)
	r.Post("/api/letter/create-form", api.FormLetter)

	// Start scheduled jobs
	go job.DeleteProcessedTmpFiles()

	err := http.ListenAndServe(PORT, r)
	if err != nil {
		log.Fatal(err)
	}
}
