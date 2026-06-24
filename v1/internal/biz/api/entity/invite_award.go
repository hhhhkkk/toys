package entity

import "time"

type InviteAwardEntity struct {
	Id        int
	InviteUid uint
	AwardUid  uint
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
} 
