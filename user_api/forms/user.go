package forms

type PasswordLoginForm struct {
	Mobile    string `forms:"mobile" json:"mobile" binding:"required,mobile"`
	Password  string `forms:"password" json:"password" binding:"required,min=6,max=20"`
	Captcha   string `forms:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `forms:"captcha_id" json:"captcha_id" binding:"required"`
}

type SmsForm struct {
	Mobile string `forms:"mobile" json:"mobile" binding:"required,mobile"`
	Type   uint   `forms:"type" json:"type" binding:"required,oneof=1 2"`
}

type RegisterForm struct {
	Mobile   string `forms:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `forms:"password" json:"password" binding:"required,min=6,max=20"`
	Code     string `forms:"code" json:"code" binding:"required,min=5,max=5"`
}
