package validator

import (
	AppError "github.com/dorlib/todo-list-manager/libs/microservice-development-kit/error"
	"strings"

	eng "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/rotisserie/eris"

	"github.com/go-playground/validator/v10/translations/en"
)

type RequestValidator struct {
	validate   *validator.Validate
	translator ut.Translator
}

func NewRequestValidator() (*RequestValidator, error) {
	translator := eng.New()
	uni := ut.New(translator, translator)
	trans, found := uni.GetTranslator("en")
	if !found {
		return nil, eris.Wrap(AppError.NewInvalidRequest("could not find universal translator"), "validation error")
	}

	validate := validator.New()
	if err := en.RegisterDefaultTranslations(validate, trans); err != nil {
		return nil, err
	}

	return &RequestValidator{
		validate:   validate,
		translator: trans,
	}, nil
}

func (r *RequestValidator) Validate(obj interface{}) error {
	err := r.validate.Struct(obj)
	if err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Translate(r.translator))
		}

		return eris.Wrap(AppError.NewInvalidRequest(strings.Join(errors, ", ")), "validation error")
	}

	return nil
}
