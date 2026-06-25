package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppService struct{}

func NewAppService() *AppService {
	return &AppService{}
}

func (t *AppService) CreateTask(ctx *gin.Context) {
	dto := &CreateTask{}
	if err := ctx.ShouldBind(dto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "param invalid",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "ok",
		"data": dto,
	})
}
