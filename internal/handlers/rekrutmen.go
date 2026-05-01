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
	GetByID(c *gin.Context)
	GetAll(c *gin.Context)
	GetAppliedByID(c *gin.Context)
	GetAppliedRekrutmen(c *gin.Context)
	GetMyRekrutmen(c *gin.Context)
	AcceptPendaftar(c *gin.Context)
	RejectPendaftar(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetByType(c *gin.Context)
	GetByRole(c *gin.Context)
	Apply(c *gin.Context)
	UploadCV(c *gin.Context)
	UploadPortfolio(c *gin.Context)
	RefreshApplyStatus(c *gin.Context)
	GetApplicantDetail(c *gin.Context)
	GetTeamMembers(c *gin.Context)
	GiveMemberRating(c *gin.Context)
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

func (h *rekrutmenHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := h.rekrutmenService.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
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

func (h *rekrutmenHandler) GetAppliedByID(c *gin.Context) {
	id := c.Param("id")

	result, err := h.rekrutmenService.GetAppliedByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *rekrutmenHandler) GetAppliedRekrutmen(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	result, err := h.rekrutmenService.GetAppliedRekrutmen(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

func (h *rekrutmenHandler) GetByType(c *gin.Context) {
	kegiatan := c.Param("type")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	result, err := h.rekrutmenService.GetAll(c.Request.Context(), page, limit, kegiatan, "", "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *rekrutmenHandler) GetByRole(c *gin.Context) {
	role := c.Param("role")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	result, err := h.rekrutmenService.GetAll(c.Request.Context(), page, limit, "", role, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}

func (h *rekrutmenHandler) Apply(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	rekrutmenID := c.Param("id")

	var req dtos.ApplyRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.rekrutmenService.Apply(c.Request.Context(), userID, rekrutmenID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *rekrutmenHandler) UploadCV(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	rekrutmenID := c.Param("id")

	file, err := c.FormFile("cv")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file cv diperlukan"})
		return
	}

	result, err := h.rekrutmenService.UploadCV(c.Request.Context(), userID, rekrutmenID, file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *rekrutmenHandler) UploadPortfolio(c *gin.Context) {
	userID := c.MustGet("user_id").(string)
	rekrutmenID := c.Param("id")

	file, err := c.FormFile("portfolio")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file portfolio diperlukan"})
		return
	}

	result, err := h.rekrutmenService.UploadPortfolio(c.Request.Context(), userID, rekrutmenID, file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *rekrutmenHandler) RefreshApplyStatus(c *gin.Context) {
	rekrutmenID := c.Param("rekrutmen_id")
	pendaftarID := c.Param("pendaftar_id")

	var req dtos.UpdateApplyStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.rekrutmenService.RefreshApplyStatus(c.Request.Context(), rekrutmenID, pendaftarID, req.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "status pendaftar berhasil diperbarui"})
}

func (h *rekrutmenHandler) AcceptPendaftar(c *gin.Context) {
	rekrutmenID := c.Param("rekrutmen_id")
	pendaftarID := c.Param("pendaftar_id")

	userID := c.MustGet("user_id").(string)

	if err := h.rekrutmenService.AcceptPendaftar(c.Request.Context(), userID, rekrutmenID, pendaftarID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "status pendaftar berhasil diubah menjadi approved dan ditambahkan ke tim"})
}

func (h *rekrutmenHandler) RejectPendaftar(c *gin.Context) {
	rekrutmenID := c.Param("rekrutmen_id")
	pendaftarID := c.Param("pendaftar_id")

	userID := c.MustGet("user_id").(string)

	if err := h.rekrutmenService.RejectPendaftar(c.Request.Context(), userID, rekrutmenID, pendaftarID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "status pendaftar berhasil diubah menjadi rejected"})
}

func (h *rekrutmenHandler) GetApplicantDetail(c *gin.Context) {
	rekrutmenID := c.Param("id")
	pendaftarID := c.Param("pendaftar_id")

	result, err := h.rekrutmenService.GetApplicantDetail(c.Request.Context(), rekrutmenID, pendaftarID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *rekrutmenHandler) GetTeamMembers(c *gin.Context) {
	timID := c.Param("id")

	result, err := h.rekrutmenService.GetTeamMembers(c.Request.Context(), timID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *rekrutmenHandler) GiveMemberRating(c *gin.Context) {
	reviewerUserID := c.MustGet("user_id").(string)
	timID := c.Param("id")
	targetUserID := c.Param("user_id")

	var req dtos.GiveMemberRatingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.rekrutmenService.GiveMemberRating(c.Request.Context(), reviewerUserID, timID, targetUserID, req.Rating, req.Deskripsi); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "rating berhasil diberikan"})
}
