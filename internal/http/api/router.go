package api

import (
	"github.com/hhhhkkk/mini-blog/router"
)

func NewApiRouterProvider() router.IRouterGroup {
	apiRG := router.NewRouterGroup("api")

	apiRG.AddRouter(router.NewGetRouter("/health", Health))

	return apiRG
}
