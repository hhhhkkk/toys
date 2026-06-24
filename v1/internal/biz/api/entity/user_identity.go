package entity

import "time"

type UserIdentityEntity struct {
	Id        int       `json:"id"`
	Uid       int       `json:"uid"`
	Channel   int       `json:"channel"`
	Identity  string    `json:"identity"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
