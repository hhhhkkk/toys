package entity

import "time"

type UserEntity struct {
	Id           int
	Name         string
	Email        string
	Age          int
	InviteCode   string
	Status       int
	LastLoginAt  time.Time
	LastActiveAt time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
