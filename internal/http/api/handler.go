package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Health gin.HandlerFunc = func(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "api ok",
	})
}
