package utils

import (
	"github.com/gorilla/websocket"
	"github.com/kataras/neffos"
)

const IsTest = true

const (
	InvalidFormat = 0
)

var State = [5]string{"战斗中", "闲置中", "房间初始化中", "结算界面", "导调介入"}
var IfPublic = [2]string{"是", "否"}
var Weather = [4]string{"晴朗", "多云", "大雨", "雷暴"}
var Map = [2]string{"地图1", "地图2"}
var Time = [4]string{"清晨", "正午", "黄昏", "午夜"}
var GroupType = [2]string{"红方", "蓝方"}
var FighterType = [2]string{"A型战机", "B型战机"}
var Ready = [2]string{"未准备", "已准备"}

var Ws *neffos.Server

var Upgrader = websocket.Upgrader{}

// ResponseBean - Structure of Response
type ResponseBean struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

/* Structure of Request Parameters */
type LoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	SID      string `json:"sid"`
}

type ModifyUserPasswordParams struct {
	Uid         int    `json:"uid"`
	NewPassword string `json:"newPassword"`
}

type ModifyUsernameParams struct {
	Uid         int    `json:"uid"`
	NewUsername string `json:"newUsername"`
}

type CreateRoomParams struct {
	Uid          int    `json:"uid"`
	RoomName     string `json:"roomName"`
	Type         string `json:"type"`
	Map          string `json:"map"`
	Weather      string `json:"weather"`
	Time         string `json:"time"`
	RoomPassword string `json:"roomPassword"`
}

type CheckRoomPasswordParams struct {
	RoomName string `json:"roomName"`
	Password string `json:"password"`
}

type EnterRoomParams struct {
	Uid      int    `json:"uid"`
	RoomName string `json:"roomName"`
}

type GetRoomByUidParam struct {
	Id int `json:"id"`
}

type QuitRoomParams struct {
	Uid      int    `json:"uid"`
	RoomName string `json:"roomName"`
}

type OwnerQuitRoomParams struct {
	Rid int `json:"rid"`
}

type GetRoomConfigParam struct {
	RoomName string `json:"roomName"`
}

type EndParam struct {
	Rid int `json:"rid"`
}

type GetRoomPlayersParam struct {
	Rid int `json:"rid"`
}
type ModifyRoomPlayerParams struct {
	Rid         int    `json:"rid"`
	Uid         int    `json:"uid"`
	GroupType   string `json:"groupType"`
	FighterType string `json:"fighterType"`
	Ready       string `json:"ready"`
}

type RoomPlayerReadyParams struct {
	Rid int `json:"rid"`
	Uid int `json:"uid"`
}

type BeginParam struct {
	Rid int `json:"rid"`
}

type KickPlayerParams struct {
	Rid      int    `json:"rid"`
	Username string `json:"username"`
}
