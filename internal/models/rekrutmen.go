package models

import "time"

type Rekrutmen struct {
	RekrutmenID    string    `gorm:"column:rekrutmen_id;primaryKey"`
	UserID         string    `gorm:"column:user_id"`
	Kegiatan       string    `gorm:"column:kegiatan"` // projek, riset, lomba
	Kriteria	   string    `gorm:"column:kriteria"`
	TanggalMulai   time.Time `gorm:"column:tanggal_mulai"`
	TanggalSelesai time.Time `gorm:"column:tanggal_selesai"`
	Fee            *float64   `gorm:"column:fee"`
	Role           string    `gorm:"column:role"` // role di kegiatan tersebut
	ContactPerson  string    `gorm:"column:contact_person"`

	User      User        `gorm:"foreignKey:UserID;references:UserID"`
	Pendaftar []Pendaftar `gorm:"foreignKey:RekrutmenID;references:RekrutmenID"`
	Bookmarks []Bookmark  `gorm:"foreignKey:RekrutmenID;references:RekrutmenID"`
	Tims      []Tim       `gorm:"foreignKey:RekrutmenID;references:RekrutmenID"`
}
