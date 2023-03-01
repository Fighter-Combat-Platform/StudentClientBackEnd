package dao

import (
	"awesomeProject/model"
	"awesomeProject/repository"
)

func CreateRoomPlayer(player model.RoomPlayer) {
	repository.CreateRoomPlayer(player)
	return
}

func DeleteRoomPlayer(player model.RoomPlayer) {
	repository.DeleteRoomPlayer(player)
	return
}

func DeleteAlloomPlayers(rid int) {
	repository.DeleteAllRoomPlayers(rid)
	return
}

func GetRoomPlayerByIds(player model.RoomPlayer) model.RoomPlayer {
	return repository.GetRoomPlayerByIds(player)
}

func GetRoomPlayersByRid(rid int) []*model.RoomPlayer {
	return repository.GetRoomPlayersByRid(rid)
}

func ModifyRoomPlayer(player model.RoomPlayer) {
	repository.ModifyRoomPlayer(player)
	return
}
