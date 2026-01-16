package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hhhhkkk/mini-blog/config"
)

func Tracer(app *config.AppConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader("X-TRACE-ID")
		if traceID == "" {
			id, err := uuid.NewV7()
			if err != nil {
				id = uuid.New()
			}
			traceID = id.String()
			c.Request.Header.Set("X-TRACE-ID", traceID)
		}
		c.Writer.Header().Set("X-TRACE-ID", traceID)
		c.Next()
	}
}
