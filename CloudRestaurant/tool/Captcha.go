package tool

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type CaptchaResult struct {
	Id          string `json:"id"`
	Base64Blob  string `json:"base_64_blob"`
	VerifyValue string `json:"code"`
}

var store base64Captcha.Store = RedisStore{}

// GenerateCaptcha 生成图形化验证码
func GenerateCaptcha(ctx *gin.Context) error {
	var driver base64Captcha.Driver
	var driverString base64Captcha.DriverString

	captchaConfig := base64Captcha.DriverString{
		Height:          30,
		Width:           60,
		NoiseCount:      0,
		ShowLineOptions: 0,
		Length:          4,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 254,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}

	driverString = captchaConfig
	driver = driverString.ConvertFonts()
	captcha := base64Captcha.NewCaptcha(driver, store)
	captchaId, b64s, err := captcha.Generate()
	if err != nil {
		return err
	}
	captchaResult := CaptchaResult{Id: captchaId, Base64Blob: b64s}
	Success(ctx, map[string]interface{}{
		"captcha_result": captchaResult,
	})
	return nil
}

func VerifyCaptcha(id string, value string) bool {
	return store.Verify(id, value, false)
}
