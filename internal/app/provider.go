package app

import (
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/internal/app/admin/user"
)

var ProviderSet = wire.NewSet(
	user.NewUserCase,
)
