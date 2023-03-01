package service

import (
	"awesomeProject/dao"
	"awesomeProject/model"
	"awesomeProject/utils"
)

func EnterRoom(params utils.EnterRoomParams) (success bool, msg string, data interface{}) {
	room := dao.GetRoomByRoomName(params.RoomName)
	if room.Capacity == room.PlayerCount {
		success, msg = false, "房间人数已满"
	} else if room.State == 0 {
		success, msg = false, "房间处于战斗中"
	} else {
		roomPlayer := dao.GetRoomPlayerByIds(model.RoomPlayer{
			Rid: room.Rid,
			Uid: params.Uid,
		})
		if roomPlayer.Rid == 0 {
			dao.AddRoomPlayerCount(params.RoomName)
			dao.CreateRoomPlayer(model.RoomPlayer{
				Rid:         room.Rid,
				Uid:         params.Uid,
				FighterType: 0,
				GroupType:   0,
				Ready:       0,
			})
			success, msg, data = true, "进入房间成功", room.Rid
		} else {
			dao.ModifyRoomPlayer(model.RoomPlayer{
				Rid:         roomPlayer.Rid,
				Uid:         roomPlayer.Uid,
				FighterType: roomPlayer.FighterType,
				GroupType:   roomPlayer.GroupType,
				Ready:       0,
			})
			success, msg, data = true, "用户已进入该房间", room.Rid
		}
	}
	return
}

func QuitRoom(params utils.QuitRoomParams) (success bool, msg string, data interface{}) {
	room := dao.GetRoomByRoomName(params.RoomName)
	if room.PlayerCount == 0 {
		success, msg = false, "房间内已无人"
	} else {
		roomPlayer := dao.GetRoomPlayerByIds(model.RoomPlayer{
			Rid: room.Rid,
			Uid: params.Uid,
		})
		if roomPlayer.Uid != 0 {
			dao.SubRoomPlayerCount(params.RoomName)
			dao.DeleteRoomPlayer(model.RoomPlayer{
				Rid: room.Rid,
				Uid: params.Uid,
			})
			success, msg = true, "退出房间成功"
		} else {
			success, msg = true, "用户已退出房间"
		}
	}
	return
}

func OwnerQuitRoom(params utils.OwnerQuitRoomParams) (success bool, msg string, data interface{}) {
	dao.DeleteAlloomPlayers(params.Rid)
	room := dao.GetRoomByRid(params.Rid)
	room.PlayerCount = 0
	dao.ModifyRoom(room)
	success, msg = true, "房主已退出"
	return
}

func KickPlayer(params utils.KickPlayerParams) (success bool, msg string, data interface{}) {
	user := dao.GetUserByUsername(params.Username)
	if user.Username == "" {
		success, msg = false, "该玩家不存在"
		return
	}
	roomPlayer := dao.GetRoomPlayerByIds(model.RoomPlayer{
		Uid: user.Uid,
		Rid: params.Rid,
	})
	if roomPlayer.Rid == 0 {
		success, msg = false, "该玩家不在房间内"
		return
	}
	dao.DeleteRoomPlayer(roomPlayer)
	room := dao.GetRoomByRid(params.Rid)
	room.PlayerCount = room.PlayerCount - 1
	dao.ModifyRoom(room)
	success, msg = true, "玩家已被踢出"
	return
}

func GetRoomPlayers(params utils.GetRoomPlayersParam) (success bool, msg string, data []*model.RoomPlayerInfo) {
	roomPlayers := dao.GetRoomPlayersByRid(params.Rid)
	var roomPlayersInfo []*model.RoomPlayerInfo
	for i := 0; i < len(roomPlayers); i++ {
		roomPlayer := roomPlayers[i]
		tmp := &model.RoomPlayerInfo{
			Uid:         roomPlayer.Uid,
			PlayerName:  dao.GetUserByUid(roomPlayer.Uid).Username,
			GroupType:   utils.GroupType[roomPlayer.GroupType],
			FighterType: utils.FighterType[roomPlayer.FighterType],
			Ready:       utils.Ready[roomPlayer.Ready],
		}
		roomPlayersInfo = append(roomPlayersInfo, tmp)
	}
	success, msg, data = true, "获取房间成员信息成功", roomPlayersInfo
	return
}

func ModifyRoomPlayer(params utils.ModifyRoomPlayerParams) (success bool, msg string, data interface{}) {
	var groupType, fighterType, ready int
	switch params.Ready {
	case "已准备":
		ready = 1
	case "未准备":
		ready = 0
	}
	switch params.GroupType {
	case "红方":
		groupType = 0
	case "蓝方":
		groupType = 1
	}
	switch params.FighterType {
	case "A型战机":
		fighterType = 0
	case "B型战机":
		fighterType = 1
	}
	dao.ModifyRoomPlayer(model.RoomPlayer{
		Uid:         params.Uid,
		Rid:         params.Rid,
		Ready:       ready,
		GroupType:   groupType,
		FighterType: fighterType,
	})
	success, msg = true, "修改房间成员信息成功"
	return
}

func RoomPlayerReady(params utils.RoomPlayerReadyParams) (success bool, msg string, data interface{}) {
	query := dao.GetRoomPlayerByIds(model.RoomPlayer{
		Uid: params.Uid,
		Rid: params.Rid,
	})
	var ready int
	if query.Ready == 0 {
		ready = 1
	} else {
		ready = 0
	}
	dao.ModifyRoomPlayer(model.RoomPlayer{
		Uid:         query.Uid,
		FighterType: query.FighterType,
		Rid:         query.Rid,
		GroupType:   query.GroupType,
		Ready:       ready,
	})
	success, msg, data = true, "房间成员准备状态切换成功", utils.Ready[ready]
	return
}
