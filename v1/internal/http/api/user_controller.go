package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hhhhkkk/mini-blog/v1/internal/app/api"
)

type UserController struct {
	uc *api.UserCase
}

func NewUser(uc *api.UserCase) *UserController {
	return &UserController{
		uc: uc,
	}
}

func (u *UserController) Register(ctx *gin.Context) {
	var req api.RegisterReq
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := u.uc.Register(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
