package models

import "time"

type BookBorrowing struct {
	BorrowingID string     `gorm:"column:borrowing_id;primaryKey"`
	UserID      string     `gorm:"column:user_id;index"`
	BookID      string     `gorm:"column:book_id;index"`
	BorrowedAt  time.Time  `gorm:"column:borrowed_at"`
	ReturnedAt  *time.Time `gorm:"column:returned_at"` // pointer;nullable
	User        User       `gorm:"foreignKey:UserID;references:UserID"`
	Book        Book       `gorm:"foreignKey:BookID;references:BookID"`
}