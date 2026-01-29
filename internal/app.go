package internal

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hhhhkkk/mini-blog/config"
	"github.com/hhhhkkk/mini-blog/internal/job"
	"github.com/hhhhkkk/mini-blog/router"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

type App struct {
	engine         *gin.Engine
	config         config.Config
	jobGroup       []job.IJobGroup
	baseMiddleware []gin.HandlerFunc
	routers        []router.IRouterGroup
	logger         *zap.Logger
}

func NewApp(
	engine *gin.Engine,
	appConfig config.Config,
	jobGroup []job.IJobGroup,
	baseMiddleware []gin.HandlerFunc,
	routers []router.IRouterGroup,
	logger *zap.Logger,
) *App {
	app := &App{
		engine:         engine,
		config:         appConfig,
		jobGroup:       jobGroup,
		baseMiddleware: baseMiddleware,
		routers:        routers,
		logger:         logger,
	}
	app.setupEngine()
	app.setupRouter()
	return app
}

func (app *App) setupEngine() {
	gin.DisableConsoleColor()

	// 根据环境设置 Gin 模式：生产环境使用 ReleaseMode，其他环境使用 DebugMode（默认）
	if app.config.Server.Env == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	app.engine.Use(func(context *gin.Context) {
		context.Set("app", app)
		context.Next()
	})
	app.engine.Use(app.baseMiddleware...)
}

// func (app *App) setupJobGroup() {
// app.jobGroup = append(app.jobGroup)
// }

func registerRoute(rg router.IRouterGroup, app *App) {
	ng := app.engine.Group("/" + rg.Name())
	ng.Use(rg.Middlewares()...)
	for _, r := range rg.Routers() {
		ng.Handle(r.Method(), r.Path(), r.Handler())
	}
	for _, subGroup := range rg.SubGroups() {
		registerRoute(subGroup, app)
	}
}

func (app *App) setupRouter() {
	for _, rg := range app.routers {
		registerRoute(rg, app)
	}

	for index, router := range app.engine.Routes() {
		app.logger.Info(fmt.Sprintf("index: %d, method: %s, path: %s, handler: %s\n", index, router.Method, router.Path, router.Handler))
	}
}

func NewLogger() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return logger
}

func (app *App) Run() error {
	g.Go(func() error {
		addr := fmt.Sprintf("%s:%d", app.config.Server.Host, app.config.Server.Port)
		if err := app.engine.Run(addr); err != nil {
			return err
		}
		return nil
	})
	// todo job
	return g.Wait()
}
