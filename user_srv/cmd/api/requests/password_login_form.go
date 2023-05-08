package requests


import (
	"errors"
	"regexp"
)

type PasswordLogin struct {
	Email    string   `json:"email"       binding:"required"`  
	Password string   `json:"password"    binding:"required,min=3,max=20"` 
	CaptchaId string  `json:"captcha_id"  binding:"required"` 
	Code string       `json:"code"        binding:"required,min=6,max=6"` 
}

func (p PasswordLogin) Check() error {
	const emailPattern = "(.+)@(.+){2,}\\.(.+){2,}"

	ok, err := regexp.Match(emailPattern, []byte(p.Email))
	if !ok || err !=nil {
		return errors.New("不正确的邮箱地址")
	}
	
	return nil
}