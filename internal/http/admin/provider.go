package admin

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserController, NewLoginController, NewAdminRouterProvider)
