package routes

import (
	"arthamna/rplLibrary/internal/handlers"
	"arthamna/rplLibrary/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userController handlers.UserHandler, bookController handlers.BookHandler, categoryController handlers.CategoryHandler) {
	
	// auth
	auth := r.Group("/auth")
	{
		// register user sekaligus onboarding data awal user
		auth.POST("/register", userHandler.Register)

		// login user
		auth.POST("/login", userHandler.Login)

		// update data user yang sedang login
		auth.PATCH("/me", middleware.AuthMiddleware(), userHandler.UpdateUser)

		// tampilkan data user yang sedang login
		auth.GET("/me", middleware.AuthMiddleware(), userHandler.GetMe)

		// upload profile picture user yang sedang login
		auth.POST("/picture", middleware.AuthMiddleware(), userHandler.UploadPicture)

		// tampilkan bookmark milik user yang sedang login
		auth.GET("/bookmarks", middleware.AuthMiddleware(), userHandler.GetBookmarks)
	}

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// rekrutmen
		rekrutmen := api.Group("/rekrutmen")
		{
			// buat rekrutmen baru oleh user yang sedang login
			rekrutmen.POST("", rekrutmenHandler.Create)

			// cari semua rekrutmen dengan pagination
			rekrutmen.GET("", rekrutmenHandler.GetAll)

			// sort / filter rekrutmen berdasarkan tipe
			rekrutmen.GET("/sort/type", rekrutmenHandler.GetByType)

			// sort / filter rekrutmen berdasarkan role
			rekrutmen.GET("/sort/role", rekrutmenHandler.GetByRole)

			// tampilkan semua rekrutmen yang dibuat oleh user yang sedang login
			rekrutmen.GET("/mine", rekrutmenHandler.GetMyRekrutmen)

			// tampilkan detail rekrutmen berdasarkan id rekrutmen dan id user yang sedang login
			// hasil: data rekrutmen + list id pendaftar (accepted/pending/rejected)
			rekrutmen.GET("/:id", rekrutmenHandler.GetByID)

			// update rekrutmen
			rekrutmen.PUT("/:id", rekrutmenHandler.Update)

			// hapus rekrutmen
			rekrutmen.DELETE("/:id", rekrutmenHandler.Delete)

			// apply ke rekrutmen dengan id user yang sedang login
			rekrutmen.POST("/:id/apply", rekrutmenHandler.Apply)

			// upload file CV untuk pendaftaran rekrutmen
			rekrutmen.POST("/:id/apply/cv", rekrutmenHandler.UploadCV)

			// upload file portfolio untuk pendaftaran rekrutmen
			rekrutmen.POST("/:id/apply/portfolio", rekrutmenHandler.UploadPortfolio)

			// refresh status pendaftaran setelah disetujui agar user masuk ke tim
			rekrutmen.PATCH("/:id/apply/:pendaftar_id/status", rekrutmenHandler.RefreshApplyStatus)

			// tampilkan detail pendaftar rekrutmen
			// hasil: id pendaftar, id rekrutmen, id user, dan detail pendaftar
			rekrutmen.GET("/:id/applicants/:pendaftar_id", rekrutmenHandler.GetApplicantDetail)
		}

		// tim
		tim := api.Group("/tim")
		{
			// tampilkan anggota tim berdasarkan tim
			tim.GET("/:id/members", rekrutmenHandler.GetTeamMembers)

			// berikan rating ke anggota tim
			// hanya untuk anggota yang status pendaftarannya accepted
			tim.POST("/:id/members/:user_id/rating", rekrutmenHandler.GiveMemberRating)
		}
	}
}