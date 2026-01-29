package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/hhhhkkk/mini-blog/config"
)

func NewBaseMiddleware(app config.Config) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		Recovery(app),
		Logger(app),
		Tracer(app),
	}
}

var ProviderSet = wire.NewSet(NewBaseMiddleware)
