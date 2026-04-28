package models

import "time"

type Pendaftar struct {
	PendaftarID     string     `gorm:"column:pendaftar_id;primaryKey"`
	RekrutmenID     string     `gorm:"column:rekrutmen_id"`
	UserID          string     `gorm:"column:user_id"`
	AlasanMendaftar *string    `gorm:"column:alasan_mendaftar"`
	CVURL           string     `gorm:"column:cv_url"`
	PortofolioURL   string     `gorm:"column:portofolio_url"`
	Status          string     `gorm:"column:status"` // diterima atau tidak
	CreatedAt       *time.Time `gorm:"column:created_at"`

	Rekrutmen Rekrutmen `gorm:"foreignKey:RekrutmenID;references:RekrutmenID"`
	User      User      `gorm:"foreignKey:UserID;references:UserID"`
}
