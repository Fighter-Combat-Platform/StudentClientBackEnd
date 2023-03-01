package utils

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

func GetContextParams(ctx iris.Context, params interface{}) bool {
	if err := ctx.ReadJSON(params); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		_, _ = ctx.JSON(ResponseBean{
			Success: false,
			Msg:     "无效的参数格式",
			Data:    nil,
		})
		fmt.Println(err)
		return false
	}
	return true
}

func SendResponse(ctx iris.Context, success bool, msg string, data interface{}) {
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ResponseBean{
		Success: success,
		Msg:     msg,
		Data:    data,
	})
}
