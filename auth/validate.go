package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/martzing/simple-wallet/helpers"
)

func registerValidate(c *gin.Context) (*RegisterParams, helpers.CustomError) {
	validate := validator.New()
	validateStruct := new(RegisterParams)

	if err := c.Bind(validateStruct); err != nil {
		msg := err.Error()
		code := http.StatusBadRequest
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	if err := validate.Struct(validateStruct); err != nil {
		msg := err.Error()
		code := http.StatusBadRequest
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	return validateStruct, nil
}

func loginValidate(c *gin.Context) (*LoginParams, helpers.CustomError) {
	validate := validator.New()
	validateStruct := new(LoginParams)

	if err := c.Bind(validateStruct); err != nil {
		msg := err.Error()
		code := http.StatusBadRequest
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	if err := validate.Struct(validateStruct); err != nil {
		msg := err.Error()
		code := http.StatusBadRequest
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	return validateStruct, nil
}
