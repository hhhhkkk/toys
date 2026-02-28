package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	app "github.com/hhhhkkk/mini-blog/internal/app/admin/user"
)

type UserController struct {
	uc *app.Case
}

func NewUserController(uc *app.Case) *UserController {
	return &UserController{
		uc: uc,
	}
}

func (c *UserController) EmailExist(ctx *gin.Context) {
	// 手写自己实现 dto
	email := ctx.Param("email")
	_uid, _ := ctx.Get("uid")
	uid, ok := _uid.(uint)
	if !ok {
		uid = 0
	}

	dto := &app.EmailExistsEmailDTO{
		Email: email,
		Uid:   uid,
	}
	exist, err := c.uc.EmailExist(ctx, dto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"exist": exist})
}

func (c *UserController) Register(ctx *gin.Context) {
	// 通过 gin 的方法实现 dto
	var req app.RegisterUserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
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

func (c *UserController) GetUserDB(ctx *gin.Context) {

}
