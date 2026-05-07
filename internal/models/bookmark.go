package models

import "github.com/google/uuid"

type Bookmark struct {
	ID          uuid.UUID `gorm:"type:uuid;column:id;primaryKey;default:uuid_generate_v4()" json:"id"`
	RekrutmenID uuid.UUID `gorm:"type:uuid;column:rekrutmen_id" json:"rekrutmen_id"`
	UserID      uuid.UUID `gorm:"type:uuid;column:user_id" json:"user_id"`

	Rekrutmen Rekrutmen `gorm:"foreignKey:RekrutmenID;references:RekrutmenID"`
	User      User      `gorm:"foreignKey:UserID;references:UserID"`
}
