package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hhhhkkk/mini-blog/v2/internal/app"
)

type Router struct {
	appService *app.AppService
}

func NewRouter(appService *app.AppService) *Router {
	return &Router{
		appService: appService,
	}
}

func (r *Router) Register(engine *gin.Engine) {
	g := engine.Group("/app")
	g.GET("health", Health)

	tg := g.Group("/task")
	{
		tg.POST("", r.appService.CreateTask)
	}
}

func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
