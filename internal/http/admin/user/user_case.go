package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hhhhkkk/mini-blog/internal/data"
	"go.uber.org/zap"
)

type UserController struct {
	logger *zap.Logger
	cache  *data.Cache
}

func NewUserController(logger *zap.Logger, cache *data.Cache) *UserController {
	return &UserController{
		logger: logger,
		cache:  cache,
	}
}

func (c *UserController) GetUserCache(ctx *gin.Context) {
	name := ctx.Param("id")

	c.logger.Info("input: " + name)

	name1, _ := c.cache.GetClient().Get("user" + name).Result()

	// time.Sleep(3 * time.Second)

	name2, _ := c.cache.GetClient().Get("user" + name).Result()

	ctx.JSON(http.StatusOK, gin.H{"msg": "get user", "id": ctx.Param("id"), "name1": name1, "name2": name2})
}
