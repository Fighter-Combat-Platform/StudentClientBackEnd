package controller

import (
	"awesomeProject/service"
	"awesomeProject/utils"
	"github.com/kataras/iris/v12"
)

func Login(ctx iris.Context) {
	var params utils.LoginParams
	if !utils.GetContextParams(ctx, &params) {
		return
	}

	success, msg, data := service.Login(params)

	utils.SendResponse(ctx, success, msg, data)

	return
}

func Register(ctx iris.Context) {
	var params utils.RegisterParams
	if !utils.GetContextParams(ctx, &params) {
		return
	}

	success, msg, data := service.Register(params)

	utils.SendResponse(ctx, success, msg, data)

	return
}

func ModifyUserPassword(ctx iris.Context) {
	var params utils.ModifyUserPasswordParams
	if !utils.GetContextParams(ctx, &params) {
		return
	}

	success, msg, data := service.ModifyUserPassword(params)

	utils.SendResponse(ctx, success, msg, data)

	return
}

func ModifyUsername(ctx iris.Context) {
	var params utils.ModifyUsernameParams
	if !utils.GetContextParams(ctx, &params) {
		return
	}

	success, msg, data := service.ModifyUsername(params)

	utils.SendResponse(ctx, success, msg, data)

	return
}
