package handlers

import (
	"arthamna/rplLibrary/internal/dtos"
	"arthamna/rplLibrary/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	CategoryHandler interface {
		CreateCategory(c *gin.Context)
		GetAllCategories(c *gin.Context)
		GetCategory(c *gin.Context)
		UpdateCategory(c *gin.Context)
		DeleteCategory(c *gin.Context)
	}

	categoryHandler struct {
		categoryService services.CategoryService
	}
)

func NewCategoryHandler(categoryService services.CategoryService) CategoryHandler {
	return &categoryHandler{
		categoryService: categoryService,
	}
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var req dtos.CategoryCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.categoryService.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, category)
}

func (h *categoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.categoryService.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *categoryHandler) GetCategory(c *gin.Context) {
	id := c.Param("id")
	category, err := h.categoryService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var req dtos.CategoryUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.categoryService.Update(c.Request.Context(), id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	err := h.categoryService.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book deleted"})
}
