package handlers

import (
	"arthamna/rplLibrary/internal/dtos"
	"arthamna/rplLibrary/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	BookHandler interface {
		CreateBook(c *gin.Context)
		UploadBookPicture(c *gin.Context)
		GetAllBooks(c *gin.Context)
		BorrowMultipleBook(c *gin.Context)
		SetMultipleReturnedBook(c *gin.Context)
		GetBook(c *gin.Context)
		UpdateBook(c *gin.Context)
		DeleteBook(c *gin.Context)
		BorrowBook(c *gin.Context)
		SetReturnedBook(c *gin.Context)
		FindByCategory(c *gin.Context)
		FindByStatus(c *gin.Context)
		SearchBooks(c *gin.Context)
	}

	bookHandler struct {
		bookService services.BookService
	}
)

func NewBookHandler(bookService services.BookService) BookHandler {
	return &bookHandler{
		bookService: bookService,
	}
}

func (h *bookHandler) CreateBook(c *gin.Context) {
	var req dtos.BookCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (h *bookHandler) UploadBookPicture(c *gin.Context) {
	var req dtos.UploadBookPictureRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookService.UploadBookPicture(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, book)
}

func (h *bookHandler) GetAllBooks(c *gin.Context) {
	books, err := h.bookService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *bookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	book, err := h.bookService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *bookHandler) UpdateBook(c *gin.Context) {
	var req dtos.BookUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookService.Update(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *bookHandler) DeleteBook(c *gin.Context) {
	bookID := c.Param("id")

	err := h.bookService.Delete(c.Request.Context(), bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
}

func (h *bookHandler) BorrowBook(c *gin.Context) {
	var req dtos.BorrowBookRequest
	userId := c.MustGet("user_id").(string)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookService.BorrowBook(c.Request.Context(), req, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}


func (h *bookHandler) BorrowMultipleBook(c *gin.Context) {
	var req dtos.BorrowMultipleBookRequest
	userId := c.MustGet("user_id").(string)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.bookService.BorrowMultipleBook(c.Request.Context(), req, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *bookHandler) SetMultipleReturnedBook(c *gin.Context) {
	var req dtos.SetMultipleReturnedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.bookService.SetMultipleBookReturned(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *bookHandler) SetReturnedBook(c *gin.Context) {
	var req dtos.SetBookReturnedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookService.SetBookReturned(c.Request.Context(), req.BookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// api/get/books/:category
func (h *bookHandler) FindByCategory(c *gin.Context) {
	query := c.Param("category")
	book, err := h.bookService.FindByCategory(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

// api/get/books/:status
func (h *bookHandler) FindByStatus(c *gin.Context) {
	query := c.Param("status")
	book, err := h.bookService.FindByStatus(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *bookHandler) SearchBooks(c *gin.Context) {
    query := c.Query("q")
    if query == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "query parameter 'q' is required"})
        return
    }

    books, err := h.bookService.SearchByTitle(c.Request.Context(), query)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, books)
}