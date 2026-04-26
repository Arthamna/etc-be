package handlers

import (
	"etc-backend/internal/dtos"
	"etc-backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RekrutmenHandler interface {
	Create(c *gin.Context)
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	GetMyRekrutmen(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type rekrutmenHandler struct {
	rekrutmenService services.RekrutmenService
}

func NewRekrutmenHandler(rekrutmenService services.RekrutmenService) RekrutmenHandler {
	return &rekrutmenHandler{rekrutmenService: rekrutmenService}
}

func (h *rekrutmenHandler) Create(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	var req dtos.CreateRekrutmenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.rekrutmenService.Create(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *rekrutmenHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	kegiatan := c.Query("kegiatan")
	role := c.Query("role")
	keyword := c.Query("q")

	result, err := h.rekrutmenService.GetAll(c.Request.Context(), page, limit, kegiatan, role, keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *rekrutmenHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := h.rekrutmenService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *rekrutmenHandler) GetMyRekrutmen(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	result, err := h.rekrutmenService.GetByUserID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *rekrutmenHandler) Update(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	id := c.Param("id")

	var req dtos.CreateRekrutmenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.rekrutmenService.Update(c.Request.Context(), userID, id, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *rekrutmenHandler) Delete(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	id := c.Param("id")

	if err := h.rekrutmenService.Delete(c.Request.Context(), userID, id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "rekrutmen berhasil dihapus"})
}
