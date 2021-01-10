package main

import (
	"fmt"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

type User struct {
	Account  string `json:"account" validate:"required"`
	PassWord string `json:"pass_word" validate:"required"`
}

func main() {
	//获取数据
	account := ""
	pass_word := ""
	user := &User{}
	user.Account = account
	user.PassWord = pass_word

	//验证
	zh_ch := zh.New()
	validate := validator.New()
	uni := ut.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")
	//验证器注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Translate(trans))
		}
		return
	}
}
