package dao

import (
	"awesomeProject/repository"
)

func GetTicket() int {
	return repository.GetTicket()
}

func UpdateTicket() {
	repository.UpdateTicket()
}
