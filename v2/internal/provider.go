//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func NewEngine() *gin.Engine {
	return gin.Default()
}

var Provider = wire.NewSet(NewEngine, NewApp)

func InitApp() (*App, error) {
	wire.Build(Provider)
	return nil, nil
}
