package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID         uuid.UUID `gorm:"type:uuid;column:user_id;primaryKey;default:uuid_generate_v4()" json:"user_id"`
	Nama           string    `gorm:"column:nama" json:"nama"`
	Jurusan        *string   `gorm:"column:jurusan" json:"jurusan"`
	NoPengenal     string    `gorm:"column:no_pengenal;unique" json:"no_pengenal"`
	Role           string    `gorm:"column:role" json:"role"`
	NoTelp         string    `gorm:"column:no_telp" json:"no_telp"`
	PasswordHash   string    `gorm:"column:password_hash" json:"-"`
	ProfilePicture *string   `gorm:"column:profile_picture" json:"profile_picture"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at" json:"updated_at"`
	Spesialisasi   []string  `json:"spesialisasi" gorm:"column:spesialisasi;type:json;serializer:json"`

	BookmarksAsUser   []Bookmark       `gorm:"foreignKey:UserID;references:UserID"`
	Pendaftar         []Pendaftar      `gorm:"foreignKey:UserID;references:UserID"`
	HistoryAsUser     []History        `gorm:"foreignKey:UserID;references:UserID"`
	HistoryAsReviewer  []History        `gorm:"foreignKey:ReviewerUserID;references:UserID"`
	TimParticipants   []TimParticipant `gorm:"foreignKey:UserID;references:UserID"`
}

