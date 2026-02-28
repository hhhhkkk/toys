package data

import "time"

type User struct {
	Id           int       `gorm:"primaryKey"`
	Name         string    `gorm:"type:varchar(255);not null"`
	Email        string    `gorm:"type:varchar(255);not null"`
	Password     string    `gorm:"type:varchar(255);not null"`
	Age          int       `gorm:"type:int(11);not null;"`
	InviteCode   string    `gorm:"type:varchar(50);not null"`
	Status       int       `gorm:"type:int(11);not null"`
	CreatedAt    time.Time `gorm:"type:datetime;not null"`
	UpdatedAt    time.Time `gorm:"type:datetime;not null"`
	LastLoginAt  time.Time `gorm:"type:datetime;not null"`
	LastActiveAt time.Time `gorm:"type:datetime;not null"`
}

func (u *User) TableName() string {
	return "users"
}
