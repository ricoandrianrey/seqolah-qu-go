package main

import (
	"net/http"
	"seqolah-qu/controllers"
	"seqolah-qu/database"
	"seqolah-qu/repositories"
	"seqolah-qu/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	db := database.ConnectDatabase()

	schoolRepo := repositories.NewSchoolRepository(db)
	schoolService := services.NewSchoolService(schoolRepo)
	schoolController := controllers.NewSchoolController(r, schoolService)

	r.Get("/{id}", schoolController.FindSchoolById())

	http.ListenAndServe(":3000", r)
}
