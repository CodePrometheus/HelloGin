package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
)

func main() {

	engine := gin.Default()

	//设置html目录
	engine.LoadHTMLGlob("./html/*")


	engine.Static("/img","./img")

	engine.GET("/hellohtml", func(context *gin.Context) {
		fullPath := "请求路径：" + context.FullPath()
		fmt.Println(fullPath)

		context.HTML(http.StatusOK, "index.html", gin.H{
			"fullPath": fullPath,
			"title":    "gin教程",
		})
	})
	
	engine.Run(":8090")
}
