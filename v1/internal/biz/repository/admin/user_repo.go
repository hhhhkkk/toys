package admin

import "github.com/hhhhkkk/mini-blog/v1/internal/biz/admin/entity"

type Repo interface {
	EmailExist(email string, uid uint) bool
	CreateUser(user *entity.User) (*entity.User, error)
}
