package admin

import (
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/internal/http/admin/user"
)

var ProviderSet = wire.NewSet(user.NewUserController, NewAdminRouterProvider)
