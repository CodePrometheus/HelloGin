package param

//手机号+验证码登录时 参数传递
type SmsLoginParam struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
}
