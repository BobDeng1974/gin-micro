package controls

import (
	"gin-micro/api/clients"
	"gin-micro/helpers/response"
	"gin-micro/helpers/status"
	"gin-micro/protos/user"
	"github.com/gin-gonic/gin"
)

var userService = user.NewUserService("micro.user.svr",clients.Client)

type UserControl struct {
	response.Control
}

// 用户登陆接口
// @Summary 用户登陆接口
// @Tags UserControl
// @Accept json
// @Produce json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Param code     query string false "验证码"
// @Success 200 {object} response.JsonObject
// @Router /user/api/login [post]
func (c *UserControl) Login( ctx *gin.Context)  {
	params := &user.LoginParams{}
	if err := ctx.Bind(params);err == nil {
		result, err := userService.Login(ctx, params)
		if err == nil {
			c.SuccessContent(ctx,status.StatusText(result.Status),result.Map)
		}else {
			c.InternalError(ctx,err.Error())
		}
	}else {
		c.BindingError(ctx)
	}
	ctx.Abort();
}

// 用户注册接口
func ( c *UserControl) SignIn(ctx *gin.Context)  {
	result, err := userService.SignIn(ctx, &user.UserInfo{})
	if err == nil {
		c.SuccessContent(ctx,status.StatusText(result.Status),result.Map)
	}else {
		c.InternalError(ctx,err.Error())
	}
	ctx.Abort()
}
