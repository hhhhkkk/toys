//go:build wireinject
// +build wireinject

// go:build wireinject
package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/config"
	"github.com/hhhhkkk/mini-blog/internal/biz"
	"github.com/hhhhkkk/mini-blog/internal/data"
	"github.com/hhhhkkk/mini-blog/internal/http"
	"github.com/hhhhkkk/mini-blog/internal/http/middleware"
	"github.com/hhhhkkk/mini-blog/internal/job"
)

func NewEngine() *gin.Engine {
	return gin.New()
}

var ProviderSet = wire.NewSet(
	NewLogger,

	http.RouterProviderSet,
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
