package main

import (
	"arthamna/rplLibrary/internal/handlers"
	"arthamna/rplLibrary/internal/repositories"
	"arthamna/rplLibrary/internal/routes"
	"arthamna/rplLibrary/internal/services"
	"arthamna/rplLibrary/pkg/auth"
	"arthamna/rplLibrary/pkg/database"
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

	jwtService := auth.NewJWTService()
	// repositories
	userRepo := repositories.NewUserRepository(db)
	bookRepo := repositories.NewBookRepository(db)
	bookBorrowing := repositories.NewBookBorrowingRepository(db)
	categoryRepo := repositories.NewCategoryRepository(db)

	// services
	userService := services.NewUserService(userRepo, jwtService)
	bookService := services.NewBookService(bookRepo, categoryRepo, userRepo, bookBorrowing)
	categoryService := services.NewCategoryService(categoryRepo)

	// handlers
	userController := handlers.NewUserHandler(userService)
	bookController := handlers.NewBookHandler(bookService)
	categoryController := handlers.NewCategoryHandler(categoryService)

	routes.SetupRoutes(r, userController, bookController, categoryController)

	r.Run(":8080")
}