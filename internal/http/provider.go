package http

import (
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/internal/http/admin"
	"github.com/hhhhkkk/mini-blog/internal/http/api"
	"github.com/hhhhkkk/mini-blog/router"
)

func NewRouterProvider(uc *admin.Controller, lc *admin.LoginController) []router.IRouterGroup {
	var r []router.IRouterGroup
	r = append(r, api.NewApiRouterProvider())
	r = append(r, admin.NewAdminRouterProvider(uc, lc))
	// other
	return r
}

var RouterProviderSet = wire.NewSet(admin.ProviderSet, NewRouterProvider)
