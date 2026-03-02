package admin

import (
	"errors"

	"github.com/hhhhkkk/mini-blog/internal/biz/admin/entity"
	repo "github.com/hhhhkkk/mini-blog/internal/biz/repository/admin"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserBiz struct {
	logger *zap.Logger
	repo   repo.Repo
}

func NewUserBiz(logger *zap.Logger, userRepo repo.Repo) *UserBiz {
	return &UserBiz{
		logger: logger,
		repo:   userRepo,
	}
}

func (u *UserBiz) EmailExists(email string, uid uint) bool {
	return u.repo.EmailExist(email, uid)
}

func (u *UserBiz) CreateUser(email, password string) (*entity.User, error) {
	if u.EmailExists(email, 0) {
		return nil, errors.New("email already exists")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	en := entity.NewUser(0, email, string(hashedPassword), 0)
	if _, err := u.repo.CreateUser(en); err != nil {
		return nil, err
	}
	return en, nil
}
