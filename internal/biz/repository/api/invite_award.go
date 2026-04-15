package api

import "github.com/hhhhkkk/mini-blog/internal/biz/api/entity"

type InviteAwardRepo interface {
	SendInviteAward(inviteUid int, award entity.AwardLevel) error
}
