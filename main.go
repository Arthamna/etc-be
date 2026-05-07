package main

import (
	"etc-backend/internal/handlers"
	"etc-backend/internal/repositories"
	"etc-backend/internal/routes"
	"etc-backend/internal/services"
	"etc-backend/middleware"
	database "etc-backend/migrations"
	common "etc-backend/utils/response"
	"etc-backend/utils/storage"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	// disable in docker image
	if err := loadEnv(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("JWT_SECRET_KEY tidak ditemukan di environment variables")
	}

	db := database.ConnectToPostgresql()
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.Use(handlePanic())

	//
	jwtService := services.NewJWTService()
	gdrive := storage.NewGdrive()

	// repositories
	userRepo := repositories.NewUserRepository(db)
	rekrutmenRepo := repositories.NewRekrutmenRepository(db)
	pendaftarRepo := repositories.NewPendaftarRepository(db)
	timRepo := repositories.NewTimRepository(db)
	historyRepo := repositories.NewHistoryRepository(db)
	bookmarkRepo := repositories.NewBookmarkRepository(db)
	settingDriveRepo := repositories.NewSettingDriveRepository(gdrive, db)

	// services
	settingDriveService := services.NewSettingDriveService(settingDriveRepo, gdrive)
	userService := services.NewUserService(rekrutmenRepo, userRepo, jwtService, settingDriveService, gdrive)
	bookService := services.NewBookmarkService(bookmarkRepo)
	rekrutmenService := services.NewRekrutmenService(userRepo, rekrutmenRepo, pendaftarRepo, timRepo, historyRepo, settingDriveService)

	// handlers
	bookController := handlers.NewBookmarkHandler(bookService)
	userController := handlers.NewUserHandler(userService)
	rekrutmenController := handlers.NewRekrutmenHandler(rekrutmenService)

	routes.SetupRoutes(r, userController, rekrutmenController, bookController)

	r.Run(":8080")
}


func loadEnv() error {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("error getting current working directory: %v", err)
	}

	// Get the directory of the executable
	ex, err := os.Executable()
	if err != nil {
		return fmt.Errorf("error getting executable path: %v", err)
	}
	exPath := filepath.Dir(ex)

	// List of possible locations for .env file
	envLocations := []string{
		filepath.Join(cwd, ".env"),
		filepath.Join(exPath, ".env"),
		"/var/www/etc-backend/.env", // Add the expected location when run as a service
	}

	// Try to load .env from each location
	for _, loc := range envLocations {
		err := godotenv.Load(loc)
		if err == nil {
			fmt.Printf("Loaded .env from: %s\n", loc)
			return nil
		}
	}

	return fmt.Errorf("no .env file found in any of the expected locations")
}

func handlePanic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var err error
				if e, ok := r.(error); ok {
					err = e
				} else {
					err = fmt.Errorf("%v", r)
				}
				fmt.Printf("\n[recovery] panic occurred: %v\n", err)
				stack := debug.Stack()
				fmt.Fprintln(os.Stderr, string(stack))

				common.BuildErrorResponse("Internal Server Error", err.Error(), nil)
			}
		}()

		ctx.Next()
	}
}

