package validation

import (
	"github.com/go-playground/locales/fa"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	fa_translations "github.com/go-playground/validator/v10/translations/fa"
	"sync"
)

var (
	validate   *validator.Validate
	translator ut.Translator
	once       sync.Once
)

func Init() {
	once.Do(func() {
		validate = validator.New()
		faLocale := fa.New()
		uni := ut.New(faLocale, faLocale)
		translator, _ = uni.GetTranslator("fa")
		_ = fa_translations.RegisterDefaultTranslations(validate, translator)
	})
}

func Validator() *validator.Validate {
	if validate == nil {
		Init()
	}
	return validate
}

func TranslatorInstance() ut.Translator {
	if translator == nil {
		Init()
	}
	return translator
}
