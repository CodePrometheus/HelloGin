package controller

import (
	"github.com/gin-gonic/gin"
	"CloudRestaurant/service"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.sendSmsCode)
}

// http://localhost:8090/api/sendcode?phone=13167582436
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	//发送验证码
	phone, exist := context.GetQuery("phone")
	if !exist {
		context.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "参数解析失败",
		})
		return
	}

	ms := service.MemberService{}
	isSend := ms.Sendcode(phone)
	if isSend {
		context.JSON(200, map[string]interface{}{
			"code": 1,
			"msg":  "发送成功",
		})
		return
	}

	context.JSON(200, map[string]interface{}{
		"code": 0,
		"msg":  "发送失败",
	})

}
