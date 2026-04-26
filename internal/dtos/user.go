// internal/dtos/user.go
package dtos

import (
	"etc-backend/internal/models"
	"time"
)

type UserRegisterRequest struct {
	Nama          string `json:"nama" binding:"required"`
	Jurusan       string `json:"jurusan" binding:"required"`
	NRP           string `json:"nrp" binding:"required"`
	ContactPerson string `json:"contact_person" binding:"required"`
	Password      string `json:"password" binding:"required,min=6"`
}

type UserLoginRequest struct {
	NRP      string `json:"nrp" binding:"required"`
	Password string `json:"password" binding:"required"`
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
	UserID        string    `json:"user_id"`
	Nama          string    `json:"nama"`
	Jurusan       string    `json:"jurusan"`
	NRP           string    `json:"nrp"`
	ContactPerson string    `json:"contact_person"`
	Role          string    `json:"role"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func ToUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		UserID:        user.UserID,
		Nama:          user.Nama,
		Jurusan:       user.Jurusan,
		NRP:           user.NRP,
		ContactPerson: user.ContactPerson,
		Role:          user.Role,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
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
	Nama          string `json:"nama" binding:"required"`
	Jurusan       string `json:"jurusan" binding:"required"`
	NRP           string `json:"nrp" binding:"required"`
	ContactPerson string `json:"contact_person" binding:"required"`
	Password      string `json:"password" binding:"required,min=6"`
	SecretKey     string `json:"secret_key" binding:"required"`
}

type UpdateUserRequest struct {
	Nama          string `json:"nama"`
	Jurusan       string `json:"jurusan"`
	ContactPerson string `json:"contact_person"`
	Role          string `json:"role"` // "mahasiswa" atau "dosen"
}
