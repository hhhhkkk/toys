package api

import (
	"errors"

	"github.com/gin-gonic/gin"
	biz "github.com/hhhhkkk/mini-blog/internal/biz/api"
	"github.com/hhhhkkk/mini-blog/internal/biz/api/entity"
)

type UserCase struct {
	userBiz   *biz.UserBiz
	inviteBiz *biz.InviteBiz
}

func NewUserCase(biz *biz.UserBiz) *UserCase {
	return &UserCase{
		userBiz: biz,
	}
}

func (u *UserCase) Register(ctx *gin.Context, req *RegisterReq) error {
	if u.userBiz.EmailExists(req.Email, 0) {
		return errors.New("email already exists")
	}

	userEntity := &entity.UserEntity{
		Email:      req.Email,
		InviteCode: req.InviteCode,
	}
	userEntity, err := u.userBiz.CreateUserByEmail(userEntity)
	if err != nil {
		return err
	}
	err = u.inviteBiz.AwardInvite(&entity.InviteEntity{
		InvitedUid: userEntity.Id,
	})
	if err != nil {
		// log
	}
	return nil
}
