package models

type TimParticipant struct {
	ID       string `gorm:"column:id;primaryKey"`
	TimID    string `gorm:"column:tim_id"`
	UserID   string `gorm:"column:user_id"`
	MemberKe int64  `gorm:"column:member_ke"`

	Tim  Tim  `gorm:"foreignKey:TimID;references:TimID"`
	User User `gorm:"foreignKey:UserID;references:UserID"`
}
