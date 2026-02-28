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

	adminRG.AddRouter(router.NewGetRouter("/demoApp", appMiddleware))

	// user
	userRG := adminRG.NewSubGroup("users")
	{
		// userRG.AddRouter(router.NewGetRouter("/:id", uc.GetUserCache))
		userRG.AddRouter(router.NewGetRouter("/db/:id", uc.GetUserDB))
	}
	// adminRG.AddRouter(router.NewGetRouter("/", appMiddleware))

	return adminRG
}
