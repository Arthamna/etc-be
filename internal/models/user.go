package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID           string         `gorm:"column:user_id;primaryKey"`
	Username         string         `gorm:"column:username"`
	Email            string         `gorm:"column:email;unique"`
	PasswordHash     string         `gorm:"column:password_hash"`
	ProfilePicture   []byte         `gorm:"column:profile_picture"`
	RegistrationDate time.Time      `gorm:"column:registration_date"`
	Role             string         `gorm:"column:role;"` 
	CreatedAt        time.Time      `gorm:"column:created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at"`
	BookBorrowings   []BookBorrowing `gorm:"foreignKey:UserID;references:UserID"`
}