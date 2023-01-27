package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()

	engine.POST("/addstudent", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		var person Person
		if err := context.BindJSON(&person); err != nil {
			log.Fatal(err.Error())
		}

		fmt.Println("姓名:", person.Name)
		fmt.Println("年龄:", person.Age)
		context.Writer.Write([]byte("添加记录：" + person.Name))
	})

	engine.Run()
}

type Person struct {
	Name string `form:"name"`
	Sex  string `form:"sex"`
	Age  int    `from:"age"`
}
