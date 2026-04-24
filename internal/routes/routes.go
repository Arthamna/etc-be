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
		auth.POST("/register", userController.Register)
		auth.POST("/login", userController.Login)
		auth.POST("/admin/register", userController.RegisterAdmin)
	}

	// user
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		api.POST("/profile", userController.UploadPicture)
		
		// kinda confused, but let's say user dan admin punya endpoint tersendiri untuk bisa mengakses categories 
		// sehingga user dapat melihat category, lalu juga mencari dan filtering dengan lebih mudah
		
		books := api.Group("/books")
		{
			// borrow book
			books.POST("/borrow", bookController.BorrowBook)
			books.POST("/borrows", bookController.BorrowMultipleBook)

			// filtering

			// status = available / borrowed
			books.GET("/status/:status", bookController.FindByStatus)

			// categories, untuk mencari buku berdasarkan kategori
			books.GET("/categories", categoryController.GetAllCategories)
			books.GET("/category/:category", bookController.FindByCategory)

			// searching : /books/search?q=harry
			books.GET("/search", bookController.SearchBooks)
		}

	}

	// Admin 
	admin := r.Group("/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		// book
		books := admin.Group("/books")
		{
			books.POST("/returned", bookController.SetMultipleReturnedBook)
		}
		
		book := admin.Group("/book")
		{
			book.POST("", bookController.CreateBook)
			book.GET("", bookController.GetAllBooks)
			book.GET("/:id", bookController.GetBook)
			book.PUT("/:id", bookController.UpdateBook)
			book.DELETE("/:id", bookController.DeleteBook)

			book.POST("/cover", bookController.UploadBookPicture)
			book.POST("/returned", bookController.SetReturnedBook)
		}

		// category
		category := admin.Group("/category")
		{
			category.POST("", categoryController.CreateCategory)
			category.GET("", categoryController.GetAllCategories)
			category.GET("/:id", categoryController.GetCategory)
			category.PUT("/:id", categoryController.UpdateCategory)
			category.DELETE("/:id", categoryController.DeleteCategory)
		}
	}
}
