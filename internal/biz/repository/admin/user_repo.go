package admin

import "github.com/hhhhkkk/mini-blog/internal/biz/admin/entity"

type Repo interface {
	EmailExist(email string, uid uint) bool
	CreateUser(user *entity.User) (*entity.User, error)
}
