package dao

import (
	"awesomeProject/model"
	"awesomeProject/repository"
)

func GetRoomByRoomName(roomName string) model.Room {
	return repository.GetRoomByRoomName(roomName)
}

func GetRidByRoomName(roomName string) int {
	return repository.GetRidByRoomName(roomName)
}

func AddRoomPlayerCount(roomName string) {
	repository.AddRoomPlayerCount(roomName)
	return
}

func SubRoomPlayerCount(roomName string) {
	repository.SubRoomPlayerCount(roomName)
	return
}

func GetRoomByRid(rid int) model.Room {
	return repository.GetRoomByRid(rid)
}

func ModifyRoomConfig(room model.Room) {
	repository.ModifyRoomConfig(room)
	return
}

func ModifyRoom(room model.Room) {
	repository.ModifyRoom(room)
	return
}

func CreateRoom(room model.Room) {
	repository.CreateRoom(room)
	return
}

func GetRoomList() []*model.Room {
	return repository.GetRoomList()
}
