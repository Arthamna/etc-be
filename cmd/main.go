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
	"runtime/debug"

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
	settingDriveRepo := repositories.NewSettingDriveRepository(gdrive, db)

	// services
	settingDriveService := services.NewSettingDriveService(settingDriveRepo, gdrive)
	userService := services.NewUserService(userRepo, jwtService, settingDriveService, gdrive)
	rekrutmenService := services.NewRekrutmenService(rekrutmenRepo, pendaftarRepo, timRepo, historyRepo, settingDriveService)

	// handlers
	userController := handlers.NewUserHandler(userService)
	rekrutmenController := handlers.NewRekrutmenHandler(rekrutmenService)

	routes.SetupRoutes(r, userController, rekrutmenController)

	r.Run(":8080")
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

