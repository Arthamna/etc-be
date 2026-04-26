package routes

import (
	"etc-backend/internal/handlers"
	"etc-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController handlers.UserHandler) {
	
	// auth
	auth := r.Group("/auth")
	{
		auth.POST("/register", userController.Register)
		auth.POST("/login", userController.Login)
		auth.POST("/admin/register", userController.RegisterAdmin)
	}

	// user
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// other user routes can be added here
	}

	// Admin 
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		// admin routes can be added here
	}
}
