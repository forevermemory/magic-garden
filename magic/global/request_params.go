package global

// RegisteruserParams 注册传入的参数
type RegisteruserParams struct {
	ID          int    `json:"user_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
	Phone       string `json:"phone"`
	Code        string `json:"code"`
	// Captcha 触发验证码后需要判断
	Captcha   string `json:"captcha"`
	IsCaptcha int    `json:"is_captcha"`
}
