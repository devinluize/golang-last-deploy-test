package validation

import (
	"errors"
	"fmt"
	"net/http"
	"user-services/api/exceptions"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate *validator.Validate
)

func translateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}
func ValidationForm(writer http.ResponseWriter, request *http.Request, form interface{}) *exceptions.BaseErrorResponse {
	validate = validator.New()
	var msg string

	err := validate.Struct(form)
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			if err.Tag() == "required" {
				msg = fmt.Sprintf("%s cant be empty", err.Field())
			} else if err.Tag() == "email" {
				msg = fmt.Sprintf("%s format not matched", err.Field())
			} else if err.Tag() == "noWhiteSpace" {
				msg = fmt.Sprintf("%s remove white space", err.Field())
			} else if err.Tag() == "eqfield" {
				msg = fmt.Sprintf("%s should matched", err.Field())
			} else if err.Tag() == "nefield" {
				msg = fmt.Sprintf("%s shouldn't matched", err.Field())
			}
		}
	}
	if msg != "" {
		return &exceptions.BaseErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    msg,
			Err:        errors.New(msg),
		}
	} else if err != nil {
		errorMsg := fmt.Sprintf("%v ", translateError(err, trans))
		return &exceptions.BaseErrorResponse{
			StatusCode: http.StatusBadRequest,
			Message:    errorMsg,
			Err:        errors.New(errorMsg),
		}
	}

	return nil
}
