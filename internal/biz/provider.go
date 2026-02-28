package biz

import (
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/internal/biz/admin"
)

var ProviderSet = wire.NewSet(admin.NewUserBiz)
