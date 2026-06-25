// go:build wireinject
//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/v2/internal/app"
	"github.com/hhhhkkk/mini-blog/v2/internal/router"
)

func NewEngine() *gin.Engine {
	return gin.Default()
}

var Provider = wire.NewSet(NewEngine, NewApp, app.ProviderSet, router.ProviderSet)

func InitApp() (*App, error) {
	wire.Build(Provider)
	return nil, nil
}
