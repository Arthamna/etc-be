package models

import (
	"time"
)

type User struct {
	UserID         string    `gorm:"column:user_id;primaryKey"`
	Nama           string    `gorm:"column:nama"`
	Jurusan        *string   `gorm:"column:jurusan"`
	NoPengenal     string   `gorm:"column:no_pengenal;unique"`
	Role           string    `gorm:"column:role"` // dosen atau mahasiswa
	NoTelp         string    `gorm:"column:no_telp"`
	PasswordHash   string    `gorm:"column:password_hash"`
	ProfilePicture *string   `gorm:"column:profile_picture"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
	Spesialisasi   []string   `json:"spesialisasi" gorm:"column:spesialisasi;type:json;serializer:json"`
	
	BookmarksAsUser      []Bookmark      `gorm:"foreignKey:UserID;references:UserID"`
	Pendaftar            []Pendaftar     `gorm:"foreignKey:UserID;references:UserID"`
	HistoryAsUser        []History       `gorm:"foreignKey:UserID;references:UserID"`
	HistoryAsReviewer    []History       `gorm:"foreignKey:ReviewerUserID;references:UserID"`
	TimParticipants      []TimParticipant `gorm:"foreignKey:UserID;references:UserID"`
}

// NRP            *string   `gorm:"column:nrp;unique"`
// NIDN           *string   `gorm:"column:nidn;unique"`
func (User) TableName() string { return "users" }
