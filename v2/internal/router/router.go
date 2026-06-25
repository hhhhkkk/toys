package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hhhhkkk/mini-blog/v2/internal/app"
	"github.com/hhhhkkk/mini-blog/v2/internal/cache"
)

type Router struct {
	appService   *app.AppService
	cacheService *cache.CacheService
}

func NewRouter(appService *app.AppService, cacheService *cache.CacheService) *Router {
	return &Router{
		appService:   appService,
		cacheService: cacheService,
	}
}

func (r *Router) Register(engine *gin.Engine) {
	g := engine.Group("/app")
	g.GET("health", Health)

	tg := g.Group("/task")
	{
		tg.POST("", r.appService.CreateTask)
	}

	cg := g.Group("cache")
	{
		cg.POST("", r.cacheService.Add)
		cg.DELETE("/:key", r.cacheService.Del)
		cg.GET("/:key", r.cacheService.Get)
	}
}

func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
