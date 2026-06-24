//go:build wireinject
// +build wireinject

// go:build wireinject
package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/v1/config"
	"github.com/hhhhkkk/mini-blog/v1/internal/app"
	"github.com/hhhhkkk/mini-blog/v1/internal/biz"
	"github.com/hhhhkkk/mini-blog/v1/internal/data"
	"github.com/hhhhkkk/mini-blog/v1/internal/http"
	"github.com/hhhhkkk/mini-blog/v1/internal/http/middleware"
	"github.com/hhhhkkk/mini-blog/v1/internal/job"
)

func NewEngine() *gin.Engine {
	return gin.New()
}

var ProviderSet = wire.NewSet(
	NewLogger,

	http.RouterProviderSet,
	app.ProviderSet,
	middleware.ProviderSet,
	biz.ProviderSet,
	data.ProviderSet,
	config.ProviderSet,
	job.ProviderSet,
	NewEngine,
	NewApp,
)

func InitApp() (*App, error) {
	wire.Build(ProviderSet)
	return nil, nil
}
