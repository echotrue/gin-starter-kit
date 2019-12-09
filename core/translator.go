package core

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zh2 "gopkg.in/go-playground/validator.v9/translations/zh"
)

type TranslateStruct struct{}

var (
	translateInstance *TranslateStruct
	uni               *ut.UniversalTranslator
	translate         ut.Translator
)

func TsInstance() *TranslateStruct {
	once.Do(func() {
		translateInstance = &TranslateStruct{}
	})
	return translateInstance
}

func (ts *TranslateStruct) NewTs() {
	zh_cn := zh.New()
	uni := ut.New(zh_cn)
	translate, _ = uni.GetTranslator("zh")

	validate := validator.New()

	zh2.RegisterDefaultTranslations(validate, translate)
}

func (ts *TranslateStruct) GetTranslate() ut.Translator {
	return translate
}
