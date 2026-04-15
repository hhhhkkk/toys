package data

import (
	"time"

	"github.com/hhhhkkk/mini-blog/internal/biz/api/entity"
)

type UserIdentity struct {
	Id        int       `gorm:"primaryKey"`
	Uid       int       `gorm:"type:int(11);not null"`
	Channel   int       `gorm:"type:tinyint;not null"`
	Identity  string    `gorm:"type:varchar(255);not null"`
	Status    int       `gorm:"type:tinyint;not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type IdentityRepoImpl struct {
	db *DB
}

func NewIdentityRepoImpl(db *DB) *IdentityRepoImpl {
	return &IdentityRepoImpl{
		db: db,
	}
}

func (impl *IdentityRepoImpl) Add(e *entity.UserIdentityEntity) (*entity.UserIdentityEntity, error) {
	userIdentity := &UserIdentity{
		Uid:      e.Uid,
		Channel:  e.Channel,
		Identity: e.Identity,
		Status:   e.Status,
	}
	if err := impl.db.GetClient().Create(userIdentity).Error; err != nil {
		return nil, err
	}
	e.Id = userIdentity.Id
	e.CreatedAt = userIdentity.CreatedAt
	e.UpdatedAt = userIdentity.UpdatedAt
	return e, nil
}

func (impl *IdentityRepoImpl) Get(e *entity.UserIdentityEntity) *entity.UserIdentityEntity {
	if e.Channel == 0 || e.Identity == "" {
		return nil
	}

	userIdentity := &UserIdentity{}
	tx := impl.db.GetClient().Model(&UserIdentity{}).
		Where("channel = ?", e.Channel).
		Where("identity = ?", e.Identity).
		First(userIdentity)

	if tx.Error != nil {
		return nil
	}
	return &entity.UserIdentityEntity{
		Id:        userIdentity.Id,
		Uid:       userIdentity.Uid,
		Channel:   userIdentity.Channel,
		Identity:  userIdentity.Identity,
		Status:    userIdentity.Status,
		CreatedAt: userIdentity.CreatedAt,
		UpdatedAt: userIdentity.UpdatedAt,
	}
}
