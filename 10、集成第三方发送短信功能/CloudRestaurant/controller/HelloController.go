package controller

import "github.com/gin-gonic/gin"

type HelloController struct {
}

func (hello *HelloController) Router(engine *gin.Engine) {
	engine.GET("/hello", hello.Hello)
}

//解析 /hello
func (hello *HelloController) Hello(context *gin.Context) {
	context.JSON(200, map[string]interface{}{
		"meeage": "hello cloudrestaurant",
	})
}
