package main

import (
	"com.nguyenonline/formipro/internal/api"
	"com.nguyenonline/formipro/internal/job"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("WARN Could not load .env file, expect environment variables to be set.")
	}

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

	r.Group(func(r chi.Router) {
		r.Post("/api/letter", api.Letter)
		r.Post("/api/letter/form", api.FormLetter)
	})

	err = os.Mkdir("tmp", 0700)
	if err != nil {
		log.Println("WARN Could not create tmp directory, expect to be able to write processed files to it.")
	}

	go job.Run()

	err = http.ListenAndServe(":"+servePort(), r)
	if err != nil {
		log.Fatal(err)
	}
}

func servePort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
