package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	engine := gin.Default()

	engine.GET("/hellojson", func(context *gin.Context) {
		fullPath := "请求路径：" + context.FullPath()
		fmt.Println(fullPath)

		context.JSON(200, map[string]interface{}{
			"code":    1,
			"message": "OK",
			"data":    fullPath,
		})
	})

	engine.GET("/jsonstruct", func(context *gin.Context) {
		fullPath := "请求路径：" + context.FullPath()
		fmt.Println(fullPath)

		resp := Response{Code: 1, Message: "OK", Data: fullPath}
		context.JSON(200, &resp)
	})

	engine.Run(":8090")
}

type Response struct {
	Code    int
	Message string
	Data    interface{}
}
