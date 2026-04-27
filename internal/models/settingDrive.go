package models

import "github.com/google/uuid"

type SettingDrive struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`

	TypeID  string `json:"type_id"`
	Month    string `json:"month"`
	Day      int    `json:"day"`
	MonthUrl string `json:"month_url"` // original folder MonthUrl
	DayUrl   string `json:"day_url"`   // original folder DayUrl
}
