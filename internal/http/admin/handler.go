package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var AdminHealth gin.HandlerFunc = func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "admin ok",
	})
}

var DemoPanic gin.HandlerFunc = func(ctx *gin.Context) {
	panic("demo panic")
}
