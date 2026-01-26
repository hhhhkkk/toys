package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	logger *zap.Logger
}

func NewUserController(logger *zap.Logger) *UserController {
	return &UserController{
		logger: logger,
	}
}

func (c *UserController) GetUser(ctx *gin.Context) {
	c.logger.Info("我是 log")
	ctx.JSON(http.StatusOK, gin.H{"msg": "get user", "id": ctx.Param("id")})
}
