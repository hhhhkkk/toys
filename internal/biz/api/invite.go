package api

import (
	"github.com/hhhhkkk/mini-blog/internal/biz/api/entity"
	"github.com/hhhhkkk/mini-blog/internal/biz/repository/api"
)

type InviteBiz struct {
	inviteRepo      api.InviteRepo
	inviteAwardRepo api.InviteAwardRepo
}

func NewInviteBiz(inviteRepo api.InviteRepo) *InviteBiz {
	return &InviteBiz{
		inviteRepo: inviteRepo,
	}
}

func (i *InviteBiz) AwardInvite(en *entity.InviteEntity) error {
	en = i.inviteRepo.GetByInvitedUid(en.InvitedUid)
	if en == nil {
		return nil
	}
	inviteTarget := entity.AwardTargetInviter
	invitedTarget := entity.AwardTargetInvitee
	// 发送被邀请人
	// 发送邀请人
	inviteTotal := i.inviteRepo.GetInviteTotal(en.InviteUid)
	i.inviteAwardRepo.SendInviteAward(en.InviteUid, inviteTarget.GetAwardLevel(inviteTotal))
	i.inviteAwardRepo.SendInviteAward(en.InvitedUid, invitedTarget.GetAwardLevel(0))
	return nil
}
