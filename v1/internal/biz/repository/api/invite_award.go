package api

import "github.com/hhhhkkk/mini-blog/v1/internal/biz/api/entity"

type InviteAwardRepo interface {
	SendInviteAward(inviteUid int, award entity.AwardLevel) error
}
