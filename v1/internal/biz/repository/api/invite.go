package api

import "github.com/hhhhkkk/mini-blog/v1/internal/biz/api/entity"

type InviteRepo interface {
	CreateInvite(invite *entity.InviteEntity) (*entity.InviteEntity, error)
	GetByInvitedUid(uid int) *entity.InviteEntity
	GetByInviteUid(uid int) []*entity.InviteEntity
	GetInviteTotal(uid int) int
}
