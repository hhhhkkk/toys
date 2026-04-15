package entity

import "time"

type InviteEntity struct {
	Id         int
	InviteUid  int
	InvitedUid int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
