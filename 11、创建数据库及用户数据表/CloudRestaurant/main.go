package main

import (
	"github.com/gin-gonic/gin"
	"CloudRestaurant/tool"
	"CloudRestaurant/controller"
	"github.com/goes/logger"
)

func main() {

	cfg, err := tool.ParseConfig("./config/app.json")

	if err != nil {
		panic(err.Error())
	}

	_, err = tool.OrmEngine(cfg)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	app := gin.Default()

	registerRouter(app)

	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//路由设置
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
}
