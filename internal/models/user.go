package models

import (
	"time"
)

type User struct {
	UserID         string    `gorm:"column:user_id;primaryKey"`
	Nama           string    `gorm:"column:nama"`
	Jurusan        string    `gorm:"column:jurusan"`
	NRP            string    `gorm:"column:nrp;unique"`
	Role           string    `gorm:"column:role"`
	ContactPerson  string    `gorm:"column:contact_person"`
	PasswordHash   string    `gorm:"column:password_hash"`
	ProfilePicture *string   `gorm:"column:profile_picture"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string { return "users" }
