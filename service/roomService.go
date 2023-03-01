package service

import (
	"awesomeProject/dao"
	"awesomeProject/model"
	"awesomeProject/utils"
	"strconv"
)

func GetRoomList() (roomList []*model.RoomListInfo) {
	var tmp []*model.Room
	tmp = dao.GetRoomList()
	for i := 0; i < len(tmp); i++ {
		info := tmp[i]

		if info.PlayerCount == 0 {
			continue
		}
		var ifPublic string
		var state string
		ifPublic = utils.IfPublic[info.IfPublic]
		state = utils.State[info.State]
		roomListInfo := &model.RoomListInfo{
			RoomName:    info.RoomName,
			OwnerName:   info.OwnerName,
			Capacity:    strconv.Itoa(info.Capacity),
			PlayerCount: strconv.Itoa(info.PlayerCount),
			IfPublic:    ifPublic,
			State:       state,
		}
		roomList = append(roomList, roomListInfo)

	}
	return
}

func CreateRoom(params utils.CreateRoomParams) (success bool, msg string, data string) {
	query := dao.GetRoomByRoomName(params.RoomName)
	if query.RoomName == "" || query.Rid == params.Uid {
		var capacity, mapID, ifPublic, weatherType, lightType int
		switch params.Type {
		case "1v1":
			capacity = 2
		case "2v2":
			capacity = 4
		case "3v3":
			capacity = 6
		}
		switch params.Map {
		case "地图1":
			mapID = 0
		case "地图2":
			mapID = 1
		}
		if params.RoomPassword != "" {
			ifPublic = 1
		} else {
			ifPublic = 0
		}
		switch params.Weather {
		case "晴朗":
			weatherType = 0
		case "多云":
			weatherType = 1
		case "大雨":
			weatherType = 2
		case "雷暴":
			weatherType = 3
		}
		switch params.Time {
		case "清晨":
			lightType = 0
		case "正午":
			lightType = 1
		case "黄昏":
			lightType = 2
		case "午夜":
			lightType = 3
		}
		dao.ModifyRoomConfig(model.Room{
			Rid:         params.Uid,
			Capacity:    capacity,
			RoomName:    params.RoomName,
			MapID:       mapID,
			WeatherType: weatherType,
			LightType:   lightType,
			IfPublic:    ifPublic,
			Password:    params.RoomPassword,
			Type:        params.Type,
		})
		success, msg, data = true, "修改房间配置成功", ""
	} else {
		success, msg, data = false, "房间名已存在", ""
	}
	return
}

func CheckRoomPassword(params utils.CheckRoomPasswordParams) (success bool, msg string, data map[string]interface{}) {
	query := dao.GetRoomByRoomName(params.RoomName)
	if query.RoomName == "" {
		success, msg, data = false, "找不到该房间", nil
	} else if query.Password != params.Password {
		success, msg, data = false, "房间密码错误", nil
	} else {
		success, msg = true, "房间密码验证成功"
	}
	return
}

func GetRoomByUid(params utils.GetRoomByUidParam) (success bool, msg string, data model.Room) {
	query := dao.GetRoomByRid(params.Id)
	if query.RoomName == "" {
		success, msg = false, "找不到该房间"
	} else {
		success, msg, data = true, "获取房间成功", query
	}
	return
}

func GetRoomConfig(params utils.GetRoomConfigParam) (success bool, msg string, data model.RoomConfig) {
	query := dao.GetRoomByRoomName(params.RoomName)
	if query.RoomName == "" {
		success, msg = false, "找不到该房间"
	} else {
		success, msg, data = true, "获取房间配置信息成功", model.RoomConfig{
			Rid:         query.Rid,
			Weather:     utils.Weather[query.WeatherType],
			Map:         utils.Map[query.MapID],
			Time:        utils.Time[query.LightType],
			OwnerName:   query.OwnerName,
			PlayerCount: strconv.Itoa(query.PlayerCount) + "/" + strconv.Itoa(query.Capacity),
			State:       query.State,
		}
	}
	return
}

func BeginPre(param utils.BeginParam) (success bool, msg string, data []*model.RoomPlayer) {
	roomPlayers := dao.GetRoomPlayersByRid(param.Rid)
	red := false
	blue := false
	if len(roomPlayers) == 1 {
		success, msg = false, "房间内只有一人，无法开始"
		return
	}
	for i := 0; i < len(roomPlayers); i++ {
		if roomPlayers[i].GroupType == 0 {
			red = true
		}
		if roomPlayers[i].GroupType == 1 {
			blue = true
		}
		if roomPlayers[i].Ready == 0 {
			if roomPlayers[i].Uid == param.Rid {
				continue
			} else {
				success, msg = false, "有玩家未准备，无法开始"
				return
			}
		}
	}
	if blue != true || red != true {
		success, msg = false, "所有玩家均在同一阵营，无法开始"
		return
	}
	for i := 0; i < len(roomPlayers); i++ {
		dao.ModifyRoomPlayer(model.RoomPlayer{
			Uid:         roomPlayers[i].Uid,
			Rid:         roomPlayers[i].Rid,
			Ready:       0,
			GroupType:   roomPlayers[i].GroupType,
			FighterType: roomPlayers[i].FighterType,
		})
	}
	return true, "开始游戏准备就绪", roomPlayers
}

func Begin(param utils.BeginParam) (success bool, msg string) {
	room := dao.GetRoomByRid(param.Rid)
	room.State = 0
	room.Mid = room.Mid + 1
	dao.ModifyRoom(room)

	//utils.Clear(strconv.Itoa(param.Rid))

	return true, "游戏开始"
}

func End(param utils.EndParam) (success bool, msg string) {
	room := dao.GetRoomByRid(param.Rid)
	room.State = 1
	dao.ModifyRoom(room)
	return true, "游戏结束"
}
