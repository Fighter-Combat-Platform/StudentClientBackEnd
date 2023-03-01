package utils

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

//定义一个websocket处理器，用于收集消息和广播消息
type _Hub struct {
	//用户列表，保存所有用户
	userList map[*User]bool
	//注册chan，用户注册时添加到chan中
	Register chan *User
	//注销chan，用户退出时添加到chan中，再从map中删除
	Unregister chan *User
	//广播消息，将消息广播给所有连接
	broadcast chan []byte
}

//定义一个websocket连接对象，连接中包含每个连接的信息
type User struct {
	Id   string
	Conn *websocket.Conn
	Msg  chan []byte
}

//定义一个升级器，将普通的http连接升级为websocket连接
var Up = &websocket.Upgrader{
	//定义读写缓冲区大小
	WriteBufferSize: 1024,
	ReadBufferSize:  1024,
	//校验请求
	CheckOrigin: func(r *http.Request) bool {
		//如果不是get请求，返回错误
		if r.Method != "GET" {
			fmt.Println("请求方式错误")
			return false
		}
		//如果路径中不包括chat，返回错误
		if r.URL.Path != "/roomInfoUpdate" {
			fmt.Println("请求路径错误")
			return false
		}
		//还可以根据其他需求定制校验规则
		return true
	},
}

//初始化处理中心，以便调用
var Hub = &_Hub{
	userList:   make(map[*User]bool),
	Register:   make(chan *User),
	Unregister: make(chan *User),
	broadcast:  make(chan []byte),
}

func Read(user *User) {
	//从连接中循环读取信息
	for {
		_, msg, err := user.Conn.ReadMessage()
		if err != nil {
			fmt.Println("用户退出:", user.Conn.RemoteAddr().String())
			Hub.Unregister <- user
			break
		}
		//将读取到的信息传入websocket处理器中的broadcast中，
		Hub.broadcast <- msg
	}
}

//func Write(user *User) {
//	for data := range user.Msg {
//		str := string(data)
//		str = strings.Replace(str, "\n", "", -1)
//		str = strings.Replace(str, "\r", "", -1)
//
//		if str == user.Id {
//			err := user.Conn.WriteMessage(1, data)
//			if err != nil {
//				fmt.Println("写入错误")
//				break
//			}
//		}
//	}
//}

func Write(user *User) {
	for data := range user.Msg {
		str := string(data)
		if str == user.Id {
			err := user.Conn.WriteMessage(1, data)
			if err != nil {
				fmt.Println("写入错误")
				break
			}
		}
	}
}

func Clear(rid string) {
	println("clear")
	for user := range Hub.Register {
		if rid == user.Id {
			Hub.Unregister <- user
		}
	}
}

//处理中心处理获取到的信息
func (h *_Hub) Run() {
	for {
		select {
		//从注册chan中取数据
		case user := <-h.Register:
			//取到数据后将数据添加到用户列表中
			h.userList[user] = true
		case user := <-h.Unregister:
			//从注销列表中取数据，判断用户列表中是否存在这个用户，存在就删掉
			if _, ok := h.userList[user]; ok {
				delete(h.userList, user)
			}
			num := 0
			for range h.userList {
				num++
			}
			print("socket count")
			println(num)
		case data := <-h.broadcast:
			//从广播chan中取消息，然后遍历给每个用户，发送到用户的msg中
			for u := range h.userList {
				select {
				case u.Msg <- data:
				default:
					delete(h.userList, u)
					close(u.Msg)
				}
			}
		}
	}
}
