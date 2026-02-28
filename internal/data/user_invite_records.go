package data

import "time"

type UserInviteRecord struct {
	Id        int       `gorm:"primaryKey"`
	UserId    int       `gorm:"type:int(11);not null"`
	InviteId  int       `gorm:"type:int(11);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

func (u *UserInviteRecord) TableName() string {
	return "user_invite_records"
}
