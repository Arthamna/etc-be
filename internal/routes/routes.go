package routes

import (
	"etc-backend/internal/handlers"
	"etc-backend/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, userHandler handlers.UserHandler, rekrutmenHandler handlers.RekrutmenHandler, bookmarkHandler handlers.BookmarkHandler) {

	// user
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

		// tampilkan semua bookmark milik user yang sedang login
		auth.GET("/bookmarks", middleware.AuthMiddleware(), userHandler.GetBookmarks)
		// kalau tampilkan detail, tembak ke get id rekrutmen (bagian cari)
		

		// tambah bookmark
		auth.POST("/:rekrutmen_id/bookmark", middleware.AuthMiddleware(), bookmarkHandler.AddBookmark)

		// hapus bookmark
		auth.DELETE("/:rekrutmen_id/bookmark", middleware.AuthMiddleware(), bookmarkHandler.RemoveBookmark)
		
		// tampilkan semua bookmark
		// auth.GET("/bookmark", middleware.AuthMiddleware(), bookmarkHandler.GetBookmarks) 
		
	}

	// contoh : GET /rekrutmen?page=1&limit=10&kegiatan=riset&role=backend&q=java
	r.GET("/api/rekrutmen", rekrutmenHandler.GetAll)
	// .GET("/sort/type/:type", rekrutmenHandler.GetByType)
	r.GET("/sort/type/:type", rekrutmenHandler.GetByType)
	r.GET("/sort/role/:role", rekrutmenHandler.GetByRole)

	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// rekrutmen
		rekrutmen := api.Group("/rekrutmen")
		{
			// buat //
			// buat rekrutmen baru oleh user yang sedang login
			rekrutmen.POST("", rekrutmenHandler.Create)

			// tampilkan semua rekrutmen yang dibuat oleh user yang sedang login
			rekrutmen.GET("/mine", rekrutmenHandler.GetMyRekrutmen)

			// tampilkan detail rekrutmen berdasarkan id rekrutmen dan (id user pembuat rekrutmen)
			// hasil: data rekrutmen + list id pendaftar (accepted/pending/rejected)
			// id = rekrutmen_id
			rekrutmen.GET("/applicants/:id", rekrutmenHandler.GetAppliedByID)
			
			// tampilkan detail pendaftar rekrutmen dari id rekrutmen dan id pendaftar, beserta history jika ada
			// hasil: id pendaftar, id rekrutmen, id user, dan detail pendaftar
			rekrutmen.GET("/:id/applicants/:pendaftar_id", rekrutmenHandler.GetApplicantDetail)

			// set status diterima/ditolak
			rekrutmen.PATCH("/:rekrutmen_id/apply/:pendaftar_id/accept", rekrutmenHandler.AcceptPendaftar)
			rekrutmen.PATCH("/:rekrutmen_id/apply/:pendaftar_id/reject", rekrutmenHandler.RejectPendaftar)

			// update rekrutmen
			rekrutmen.PUT("/:id", rekrutmenHandler.Update)

			// hapus rekrutmen
			rekrutmen.DELETE("/:id", rekrutmenHandler.Delete)

			// cari //
			// cari semua rekrutmen dengan pagination
			// contoh : GET /rekrutmen?page=1&limit=10&kegiatan=riset&role=backend&q=java
			// rekrutmen.GET("", rekrutmenHandler.GetAll)

			// tampilkan rekrutmen yang pernah di-apply oleh user yang sedang login
			rekrutmen.GET("/applied", rekrutmenHandler.GetAppliedRekrutmen)

			// get by id
			rekrutmen.GET("/:id", rekrutmenHandler.GetByID)

			// sort / filter rekrutmen berdasarkan tipe kegiatan, (projek, riset, lomba)
			// rekrutmen.GET("/sort/type/:type", rekrutmenHandler.GetByType)

			// sort / filter rekrutmen berdasarkan role
			// rekrutmen.GET("/sort/role/:role", rekrutmenHandler.GetByRole)

			// apply //
			// id = rekrutmen_id
			// apply ke rekrutmen dengan id user yang sedang login
			// persyaratan, sudah ada cv dan portofolio (jadi tembak endpoint cv ama porto dulu)
			rekrutmen.POST("/:id/apply", rekrutmenHandler.Apply)

			// upload file CV untuk pendaftaran rekrutmen
			rekrutmen.POST("/:id/apply/cv", rekrutmenHandler.UploadCV)

			// upload file portfolio untuk pendaftaran rekrutmen
			rekrutmen.POST("/:id/apply/portfolio", rekrutmenHandler.UploadPortfolio)

			// refresh status pendaftaran setelah disetujui agar user masuk ke tim
			rekrutmen.PATCH("/:rekrutmen_id/apply/:pendaftar_id/refresh-status", rekrutmenHandler.RefreshApplyStatus)

		}

		// tim
		tim := api.Group("/tim")
		{
			// tampilkan anggota tim berdasarkan tim
			// id = tim_id
			tim.GET("/:id/members", rekrutmenHandler.GetTeamMembers)

			// berikan rating ke anggota tim
			// hanya untuk anggota yang status pendaftarannya accepted
			tim.POST("/:id/members/:target_id/rating", rekrutmenHandler.GiveMemberRating)
		}
	}
}