package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/martzing/simple-wallet/helpers"
)

func registerValidate(c *gin.Context) *RegisterParams {
	validate := validator.New()
	validateStruct := new(RegisterParams)

	if err := c.Bind(validateStruct); err != nil {
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    err.Error(),
			StatusCode: http.StatusUnauthorized,
		}
		panic(ce)
	}

	if err := validate.Struct(validateStruct); err != nil {
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    err.Error(),
			StatusCode: http.StatusUnauthorized,
		}
		panic(ce)
	}

	return validateStruct
}

func loginValidate(c *gin.Context) *LoginParams {
	validate := validator.New()
	validateStruct := new(LoginParams)

	if err := c.Bind(validateStruct); err != nil {
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    err.Error(),
			StatusCode: http.StatusUnauthorized,
		}
		panic(ce)
	}

	if err := validate.Struct(validateStruct); err != nil {
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    err.Error(),
			StatusCode: http.StatusUnauthorized,
		}
		panic(ce)
	}

	return validateStruct
}
