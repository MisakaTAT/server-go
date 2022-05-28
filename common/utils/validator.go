package utils

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translation "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

func init() {
	translator := zh.New()
	uni = ut.New(translator, translator)
	trans, _ = uni.GetTranslator("zh")
	validate := binding.Validator.Engine().(*validator.Validate)
	_ = translation.RegisterDefaultTranslations(validate, trans)
}

func Translate(err error) string {
	var result string
	errors := err.(validator.ValidationErrors)
	for _, err := range errors {
		result += err.Translate(trans) + ";"
	}
	return result
}
