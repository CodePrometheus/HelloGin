package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	engine := gin.Default()

	// http://localhost:8080/hello?name=davie
	engine.Handle("GET", "/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Println(path)

		name := context.DefaultQuery("name", "hello")
		fmt.Println(name)

		//输出
		context.Writer.Write([]byte("Hello ," + name))
	})

	//post
	// http://loclahost:8080/login
	engine.Handle("POST", "/login", func(context *gin.Context) {

		fmt.Println(context.FullPath())

		username := context.PostForm("username")
		password := context.PostForm("password")

		fmt.Println(username)
		fmt.Println(password)

		context.Writer.Write([]byte(username + " 登录 "))
	})

	engine.Run()
}
