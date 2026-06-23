package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Health gin.HandlerFunc = func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "admin ok",
	})
}

var DemoPanic gin.HandlerFunc = func(ctx *gin.Context) {
	panic("demo panic")
}

// var appMiddleware gin.HandlerFunc = func(ctx *gin.Context) {
// 	_, exists := ctx.Get("app")
// 	if !exists {
// 		ctx.Abort()
// 	}

// 	logger := zap.NewNop()
// 	logger.Info("logggggggg")
// 	ctx.JSON(http.StatusOK, gin.H{"msg": "logged"})
// }
