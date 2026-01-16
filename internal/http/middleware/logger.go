package middleware

import (
	"fmt"

	"github.com/dromara/carbon/v2"
	"github.com/gin-gonic/gin"
	"github.com/hhhhkkk/mini-blog/config"
)

func Logger(app *config.AppConfig) gin.HandlerFunc {
	date := carbon.Now().Format("Y-m-d")
	accessLog := fmt.Sprintf("%s-access-%s.log", app.Name, date)
	return gin.LoggerWithWriter(NewMiniBlogErrorWriter(app, accessLog))
}
