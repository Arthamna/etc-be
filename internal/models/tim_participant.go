package models

import "github.com/google/uuid"

type TimParticipant struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	TimID    uuid.UUID `gorm:"type:uuid;column:tim_id" json:"tim_id"`
	UserID   uuid.UUID `gorm:"type:uuid;column:user_id" json:"user_id"`
	MemberKe int64     `gorm:"column:member_ke" json:"member_ke"`

	Tim  Tim  `gorm:"foreignKey:TimID;references:TimID"`
	User User `gorm:"foreignKey:UserID;references:UserID"`
}