package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	app "github.com/hhhhkkk/mini-blog/v1/internal/app/admin"
)

type Controller struct {
	uc *app.UserCase
}

func NewUserController(uc *app.UserCase) *Controller {
	return &Controller{
		uc: uc,
	}
}

func (c *Controller) EmailExist(ctx *gin.Context) {
	// 手写自己实现 dto
	var req app.EmailExistsDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	exist := c.uc.EmailExist(ctx, &req)
	ctx.JSON(http.StatusOK, gin.H{"valid": exist})
}

func (c *Controller) GetUserDB(ctx *gin.Context) {

}
