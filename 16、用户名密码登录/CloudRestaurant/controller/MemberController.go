package controller

import (
	"github.com/gin-gonic/gin"
	"CloudRestaurant/service"
	"CloudRestaurant/param"
	"CloudRestaurant/tool"
	"fmt"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {

	engine.GET("/api/sendcode", mc.sendSmsCode)

	engine.POST("/api/login_sms", mc.smsLogin)

	engine.GET("/api/captcha", mc.captcha)

	//postman测试
	engine.POST("/api/vertifycha", mc.vertifyCaptcha)

	//login_pwd
	engine.POST("/api/login_pwd", mc.nameLogin)
}

//用户名+密码、验证码登录
func (mc *MemberController) nameLogin(context *gin.Context) {

	//1、解析用户登录传递参数
	var loginParam param.LoginParam
	err := tool.Decode(context.Request.Body, &loginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}

	//2、验证验证码
	validate := tool.VertifyCaptcha(loginParam.Id, loginParam.Value)
	if !validate {
		tool.Failed(context, "验证码不正确，请重新验证")
		return
	}

	//3、登录
	ms := service.MemberService{}
	member := ms.Login(loginParam.Name, loginParam.Password)
	if member.Id != 0 {
		tool.Success(context, &member)
		return
	}

	tool.Failed(context, "登录失败")
}

//生成验证码
func (mc *MemberController) captcha(context *gin.Context) {
	tool.GenerateCaptcha(context)
}

//验证验证码是否正确
func (mc *MemberController) vertifyCaptcha(context *gin.Context) {
	var captcha tool.CaptchaResult
	err := tool.Decode(context.Request.Body, &captcha)
	if err != nil {
		tool.Failed(context, " 参数解析失败 ")
		return
	}

	result := tool.VertifyCaptcha(captcha.Id, captcha.VertifyValue)
	if result {
		fmt.Println("验证通过")
	} else {
		fmt.Println("验证失败")
	}
}

// http://localhost:8090/api/sendcode?phone=13167582436
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	//发送验证码
	phone, exist := context.GetQuery("phone")
	if !exist {
		tool.Failed(context, "参数解析失败")
		return
	}

	ms := service.MemberService{}
	isSend := ms.Sendcode(phone)
	if isSend {
		tool.Success(context, "发送成功")
		return
	}

	tool.Failed(context, "发送失败")
}

//手机号+短信  登录的方法
func (mc *MemberController) smsLogin(context *gin.Context) {

	var smsLoginParam param.SmsLoginParam
	err := tool.Decode(context.Request.Body, &smsLoginParam)

	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}

	//完成手机+验证码登录
	us := service.MemberService{}
	member := us.SmsLogin(smsLoginParam)

	if member != nil {
		tool.Success(context, member)
		return
	}

	tool.Failed(context, "登录失败")
}
