package api

import (
	"github.com/hhhhkkk/mini-blog/v1/router"
)

func NewApiRouterProvider(user *UserController) router.IRouterGroup {
	apiRG := router.NewRouterGroup("api")

	users := apiRG.NewSubGroup("/users")
	{
		users.AddRouter(router.NewPostRouter("/register", user.Register))
	}

	return apiRG
}
