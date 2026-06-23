package api

import (
	"github.com/hhhhkkk/mini-blog/internal/biz/api/entity"
	"github.com/hhhhkkk/mini-blog/internal/biz/repository/api"
)

type UserBiz struct {
	userRepo         api.UserRepo
	userIdentityRepo api.UserIdentityRepo
	inviteRepo       api.InviteRepo
}

func NewUserBiz(userRepo api.UserRepo, userIdentityRepo api.UserIdentityRepo, inviteRepo api.InviteRepo) *UserBiz {
	return &UserBiz{
		userRepo:         userRepo,
		userIdentityRepo: userIdentityRepo,
		inviteRepo:       inviteRepo,
	}
}

func (u *UserBiz) EmailExists(email string, id int) bool {
	return u.userRepo.EmailExists(email, id)
}

func (u *UserBiz) CreateUserByEmail(en *entity.UserEntity) (*entity.UserEntity, error) {
	// 创建用户
	user, err := u.userRepo.CreateUser(en)
	if err != nil {
		return nil, err
	}

	// 创建用户渠道
	u.createUserIdentity(user, entity.Email)

	// 创建拉新记录
	if en.InviteCode != "" {
		inviteUser, err := u.userRepo.GetUserByInviteCode(en.InviteCode)
		if err != nil {
			// todo log
			return en, nil
		}
		_, err = u.inviteRepo.CreateInvite(&entity.InviteEntity{
			InvitedUid: user.Id,
			InviteUid:  inviteUser.Id,
		})
		// if err != nil {
		// }
	}
	return user, nil
}

func (u *UserBiz) createUserIdentity(en *entity.UserEntity, channels entity.Channels) error {
	// 创建用户身份
	_, err := u.userIdentityRepo.Add(&entity.UserIdentityEntity{
		Uid:      en.Id,
		Channel:  channels.Value(),
		Identity: en.Email,
	})
	return err
}
