package main

import (
	"etc-backend/internal/handlers"
	"etc-backend/internal/repositories"
	"etc-backend/internal/routes"
	"etc-backend/internal/services"
	database "etc-backend/migrations"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("JWT_SECRET_KEY tidak ditemukan di environment variables")
	}

	db := database.ConnectToPostgresql()
	r := gin.Default()

	//

	jwtService := services.NewJWTService()
	// repositories
	userRepo := repositories.NewUserRepository(db)

	// services
	userService := services.NewUserService(userRepo, jwtService)

	// handlers
	userController := handlers.NewUserHandler(userService)

	routes.SetupRoutes(r, userController)

	r.Run(":8080")
}
