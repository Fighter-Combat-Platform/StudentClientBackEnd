package controller

import (
	"awesomeProject/model"
	"awesomeProject/service"
	"awesomeProject/utils"
	"fmt"
	"github.com/kataras/iris/v12"
	"log"
	"time"
)

func GetRoomList(ctx iris.Context) {
	utils.SendResponse(ctx, true, "房间列表获取成功", service.GetRoomList())
}

func CreateRoom(ctx iris.Context) {
	var params utils.CreateRoomParams
	if !utils.GetContextParams(ctx, &params) {
		return
	}

	success, msg, data := service.CreateRoom(params)

	utils.SendResponse(ctx, success, msg, data)

	return
}

func CheckRoomPassword(ctx iris.Context) {
	var params utils.CheckRoomPasswordParams
	if !utils.GetContextParams(ctx, &params) {
		return
	}

	success, msg, data := service.CheckRoomPassword(params)

	utils.SendResponse(ctx, success, msg, data)

	return
}

func GetRoomByUid(ctx iris.Context) {
	var params utils.GetRoomByUidParam
	if !utils.GetContextParams(ctx, &params) {
		return
	}

	success, msg, data := service.GetRoomByUid(params)

	utils.SendResponse(ctx, success, msg, data)

	return
}

func GetRoomConfig(ctx iris.Context) {
	var params utils.GetRoomConfigParam
	if !utils.GetContextParams(ctx, &params) {
		return
	}

	success, msg, data := service.GetRoomConfig(params)

	utils.SendResponse(ctx, success, msg, data)

	return
}

func Begin(ctx iris.Context) {
	var param utils.BeginParam
	if !utils.GetContextParams(ctx, &param) {
		return
	}

	success, msg, roomPlayers := service.BeginPre(param)
	if success != false {
		_, _, room := service.GetRoomByUid(utils.GetRoomByUidParam{
			Id: param.Rid,
		})

		conn, err := InitTcpSocket()
		if err != nil {
			log.Println(err)
			return
		}

		defer conn.Close()

		success, msg = Sender(conn, model.Begin{
			OperationType: 3,
			DaotiaoType:   1,
			Timestamp:     int32(time.Now().Unix()),
			ContentLength: int32(28 + 12*len(roomPlayers)),
			CallbackID:    0,
			Room:          room,
			UsersInfo:     roomPlayers,
		})

	}
	println(msg)
	if success == true {
		success, msg = service.Begin(param)
	}

	utils.SendResponse(ctx, success, msg, "")

	return
}

func End(ctx iris.Context) {
	ctx.ResponseWriter()
	var param utils.EndParam
	if !utils.GetContextParams(ctx, &param) {
		return
	}

	success, msg := service.End(param)

	utils.SendResponse(ctx, success, msg, "")

	return
}

func RoomInfoUpdate(ctx iris.Context) {
	w := ctx.ResponseWriter()
	r := ctx.Request()
	//通过升级后的升级器得到链接
	conn, err := utils.Up.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("获取连接失败:", err)
		return
	}
	var arg string
	values := r.URL.Query()
	arg = values.Get("id")
	//println(arg)
	//连接成功后注册用户
	user := &utils.User{
		Id:   arg,
		Conn: conn,
		Msg:  make(chan []byte),
	}
	utils.Hub.Register <- user
	defer func() {
		utils.Hub.Unregister <- user
	}()
	//得到连接后，就可以开始读写数据了
	go utils.Read(user)
	utils.Write(user)
}
