package models

type Bookmark struct {
	ID          string `gorm:"column:id;primaryKey"`
	RekrutmenID string `gorm:"column:rekrutmen_id"`
	UserID      string `gorm:"column:user_id"`

	Rekrutmen Rekrutmen `gorm:"foreignKey:RekrutmenID;references:RekrutmenID"`
	User      User      `gorm:"foreignKey:UserID;references:UserID"`
}
