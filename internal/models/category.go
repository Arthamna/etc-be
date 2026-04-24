package models

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	CategoryID  string         `gorm:"column:category_id;primaryKey"`
	Name        string         `gorm:"column:name;unique"`
	Description string         `gorm:"column:description"`
	CreatedAt   time.Time      `gorm:"column:created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at"`
	
	Books []Book `gorm:"many2many:book_categories;joinForeignKey:CategoryID;JoinReferences:BookID"`
}
