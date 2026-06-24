package app

import (
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/v1/internal/app/admin"
	"github.com/hhhhkkk/mini-blog/v1/internal/app/api"
)

var ProviderSet = wire.NewSet(
	admin.NewUserCase,
	api.NewUserCase,
)
