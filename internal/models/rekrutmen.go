package models

import "time"

type Rekrutmen struct {
	RekrutmenID    string    `gorm:"column:rekrutmen_id;primaryKey"`
	UserID         string    `gorm:"column:user_id"`
	KegiatanID     string    `gorm:"column:kegiatan_id"`
	Kegiatan       string    `gorm:"column:kegiatan"` // projek, riset, lomba
	TanggalMulai   time.Time `gorm:"column:tanggal_mulai"`
	TanggalSelesai time.Time `gorm:"column:tanggal_selesai"`
	Fee            float64   `gorm:"column:fee"`
	Role           string    `gorm:"column:role"`
	ContactPerson  string    `gorm:"column:contact_person"`

	User User `gorm:"foreignKey:UserID;references:UserID"`
}

func (Rekrutmen) TableName() string { return "rekrutmen" }
