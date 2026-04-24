// internal/dtos/user.go
package dtos

import (
	"arthamna/rplLibrary/internal/models"
	"encoding/base64"
	"mime/multipart"
	"time"
)

type UserRegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UploadProfilePictureRequest struct {
	ProfilePicture *multipart.FileHeader `form:"profile_picture" binding:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type UserUpdateRequest struct {
	Username string `json:"username" binding:""`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:""`
}

type UserRegisterResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}

type UserResponse struct {
	UserID           string    `json:"user_id"`
	Username         string    `json:"username"`
	Email            string    `json:"email"`
	ProfilePicture   string    `json:"profile_picture"`
	Role             string    `json:"role"`
	RegistrationDate time.Time `json:"registration_date"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type UpdateProfilePictureResponse struct {
	ProfilePicture string `json:"profile_picture"`
}

func ToUserResponse(user *models.User) *UserResponse {
	profilePicture := ""
	if len(user.ProfilePicture) > 0 {
		profilePicture = base64.StdEncoding.EncodeToString(user.ProfilePicture)
	}

	return &UserResponse{
		UserID:           user.UserID,
		Username:         user.Username,
		Email:            user.Email,
		ProfilePicture:   profilePicture,
		Role:             user.Role,
		RegistrationDate: user.RegistrationDate,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
	}
}

func ToUserResponseList(users []models.User) []*UserResponse {
	var responses []*UserResponse
	for _, user := range users {
		responses = append(responses, ToUserResponse(&user))
	}
	return responses
}

type AdminRegisterRequest struct {
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	SecretKey string `json:"secret_key" binding:"required"`
}
