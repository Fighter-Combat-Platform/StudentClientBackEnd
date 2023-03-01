package dao

import (
	"awesomeProject/model"
	"awesomeProject/repository"
)

func CreateUser(user model.User) {
	repository.CreateUser(user)
	return
}

func ModifyUser(user model.User) {
	repository.ModifyUser(user)
	return
}

func GetUserByUsername(username string) model.User {
	return repository.GetUserByUsername(username)
}

func GetIDByUsername(username string) int {
	return repository.GetIDByUsername(username)
}
func GetUserByUid(uid int) model.User {
	return repository.GetUserByUid(uid)
}
