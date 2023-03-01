package service

import (
	"awesomeProject/dao"
	"awesomeProject/model"
	"awesomeProject/utils"
)

func Login(params utils.LoginParams) (success bool, msg string, data int) {
	query := dao.GetUserByUsername(params.Username)
	if query.Username == "" {
		success, msg = false, "用户不存在"
	} else if query.Password != params.Password {
		success, msg = false, "密码错误"
	} else {
		query.State = 2
		dao.ModifyUser(query)
		success, msg, data = true, "登录成功", query.Uid
	}
	return

}

func Register(params utils.RegisterParams) (success bool, msg string, data string) {
	query := dao.GetUserByUsername(params.Username)
	if query.Username == "" {
		//此处需要考虑update喝get操作的原子性
		dao.UpdateTicket()
		count := dao.GetTicket()

		dao.CreateUser(model.User{
			Uid:      count,
			Username: params.Username,
			Password: params.Password,
			SID:      params.SID,
			Role:     0,
			State:    1,
			Point:    0,
		})

		var roomName = params.Username + "的房间"
		dao.CreateRoom(model.Room{
			Rid:             count,
			Mid:             0,
			State:           1,
			Capacity:        4,
			PlayerCount:     0,
			RoomName:        roomName,
			OwnerName:       params.Username,
			GeneralSetting:  0,
			MaxCampCount:    2,
			MaxWatcherCount: 10,
			MapID:           0,
			WeatherType:     0,
			LightType:       0,
			IfPublic:        1,
			Password:        "123456",
			Type:            "1v1",
		})
		success, msg, data = true, "注册成功", ""
	} else {
		success, msg, data = false, "用户已存在", ""
	}
	return
}

func ModifyUserPassword(params utils.ModifyUserPasswordParams) (success bool, msg string, data string) {
	query := dao.GetUserByUid(params.Uid)
	if query.Username == "" {
		success, msg = false, "用户不存在"
	} else {
		query.Password = params.NewPassword
		dao.ModifyUser(query)
		success, msg = true, "修改用户密码成功"
	}
	return
}

func ModifyUsername(params utils.ModifyUsernameParams) (success bool, msg string, data string) {
	query := dao.GetUserByUsername(params.NewUsername)
	user := dao.GetUserByUid(params.Uid)
	if query.Username != "" {
		success, msg = false, "用户名已存在"
	} else {
		user.Username = params.NewUsername
		dao.ModifyUser(user)
		success, msg = true, "修改用户名成功"
	}
	return
}
