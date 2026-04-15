package biz

import (
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/internal/biz/admin"
	"github.com/hhhhkkk/mini-blog/internal/biz/api"
)

var ProviderSet = wire.NewSet(admin.NewUserBiz, api.NewUserBiz, api.NewInviteBiz)
