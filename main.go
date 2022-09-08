package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"seqolah-qu/controllers"
	"seqolah-qu/database"
	"seqolah-qu/middlewares"
	"seqolah-qu/repositories"
	"seqolah-qu/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	db := database.ConnectDatabase()

	schoolRepo := repositories.NewSchoolRepository(db)
	schoolService := services.NewSchoolService(schoolRepo)
	schoolController := controllers.NewSchoolController(r, schoolService)

	r.Get("/check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halo Dunia!"))
	})

	r.With(middlewares.JwtMiddleware).Get("/{id}", schoolController.FindSchoolById())

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	http.ListenAndServe(port, r)
}
