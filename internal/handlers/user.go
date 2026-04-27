package handlers

import (
	"etc-backend/internal/dtos"
	"etc-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	UserHandler interface {
		Register(c *gin.Context)
		Login(c *gin.Context)
		// RegisterAdmin(c *gin.Context)
		UpdateUser(c *gin.Context)
		GetMe(c *gin.Context)
		UploadPicture(c *gin.Context)
	}

	userHandler struct {
		userService services.UserService
	}
)

func NewUserHandler(userService services.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) Register(c *gin.Context) {
	var req dtos.UserRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Register(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *userHandler) Login(c *gin.Context) {
	var req dtos.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *userHandler) GetMe(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	result, err := h.userService.GetByID(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *userHandler) UploadPicture(c *gin.Context) {
	var req dtos.UploadProfilePictureRequest
	userId := c.MustGet("user_id").(string)
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.UploadProfilePicture(c.Request.Context(), req, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}


// func (h *userHandler) RegisterAdmin(c *gin.Context) {
// 	var req dtos.AdminRegisterRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	user, err := h.userService.RegisterAdmin(c.Request.Context(), req)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, user)
// }

func (h *userHandler) UpdateUser(c *gin.Context) {
	userID := c.MustGet("user_id").(string)

	var req dtos.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.UpdateUser(c.Request.Context(), userID, req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
