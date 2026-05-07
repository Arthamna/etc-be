package models

import (
	"time"

	"github.com/google/uuid"
)

type Rekrutmen struct {
	RekrutmenID    uuid.UUID `gorm:"type:uuid;column:rekrutmen_id;primaryKey;default:uuid_generate_v4()" json:"rekrutmen_id"`
	UserID         uuid.UUID `gorm:"type:uuid;column:user_id" json:"user_id"`
	Kegiatan       string    `gorm:"column:kegiatan" json:"kegiatan"`
	Kriteria       string    `gorm:"column:kriteria" json:"kriteria"`
	TanggalMulai   time.Time `gorm:"column:tanggal_mulai" json:"tanggal_mulai"`
	TanggalSelesai time.Time `gorm:"column:tanggal_selesai" json:"tanggal_selesai"`
	Fee            *float64  `gorm:"column:fee" json:"fee"`
	Role           string    `gorm:"column:role" json:"role"`
	ContactPerson  string    `gorm:"column:contact_person" json:"contact_person"`

	User      User       `gorm:"foreignKey:UserID;references:UserID"`
	Pendaftar []Pendaftar `gorm:"foreignKey:RekrutmenID;references:RekrutmenID"`
	Bookmarks []Bookmark  `gorm:"foreignKey:RekrutmenID;references:RekrutmenID"`
	Tims      []Tim       `gorm:"foreignKey:RekrutmenID;references:RekrutmenID"`
}
