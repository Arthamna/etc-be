package models

import "time"

type History struct {
	ID              string    `gorm:"column:id;primaryKey"`
	UserID          string    `gorm:"column:user_id"`
	ReviewerUserID  string    `gorm:"column:reviewer_user_id"`
	TimID           string    `gorm:"column:tim_id"`
	Rating          int64     `gorm:"column:rating"`
	Deskripsi       string    `gorm:"column:deskripsi"`
	CreatedAt       time.Time `gorm:"column:created_at"`

	User     User `gorm:"foreignKey:UserID;references:UserID"`
	Reviewer User `gorm:"foreignKey:ReviewerUserID;references:UserID"`
	Tim      Tim  `gorm:"foreignKey:TimID;references:TimID"`
}