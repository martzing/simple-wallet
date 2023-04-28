package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func registerValidate(c *gin.Context) *RegisterParams {
	validate := validator.New()
	validateStruct := new(RegisterParams)

	if err := c.Bind(validateStruct); err != nil {
		panic(err.Error())
	}

	if err := validate.Struct(validateStruct); err != nil {
		panic(err.Error())
	}

	return validateStruct
}

func loginValidate(c *gin.Context) *LoginParams {
	validate := validator.New()
	validateStruct := new(LoginParams)

	if err := c.Bind(validateStruct); err != nil {
		panic(err.Error())
	}

	if err := validate.Struct(validateStruct); err != nil {
		panic(err.Error())
	}

	return validateStruct
}
