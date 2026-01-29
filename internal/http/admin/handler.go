package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var AdminHealth gin.HandlerFunc = func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "admin ok",
	})
}

var DemoPanic gin.HandlerFunc = func(ctx *gin.Context) {
	panic("demo panic")
}

var appMiddleware gin.HandlerFunc = func(ctx *gin.Context) {
	_, exists := ctx.Get("app")
	if !exists {
		ctx.Abort()
	}

	logger := zap.NewNop()
	logger.Info("logggggggg")
	ctx.JSON(http.StatusOK, gin.H{"msg": "logged"})
}
