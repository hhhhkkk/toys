package http

import (
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/internal/http/admin"
	"github.com/hhhhkkk/mini-blog/internal/http/api"
	"github.com/hhhhkkk/mini-blog/router"
)

func NewRouterProvider() []router.IRouterGroup {
	var r []router.IRouterGroup
	r = append(r, api.NewApiRouterProvider())
	r = append(r, admin.NewAdminRouterProvider())
	// other
	return r
}

var RouterProviderSet = wire.NewSet(NewRouterProvider)
