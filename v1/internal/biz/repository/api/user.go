package api

import "github.com/hhhhkkk/mini-blog/v1/internal/biz/api/entity"

type UserRepo interface {
	EmailExists(email string, id int) bool
	GetUserByInviteCode(code string) (*entity.UserEntity, error)
	CreateUser(entity *entity.UserEntity) (*entity.UserEntity, error)
}
