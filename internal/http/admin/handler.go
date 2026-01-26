package admin

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hhhhkkk/mini-blog/config"
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

var configPath gin.HandlerFunc = func(ctx *gin.Context) {
	ret := make(map[int]int)
	c := config.NewCacheConfig()
	go func() {
		select {
		case <-ctx.Done():
			return
		case nc := <-c.WatchConfig():
			c = nc
		}
	}()
	for i := range 10 {
		ret[i] = c.DB
		fmt.Println(c.Host)
		time.Sleep(500 * time.Millisecond)
	}
	ctx.JSON(http.StatusOK, ret)
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
