package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	engin := gin.Default()

	//post register
	routerGroup := engin.Group("/user")
	routerGroup.POST("/register", registerHandle)
	routerGroup.POST("/login", loginHandle)
	routerGroup.DELETE("/:id", deleteHandle)

	engin.Run(":9001")
}

func registerHandle(context *gin.Context) {
	fullPath := "用户注册" + context.FullPath()
	fmt.Println(fullPath)
	context.Writer.WriteString(fullPath)
}

func loginHandle(context *gin.Context) {
	fullPath := "用户登录" + context.FullPath()
	fmt.Println(fullPath)
	context.Writer.WriteString(fullPath)
}

func deleteHandle(context *gin.Context) {
	fullPath := "用户删除" + context.FullPath()
	userID := context.Param("id")
	fmt.Println(fullPath + "  " + userID)
	context.Writer.WriteString(fullPath + "  " + userID)
}
