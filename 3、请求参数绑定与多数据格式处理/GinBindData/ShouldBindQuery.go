package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	engine := gin.Default()

	// http://localhost:8080/hello?name=davie&classes=软件工程
	engine.GET("/hello", func(context *gin.Context) {
		fmt.Println(context.FullPath())

		var student Student
		err := context.ShouldBindQuery(&student)
		if err != nil {
			log.Fatal(err.Error())
			return
		}

		fmt.Println(student.Name)
		fmt.Println(student.Classes)
		context.Writer.Write([]byte("hello," + student.Name))

	})

	engine.Run()
}

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}
