package internal

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hhhhkkk/mini-blog/v2/internal/router"
)

var START_ERROR = errors.New("gin 启动失败")

type App struct {
	engine *gin.Engine
	router *router.Router
}

func NewApp(engine *gin.Engine, router *router.Router) *App {
	return &App{
		engine: engine,
		router: router,
	}
}

func (app *App) Run() error {
	fmt.Println("v2 starting...")

	app.router.Register(app.engine)

	if err := app.engine.Run(":8081"); err != nil {
		return START_ERROR
	}
	return nil
}
