package biz

import (
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/v1/internal/biz/admin"
	"github.com/hhhhkkk/mini-blog/v1/internal/biz/api"
)

var ProviderSet = wire.NewSet(admin.NewUserBiz, api.NewUserBiz, api.NewInviteBiz)
