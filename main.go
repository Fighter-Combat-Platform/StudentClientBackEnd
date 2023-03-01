package main

import (
	"awesomeProject/server"
	"awesomeProject/utils"
	"github.com/kataras/iris/v12"
)

func main() {
	app := server.NewApp()
	go utils.Hub.Run()
	if err := app.Run(iris.Addr("0.0.0.0:8080"), iris.WithoutServerError(iris.ErrServerClosed)); err != nil {
		panic("Failed to Start Server!")
	}
	
}
