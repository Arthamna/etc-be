package handlers

import (
	"etc-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookmarkHandler interface {
	AddBookmark(c *gin.Context)
	RemoveBookmark(c *gin.Context)
	GetBookmarks(c *gin.Context) 
}

type bookmarkHandler struct {
	bookmarkService services.BookmarkService
}

func NewBookmarkHandler(bookmarkService services.BookmarkService) BookmarkHandler {
	return &bookmarkHandler{
		bookmarkService: bookmarkService,
	}
}

func (h *bookmarkHandler) AddBookmark(c *gin.Context) {
	userID := c.GetString("user_id") 
	rekrutmenID := c.Param("rekrutmen_id")

	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := h.bookmarkService.AddBookmark(c.Request.Context(), userID, rekrutmenID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "bookmark berhasil ditambahkan"})
}

func (h *bookmarkHandler) RemoveBookmark(c *gin.Context) {
	userID := c.GetString("user_id")
	rekrutmenID := c.Param("rekrutmen_id")

	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	if err := h.bookmarkService.RemoveBookmark(c.Request.Context(), userID, rekrutmenID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "bookmark berhasil dihapus"})
}

func (h *bookmarkHandler) GetBookmarks(c *gin.Context) {
	userID := c.GetString("user_id")

	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	data, err := h.bookmarkService.GetBookmarks(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}