package models

import (
	"gorm.io/gorm"
	"time"
)

type Book struct {
	BookID      string         `gorm:"column:book_id;primaryKey"`
	Author    string         `gorm:"column:author"`
	BookPicture []byte         `gorm:"column:book_picture"`
	Title       string         `gorm:"column:title;index"`
	Description string         `gorm:"column:description"`
	Status      string         `gorm:"column:status;default:'available'"` 
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
	
	Categories []Category `gorm:"many2many:book_categories;joinForeignKey:BookID;JoinReferences:CategoryID"`
	Borrowings  []BookBorrowing `gorm:"foreignKey:BookID;references:BookID"`
}