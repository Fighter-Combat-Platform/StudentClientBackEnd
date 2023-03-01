package router

import (
	"awesomeProject/controller"
	"github.com/kataras/iris/v12"
)

func SetRouter(app *iris.Application) {
	root := app.Party("/")

	root.Handle("POST", "/login", controller.Login)
	root.Handle("POST", "/register", controller.Register)
	root.Handle("POST", "/modifyUserPassword", controller.ModifyUserPassword)
	root.Handle("POST", "/modifyUsername", controller.ModifyUsername)

	root.Handle("POST", "/createRoom", controller.CreateRoom)
	root.Handle("GET", "/getRoomList", controller.GetRoomList)
	root.Handle("POST", "/checkRoomPassword", controller.CheckRoomPassword)
	root.Handle("POST", "/getRoomByUid", controller.GetRoomByUid)
	root.Handle("POST", "/getRoomConfig", controller.GetRoomConfig)
	root.Handle("POST", "/begin", controller.Begin)
	root.Handle("POST", "/end", controller.End)

	root.Handle("POST", "/enterRoom", controller.EnterRoom)
	root.Handle("POST", "/quitRoom", controller.QuitRoom)
	root.Handle("POST", "/getRoomPlayers", controller.GetRoomPlayers)
	root.Handle("POST", "/modifyRoomPlayer", controller.ModifyRoomPlayer)
	root.Handle("POST", "/roomPlayerReady", controller.RoomPlayerReady)
	root.Handle("POST", "/ownerQuitRoom", controller.OwnerQuitRoom)
	root.Handle("POST", "/kickPlayer", controller.KickPlayer)

	root.Handle("GET", "/roomInfoUpdate", controller.RoomInfoUpdate)

}
