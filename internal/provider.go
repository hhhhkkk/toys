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
	"go.uber.org/zap"
)

func NewEngine() *gin.Engine {
	return gin.New()
}

func NewLogger() *zap.Logger {
	logger := zap.NewNop()
	return logger.Named("test")
}

var ProviderSet = wire.NewSet(
	http.RouterProviderSet,
	middleware.ProviderSet,
	biz.ProviderSet,
	data.ProviderSet,
	config.ProviderSet,
	job.ProviderSet,
	NewEngine,
	NewLogger,
	NewApp,
)

func InitApp() (*App, error) {
	wire.Build(ProviderSet)
	return nil, nil
}
