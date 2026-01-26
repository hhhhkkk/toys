package admin

import (
	"github.com/hhhhkkk/mini-blog/internal/data"
	"github.com/hhhhkkk/mini-blog/internal/http/admin/user"
	"github.com/hhhhkkk/mini-blog/router"
)

func NewAdminRouterProvider(uc *user.UserController, cache *data.Cache) router.IRouterGroup {
	adminRG := router.NewRouterGroup("admin")

	adminRG.AddRouter(router.NewGetRouter("/health", AdminHealth))
	adminRG.AddRouter(router.NewGetRouter("/demoPanic", DemoPanic))

	adminRG.AddRouter(router.NewGetRouter("/demoConfig", configPath))
	adminRG.AddRouter(router.NewGetRouter("/demoApp", appMiddleware))

	// user
	userRG := adminRG.NewSubGroup("users")
	{
		userRG.AddRouter(router.NewGetRouter("/:id", uc.GetUser))
	}
	// adminRG.AddRouter(router.NewGetRouter("/", appMiddleware))

	return adminRG
}
