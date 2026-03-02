package app

import (
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/internal/app/admin"
)

var ProviderSet = wire.NewSet(
	admin.NewUserCase,
)
