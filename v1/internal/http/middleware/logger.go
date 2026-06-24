package middleware

import (
	"fmt"

	"github.com/dromara/carbon/v2"
	"github.com/gin-gonic/gin"
	"github.com/hhhhkkk/mini-blog/v1/config"
)

func Logger(cfg config.Config) gin.HandlerFunc {
	date := carbon.Now().Format("Y-m-d")
	accessLog := fmt.Sprintf("%s-access-%s.log", cfg.Server.Name, date)
	return gin.LoggerWithWriter(NewMiniBlogErrorWriter(cfg, accessLog))
}
