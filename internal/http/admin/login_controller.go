package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	app "github.com/hhhhkkk/mini-blog/internal/app/admin"
)

type LoginController struct {
	uc *app.UserCase
}

func NewLoginController(uc *app.UserCase) *LoginController {
	return &LoginController{
		uc: uc,
	}
}

func (c LoginController) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "login success"})
}

func (c LoginController) Register(ctx *gin.Context) {
	// 通过 gin 的方法实现 dto
	var req app.RegisterUserDTO
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := c.uc.Register(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}
