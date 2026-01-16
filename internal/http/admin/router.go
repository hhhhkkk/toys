package admin

import (
	"github.com/hhhhkkk/mini-blog/router"
)

func NewAdminRouterProvider() router.IRouterGroup {
	adminRG := router.NewRouterGroup("admin")

	adminRG.AddRouter(router.NewGetRouter("/health", AdminHealth))
	adminRG.AddRouter(router.NewGetRouter("/demoPanic", DemoPanic))

	return adminRG
}
