package controller

import (
	"github.com/gin-gonic/gin"
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"fmt"
)

type ShopController struct {
}

/**
 * shop模块的路由解析
 */
func (sc *ShopController) Router(app *gin.Engine) {
	app.GET("/api/shops", sc.GetShopList)
}

/**
 * 获取商铺列表
 */
func (sc *ShopController) GetShopList(context *gin.Context) {

	longitude := context.Query("longitude")
	latitude := context.Query("latitude")

	fmt.Println(longitude, latitude)
	if longitude == "" || longitude == "undefined" || latitude == "" || latitude == "undefined" {
		longitude = "116.34" //北京
		latitude = "40.34"
	}

	fmt.Println(longitude, latitude)

	shopService := service.ShopService{}
	shops := shopService.ShopList(longitude, latitude)
	if len(shops) != 0 {
		tool.Success(context, shops)
		return
	}
	tool.Failed(context, "暂未获取到商户信息")
}
