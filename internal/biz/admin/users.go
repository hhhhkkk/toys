package admin

import (
	"context"
	"errors"

	"github.com/hhhhkkk/mini-blog/internal/biz/admin/entity"
	repo "github.com/hhhhkkk/mini-blog/internal/biz/repository/admin"
	"github.com/hhhhkkk/mini-blog/internal/data"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserBiz struct {
	db     *data.DB
	cache  *data.Cache
	logger *zap.Logger
	repo   repo.Repo
}

func NewUserBiz(db *data.DB, cache *data.Cache, logger *zap.Logger, userRepo repo.Repo) *UserBiz {
	return &UserBiz{
		db:     db,
		cache:  cache,
		logger: logger,
		repo:   userRepo,
	}
}

func (u *UserBiz) EmailExists(ctx context.Context, email string, uid uint) (bool, error) {
	query := gorm.G[data.User](u.db.GetClient()).Where("email = ?", email)
	if uid > 0 {
		query = query.Where("uid != ?", uid)
	}
	_, err := query.Where("uid != ?", uid).First(ctx)
	return err == nil, err
}

func (u *UserBiz) CreateUser(email, password, repeatPassword string) (*entity.User, error) {
	if password != repeatPassword {
		return nil, errors.New("两次输入不一致")
	}
	en := entity.NewUser(0, email, password)
	if _, err := u.repo.CreateUser(en); err != nil {
		return nil, err
	}
	return en, nil
}
