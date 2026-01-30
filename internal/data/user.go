package data

import "time"

type User struct {
	Id        int       `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Age       int       `gorm:"type:int(11);not null;"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

func (u *User) TableName() string {
	return "aaa"
}
