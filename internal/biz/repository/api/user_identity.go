package api

import "github.com/hhhhkkk/mini-blog/internal/biz/api/entity"

type UserIdentityRepo interface {
	Add(e *entity.UserIdentityEntity) (*entity.UserIdentityEntity, error)
	Get(e *entity.UserIdentityEntity) *entity.UserIdentityEntity
}
