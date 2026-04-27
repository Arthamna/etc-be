package routes

import (
	"etc-backend/internal/handlers"
	"etc-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userHandler handlers.UserHandler, rekrutmenHandler handlers.RekrutmenHandler) {

	// user
	auth := r.Group("/auth")
	{
		auth.POST("/register", userHandler.Register)
		auth.POST("/login", userHandler.Login)
		// auth.POST("/admin/register", userHandler.RegisterAdmin)

		auth.PATCH("/update", middleware.AuthMiddleware(), userHandler.UpdateUser)
		auth.POST("/picture",middleware.AuthMiddleware(), userHandler.UploadPicture)
		auth.GET("/me", middleware.AuthMiddleware(), userHandler.GetMe)  
	}


	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		rekrutmen := api.Group("/rekrutmen")
		{
			rekrutmen.POST("", rekrutmenHandler.Create)
			rekrutmen.GET("", rekrutmenHandler.GetAll) 
			rekrutmen.GET("/rekrutmen", rekrutmenHandler.GetMyRekrutmen)
			rekrutmen.GET("/:id", rekrutmenHandler.GetByID)
			rekrutmen.PUT("/:id", rekrutmenHandler.Update)
			rekrutmen.DELETE("/:id", rekrutmenHandler.Delete)
		}
	}
}
