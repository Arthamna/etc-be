package models

import (
	"time"

	"github.com/google/uuid"
)

type Pendaftar struct {
	PendaftarID     uuid.UUID  `gorm:"type:uuid;column:pendaftar_id;primaryKey;default:uuid_generate_v4()" json:"pendaftar_id"`
	RekrutmenID     uuid.UUID  `gorm:"type:uuid;column:rekrutmen_id" json:"rekrutmen_id"`
	UserID          uuid.UUID  `gorm:"type:uuid;column:user_id" json:"user_id"`
	AlasanMendaftar *string    `gorm:"column:alasan_mendaftar" json:"alasan_mendaftar"`
	CVURL           string     `gorm:"column:cv_url" json:"cv_url"`
	PortofolioURL   string     `gorm:"column:portofolio_url" json:"portofolio_url"`
	Status          string     `gorm:"column:status" json:"status"`
	CreatedAt       *time.Time `gorm:"column:created_at" json:"created_at"`

	Rekrutmen Rekrutmen `gorm:"foreignKey:RekrutmenID;references:RekrutmenID"`
	User      User      `gorm:"foreignKey:UserID;references:UserID"`
}
