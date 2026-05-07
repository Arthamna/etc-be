package models

import (
	"time"

	"github.com/google/uuid"
)

type History struct {
	ID             uuid.UUID `gorm:"type:uuid;column:id;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserID         uuid.UUID `gorm:"type:uuid;column:user_id" json:"user_id"`
	ReviewerUserID uuid.UUID `gorm:"type:uuid;column:reviewer_user_id" json:"reviewer_user_id"`
	TimID          uuid.UUID `gorm:"type:uuid;column:tim_id" json:"tim_id"`
	Rating         int64     `gorm:"column:rating" json:"rating"`
	Deskripsi      string    `gorm:"column:deskripsi" json:"deskripsi"`
	CreatedAt      time.Time `gorm:"column:created_at" json:"created_at"`

	User     User `gorm:"foreignKey:UserID;references:UserID"`
	Reviewer User `gorm:"foreignKey:ReviewerUserID;references:UserID"`
	Tim      Tim  `gorm:"foreignKey:TimID;references:TimID"`
}