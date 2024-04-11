package validation

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type BaseResponseErrorValidation struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    any    `json:"value"`
	Message  string `json:"message"`
}

func ErrorProvider(err error) *[]BaseResponseErrorValidation {
	var BaseResponseErrorsList []BaseResponseErrorValidation
	var er *validator.ValidationErrors
	if errors.As(err, er) {
		var BaseResponseStruct BaseResponseErrorValidation
		for _, e := range err.(validator.ValidationErrors) {
			BaseResponseStruct.Property = e.Field()
			BaseResponseStruct.Tag = e.Tag()
			BaseResponseStruct.Value = e.Param()
			BaseResponseErrorsList = append(BaseResponseErrorsList, BaseResponseStruct)
		}

	}
	return &BaseResponseErrorsList
}
