// go:build wireinject
//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/v2/config"
	"github.com/hhhhkkk/mini-blog/v2/internal/app"
	"github.com/hhhhkkk/mini-blog/v2/internal/cache"
	"github.com/hhhhkkk/mini-blog/v2/internal/router"
	"github.com/hhhhkkk/mini-blog/v2/internal/service"
)

func NewEngine() *gin.Engine {
	return gin.Default()
}

var Provider = wire.NewSet(NewEngine, NewApp, app.ProviderSet, router.ProviderSet, cache.ProviderSet, service.ProviderSet, config.ProviderSet)

func InitApp() (*App, error) {
	wire.Build(Provider)
	return nil, nil
}
