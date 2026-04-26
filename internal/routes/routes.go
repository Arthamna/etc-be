package routes

import (
	"etc-backend/internal/handlers"
	"etc-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController handlers.UserHandler, rekrutmenController handlers.RekrutmenHandler) {

	auth := r.Group("/auth")
	{
		auth.POST("/register", userController.Register)
		auth.POST("/login", userController.Login)
		auth.POST("/admin/register", userController.RegisterAdmin)
	}

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.PATCH("/user/profile", userController.UpdateUser)

		rekrutmen := api.Group("/rekrutmen")
		{
			rekrutmen.POST("", rekrutmenController.Create)
			rekrutmen.GET("", rekrutmenController.GetAll)
			rekrutmen.GET("/me", rekrutmenController.GetMyRekrutmen)
			rekrutmen.GET("/:id", rekrutmenController.GetByID)
			rekrutmen.PUT("/:id", rekrutmenController.Update)
			rekrutmen.DELETE("/:id", rekrutmenController.Delete)
		}
	}
}
