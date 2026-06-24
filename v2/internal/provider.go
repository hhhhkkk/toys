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

var Provider = wire.NewSet(NewEngine)

func InitApp() (*App, error) {
	wire.Build()
	return nil, nil
}
