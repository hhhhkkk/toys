package internal

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hhhhkkk/mini-blog/config"
	"github.com/hhhhkkk/mini-blog/internal/job"
	"github.com/hhhhkkk/mini-blog/router"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

type App struct {
	engine         *gin.Engine
	config         *config.AppConfig
	jobGroup       []job.IJobGroup
	baseMiddleware []gin.HandlerFunc
	routers        []router.IRouterGroup
}

func NewApp(
	engine *gin.Engine,
	appConfig *config.AppConfig,
	jobGroup []job.IJobGroup,
	baseMiddleware []gin.HandlerFunc,
	routers []router.IRouterGroup,
) *App {
	app := &App{
		engine:         engine,
		config:         appConfig,
		jobGroup:       jobGroup,
		baseMiddleware: baseMiddleware,
		routers:        routers,
	}
	app.setupEngine()
	app.setupRouter()
	return app
}

func (app *App) setupEngine() {
	gin.DisableConsoleColor()
	
	// 根据环境设置 Gin 模式：生产环境使用 ReleaseMode，其他环境使用 DebugMode（默认）
	if app.config.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	app.engine.Use(app.baseMiddleware...)
}

// func (app *App) setupJobGroup() {
// app.jobGroup = append(app.jobGroup)
// }

func (app *App) setupRouter() {
	for _, rg := range app.routers {
		ng := app.engine.Group("/" + rg.Name())
		ng.Use(rg.Middlewares()...)
		for _, r := range rg.Routers() {
			ng.Handle(r.Method(), r.Path(), r.Handler())
		}
	}
}

func (app *App) Run() error {
	g.Go(func() error {
		addr := fmt.Sprintf("%s:%d", app.config.Host, app.config.Port)
		if err := app.engine.Run(addr); err != nil {
			return err
		}
		return nil
	})
	// todo job
	return g.Wait()
}
