package admin

import (
	"github.com/gin-gonic/gin"
	biz "github.com/hhhhkkk/mini-blog/internal/biz/admin"
)

type UserCase struct {
	userBiz *biz.UserBiz
}

func NewUserCase(userBiz *biz.UserBiz) *UserCase {
	return &UserCase{
		userBiz: userBiz,
	}
}

func (c *UserCase) EmailExist(ctx *gin.Context, dto *EmailExistsDTO) bool {
	return c.userBiz.EmailExists(dto.Email, dto.Uid)
}

func (c *UserCase) Register(ctx *gin.Context, dto *RegisterUserDTO) (map[string]interface{}, error) {
	user, err := c.userBiz.CreateUser(dto.Email, dto.Password)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]interface{})
	ret["id"] = user.Id
	ret["email"] = user.Email
	return ret, nil
}

func (c *UserCase) GetUserDB(ctx *gin.Context) {

}
