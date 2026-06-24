package middleware

import (
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.GetHeader("Authorization")
		// if token == "" {
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// 	return
		// }
		c.Set("uid", 100)
		c.Next()
	}
}
