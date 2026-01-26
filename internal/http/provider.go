package http

import (
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/internal/data"
	"github.com/hhhhkkk/mini-blog/internal/http/admin"
	"github.com/hhhhkkk/mini-blog/internal/http/admin/user"
	"github.com/hhhhkkk/mini-blog/internal/http/api"
	"github.com/hhhhkkk/mini-blog/router"
)

func NewRouterProvider(uc *user.UserController, cache *data.Cache) []router.IRouterGroup {
	var r []router.IRouterGroup
	r = append(r, api.NewApiRouterProvider())
	r = append(r, admin.NewAdminRouterProvider(uc, cache))
	// other
	return r
}

var RouterProviderSet = wire.NewSet(admin.ProviderSet, NewRouterProvider)
