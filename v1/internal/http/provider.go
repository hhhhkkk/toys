package http

import (
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/v1/internal/http/admin"
	"github.com/hhhhkkk/mini-blog/v1/internal/http/api"
	"github.com/hhhhkkk/mini-blog/v1/router"
)

func NewRouterProvider(uc *admin.Controller, lc *admin.LoginController, apiUser *api.UserController) []router.IRouterGroup {
	var r []router.IRouterGroup
	r = append(r, api.NewApiRouterProvider(apiUser))
	r = append(r, admin.NewAdminRouterProvider(uc, lc))
	// other
	return r
}

var RouterProviderSet = wire.NewSet(admin.ProviderSet, api.ProviderSet, NewRouterProvider)
