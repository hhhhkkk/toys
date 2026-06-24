package admin

import (
	"github.com/hhhhkkk/mini-blog/v1/router"
)

func NewAdminRouterProvider(uc *Controller, lc *LoginController) router.IRouterGroup {
	adminRG := router.NewRouterGroup("admin")

	adminRG.AddRouter(router.NewGetRouter("/health", Health))
	// adminRG.AddRouter(router.NewGetRouter("/demoPanic", DemoPanic))
	// adminRG.AddRouter(router.NewGetRouter("/demoApp", appMiddleware))

	adminRG.AddRouter(router.NewPostRouter("/login", lc.Login))
	adminRG.AddRouter(router.NewPostRouter("/register", lc.Register))
	adminRG.AddRouter(router.NewGetRouter("/email_validate", uc.EmailExist))
	// user
	userRG := adminRG.NewSubGroup("users")
	{
		// userRG.AddRouter(router.NewGetRouter("/:id", uc.GetUserCache))
		userRG.AddRouter(router.NewGetRouter("/:id", uc.GetUserDB))
	}
	// adminRG.AddRouter(router.NewGetRouter("/", appMiddleware))

	return adminRG
}
