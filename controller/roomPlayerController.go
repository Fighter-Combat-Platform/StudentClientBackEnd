package controller

import (
	"awesomeProject/service"
	"awesomeProject/utils"
	"github.com/kataras/iris/v12"
)

func EnterRoom(ctx iris.Context) {
	var params utils.EnterRoomParams
	if !utils.GetContextParams(ctx, &params) {
		return
	}

	success, msg, data := service.EnterRoom(params)

	utils.SendResponse(ctx, success, msg, data)

	return
}

func QuitRoom(ctx iris.Context) {
	var params utils.QuitRoomParams
	if !utils.GetContextParams(ctx, &params) {
		return
	}

	success, msg, data := service.QuitRoom(params)

	utils.SendResponse(ctx, success, msg, data)

	return
}

func OwnerQuitRoom(ctx iris.Context) {
	var params utils.OwnerQuitRoomParams
	if !utils.GetContextParams(ctx, &params) {
		return
	}

	success, msg, data := service.OwnerQuitRoom(params)

	utils.SendResponse(ctx, success, msg, data)

	return
}

func KickPlayer(ctx iris.Context) {
	var params utils.KickPlayerParams
	if !utils.GetContextParams(ctx, &params) {
		return
	}

	success, msg, data := service.KickPlayer(params)

	utils.SendResponse(ctx, success, msg, data)

	return
}

func GetRoomPlayers(ctx iris.Context) {
	var params utils.GetRoomPlayersParam
	if !utils.GetContextParams(ctx, &params) {
		return
	}

	success, msg, data := service.GetRoomPlayers(params)

	utils.SendResponse(ctx, success, msg, data)

	return
}

func ModifyRoomPlayer(ctx iris.Context) {
	var params utils.ModifyRoomPlayerParams
	if !utils.GetContextParams(ctx, &params) {
		return
	}

	success, msg, data := service.ModifyRoomPlayer(params)

	utils.SendResponse(ctx, success, msg, data)

	return
}

func RoomPlayerReady(ctx iris.Context) {
	println("room player ready")
	var params utils.RoomPlayerReadyParams
	if !utils.GetContextParams(ctx, &params) {
		return
	}

	success, msg, data := service.RoomPlayerReady(params)

	utils.SendResponse(ctx, success, msg, data)

	return
}
