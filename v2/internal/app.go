package internal

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
)

var START_ERROR = errors.New("gin 启动失败")

type App struct {
	engine *gin.Engine
}

func NewApp(engine *gin.Engine) *App {
	return &App{
		engine: engine,
	}
}

func (app *App) Run() error {
	fmt.Println("v2 starting...")
	if err := app.engine.Run(":8081"); err != nil {
		return START_ERROR
	}
	return nil
}
