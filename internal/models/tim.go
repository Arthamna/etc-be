package models

import "time"

type Tim struct {
	TimID        string    `gorm:"column:tim_id;primaryKey"`
	TipeTim      string    `gorm:"column:tipe_tim"` // enum
	RekrutmenID  string    `gorm:"column:rekrutmen_id"`
	NamaKetua    string    `gorm:"column:nama_ketua"`
	CreatedAt    time.Time `gorm:"column:created_at"`

	Rekrutmen    Rekrutmen        `gorm:"foreignKey:RekrutmenID;references:RekrutmenID"`
	Participants []TimParticipant `gorm:"foreignKey:TimID;references:TimID"`
	Histories    []History        `gorm:"foreignKey:TimID;references:TimID"`
}
