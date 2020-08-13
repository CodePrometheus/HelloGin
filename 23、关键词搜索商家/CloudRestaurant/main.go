package main

import (
	"github.com/gin-gonic/gin"
	"CloudRestaurant/tool"
	"github.com/goes/logger"
	"CloudRestaurant/controller"
	"strings"
	"fmt"
	"net/http"
)

type Data struct {
	Stream   string    `json:"Stream"`
	Messages []Message `json:"Messages"`
}

type Message struct {
	ID     string `json:"ID"`
	Values Value  `json:"Values"`
}

type Value struct {
	OrderId string `json:"order_id"`
}

func main() {

	cfg, err := tool.ParseConfig("./config/app.json")

	if err != nil {
		panic(err.Error())
	}

	//实例化数据库
	_, err = tool.OrmEngine(cfg)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	//初始化redis配置
	tool.InitRedisStore()

	app := gin.Default()

	//设置全局跨域访问
	app.Use(Cors())

	//集成session
	tool.InitSession(app)

	registerRouter(app)

	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}

//路由设置
func registerRouter(router *gin.Engine) {
	new(controller.HelloController).Router(router)
	new(controller.MemberController).Router(router)
	new(controller.FoodCategoryController).Router(router)
	new(controller.ShopController).Router(router)
}

//跨域访问：cross  origin resource share
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headerKeys []string
		for key, _ := range context.Request.Header {
			headerKeys = append(headerKeys, key)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false")
			context.Set("content-type", "application/json") //// 设置返回格式是json
		}

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}

		//处理请求
		context.Next()
	}
}
