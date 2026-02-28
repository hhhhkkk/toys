package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	biz "github.com/hhhhkkk/mini-blog/internal/biz/admin"
	"github.com/hhhhkkk/mini-blog/internal/biz/admin/entity"
)

type Case struct {
	userBiz *biz.UserBiz
}

func NewUserCase(userBiz *biz.UserBiz) *Case {
	return &Case{
		userBiz: userBiz,
	}
}

func (c *Case) EmailExist(ctx *gin.Context, dto *EmailExistsEmailDTO) (bool, error) {
	exist, err := c.userBiz.EmailExists(ctx, dto.Email, dto.Uid)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (c *Case) Register(ctx *gin.Context, dto *RegisterUserDTO) (*entity.User, error) {
	// 通过 gin 的方法实现 dto
	var req RegisterUserDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return nil, err
	}
	// app 层负责将 DTO 转换为基本类型传给 biz 层
	return c.userBiz.CreateUser(req.Email, req.Password, req.RepeatPassword)
}

func (c *Case) GetUserDB(ctx *gin.Context) {

}
