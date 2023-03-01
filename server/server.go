package server

import (
	"awesomeProject/repository"
	"awesomeProject/router"
	"github.com/kataras/iris/v12"
)

func NewApp() *iris.Application {
	app := iris.New()

	//controller.InitTcpSocket()

	router.SetRouter(app)

	repository.InitDBConn()

	return app
}
