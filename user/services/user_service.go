package services

import (
	"context"
	"gin-micro/helpers/status"
	"gin-micro/helpers/token"
	"gin-micro/protos/user"
)

type UserService struct { }

func (userService *UserService) SignIn(ctx context.Context, userinfo *user.UserInfo, result *user.Result) error {
	result.Status = status.SaveStatusOK
	return nil
}

// 用户方法
func (userService *UserService) Login(ctx context.Context,loginParams *user.LoginParams, result *user.Result) error {
	username := loginParams.GetAccount();password := loginParams.GetPassword()
	if username == "admin" && password == "123456" {
		result.Status = status.LoginStatusOK
		infos := map[string]string{"account":"admin","nickname":"皮卡丘"}
		astoken := token.GenerateToken(infos)
		infos["ACCESS_TOKEN"]= astoken
		result.Map = infos
	}else {
		result.Status = status.LoginStatusErr
	}
	return nil
}


