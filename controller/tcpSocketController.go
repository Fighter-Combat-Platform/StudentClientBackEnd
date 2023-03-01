package controller

import (
	"awesomeProject/model"
	"bytes"
	"log"
	"net"
	"os"
	"time"
)

//var conn *net.TCPConn

func InitTcpSocket() (*net.TCPConn, error) {
	server := "10.119.6.138:35030"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		println(os.Stderr, "Fatal error: %s", err.Error())
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println(os.Stderr, "Fatal error: %s", err.Error())
	}

	println("tcp connection success")

	return conn, err
}

func Int32ToBytes(i int32) []byte {
	buf := make([]byte, 4)
	for k := 0; k < 4; k++ {
		buf[k] = byte(int8(i))
		i = i >> 8
	}
	return buf
}

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func Sender(conn net.Conn, param model.Begin) (bool, string) {

	var data []byte
	var user []byte
	for i := 0; i < len(param.UsersInfo); i++ {
		user = BytesCombine(user, Int32ToBytes(int32(param.UsersInfo[i].Uid)),
			Int32ToBytes(int32(param.UsersInfo[i].GroupType)),
			Int32ToBytes(int32(param.UsersInfo[i].FighterType)))
	}

	data = BytesCombine(Int32ToBytes(param.OperationType), Int32ToBytes(param.DaotiaoType), Int32ToBytes(int32(time.Now().Unix())),
		Int32ToBytes(param.ContentLength), Int32ToBytes(param.CallbackID), Int32ToBytes(int32(param.Room.Rid)), Int32ToBytes(int32(param.Room.Mid)),
		Int32ToBytes(int32(param.Room.MapID)), Int32ToBytes(int32(param.Room.WeatherType)), Int32ToBytes(int32(param.Room.LightType)),
		Int32ToBytes(100), Int32ToBytes(int32(param.Room.PlayerCount)), user)

	_, err := conn.Write(data)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "waiting server back msg error: ", err)
		return false, "战斗服务器创建房间失败"
	}

	println("send over")
	println(data)

	buffer := make([]byte, 2048)

	_, err = conn.Read(buffer)
	if err != nil {
		log.Println(conn.RemoteAddr().String(), "waiting server back msg error: ", err)
		return false, "战斗服务器创建房间失败"
	}
	log.Println(conn.RemoteAddr().String(), "receive server back msg: ", string(buffer))
	return true, "战斗服务器创建房间成功"

}
