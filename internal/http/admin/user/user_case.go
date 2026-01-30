package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hhhhkkk/mini-blog/internal/data"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserController struct {
	logger *zap.Logger
	cache  *data.Cache
	db     *data.DB
}

func NewUserController(db *data.DB, cache *data.Cache, logger *zap.Logger) *UserController {
	return &UserController{
		logger: logger,
		cache:  cache,
		db:     db,
	}
}

func (c *UserController) GetUserDB(ctx *gin.Context) {
	id := ctx.Param("id")
	db := c.db.GetClient()
	u, err := gorm.G[data.User](db).Where("id", id).First(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "get user failed", "err": err.Error(), "dsn": ""})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"msg": "get user", "id": ctx.Param("id"), "name": u.Name, "age": u.Age})
}

func (c *UserController) GetUserCache(ctx *gin.Context) {
	name := ctx.Param("id")

	c.logger.Info("input: " + name)

	name1, _ := c.cache.GetClient().Get("user" + name).Result()

	// time.Sleep(3 * time.Second)

	name2, _ := c.cache.GetClient().Get("user" + name).Result()

	ctx.JSON(http.StatusOK, gin.H{"msg": "get user", "id": ctx.Param("id"), "name1": name1, "name2": name2})
}
