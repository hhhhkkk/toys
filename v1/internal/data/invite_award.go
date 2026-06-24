package data

import (
	"time"

	"github.com/hhhhkkk/mini-blog/v1/internal/biz/api/entity"
)

type UserInviteAward struct {
	Id        int
	Uid       int
	Award     string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type InviteAwardRepoImpl struct {
	db *DB
}

func NewInviteAwardRepoImpl(db *DB) *InviteAwardRepoImpl {
	return &InviteAwardRepoImpl{
		db: db,
	}
}

func (impl *InviteAwardRepoImpl) SendInviteAward(uid int, award entity.AwardLevel) error {
	userInviteAward := &UserInviteAward{
		Uid:   uid,
		Award: award.String(),
	}
	if err := impl.db.GetClient().Create(userInviteAward).Error; err != nil {
		return err
	}
	return nil
}
