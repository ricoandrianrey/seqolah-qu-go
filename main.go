package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"seqolah-qu/controllers"
	"seqolah-qu/database"
	"seqolah-qu/repositories"
	"seqolah-qu/services"
	"seqolah-qu/utils"

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
	userRepo := repositories.NewUserRepository(db)

	schoolService := services.NewSchoolService(schoolRepo)
	authService := services.NewAuthService(userRepo)

	schoolController := controllers.NewSchoolController(schoolService)
	authController := controllers.NewAuthController(authService)

	r.Get("/check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Halo Dunia!"))
	})

	r.Post("/auth/register", utils.HandlerWrapper(authController.Register))
	r.Post("/auth/login", utils.HandlerWrapper(authController.Login))

	r.Post("/school", utils.HandlerWrapper(schoolController.CreateSchool))
	r.Put("/school/{id}", utils.HandlerWrapper(schoolController.UpdateSchool))
	r.Get("/school/{id}", utils.HandlerWrapper(schoolController.FindSchoolById))
	r.Delete("/school/{id}", utils.HandlerWrapper(schoolController.DeleteSchool))

	port := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))
	http.ListenAndServe(port, r)
}
