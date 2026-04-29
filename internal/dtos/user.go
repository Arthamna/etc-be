// internal/dtos/user.go
package dtos

import (
	"mime/multipart"
	// "time"
)

type UserRegisterRequest struct {
	Nama string `json:"nama" binding:"required"`
	// mahasiswa only
	Jurusan *string `json:"jurusan"`

	Role         string   `json:"role" binding:"required,oneof=mahasiswa dosen"`
	Password     string   `json:"password" binding:"required,min=6"`
	NoTelp       string   `json:"no_telp" binding:"required"`
	NoPengenal   string   `json:"no_pengenal"`
	Spesialisasi []string `json:"spesialisasi"`
}

type UserLoginRequest struct {
	NoPengenal string `json:"no_pengenal"`
	Password   string `json:"password" binding:"required"`
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
	UserID         string   `json:"user_id"`
	Nama           string   `json:"nama"`
	Jurusan        *string  `no_telpjson:"jurusan"`
	NoPengenal     string   `json:"no_pengenal"`
	NoTelp         string   `json:"no_telp"`
	Role           string   `json:"role"`
	ProfilePicture *string  `json:"profile_picture"`
	Spesialisasi   []string `json:"spesialisasi"`
	// CreatedAt     time.Time `json:"created_at"`
	// UpdatedAt     time.Time `json:"updated_at"`
}

type UserGetMe struct {
	PersonalInfo UserResponse `json:"personal_info"`
}

type UploadProfilePictureRequest struct {
	ProfilePicture *multipart.FileHeader `form:"profile_picture" binding:"required"`
}

type UpdateProfilePictureResponse struct {
	ProfilePicture *string `json:"profile_picture"`
}

type UpdateUserRequest struct {
	Nama         string `json:"nama"`
	Jurusan      string `json:"jurusan"`
	NoTelp       string `json:"no_telp"`
	Spesialisasi []string `json:"spesialisasi"`
	NoPengenal   string `json:"no_pengenal"`
	// Role          string `json:"role"` // "mahasiswa" atau "dosen"
}
