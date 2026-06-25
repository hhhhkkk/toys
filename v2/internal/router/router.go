package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Router struct{}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Register(engine *gin.Engine) {
	g := engine.Group("/app")
	{
		g.GET("health", Health)
	}
}

func Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
}
