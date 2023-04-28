package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func registerValidate(c *gin.Context) *RegisterData {
	validate := validator.New()
	validateStruct := new(RegisterData)

	if err := c.Bind(validateStruct); err != nil {
		panic(err.Error())
	}

	if err := validate.Struct(validateStruct); err != nil {
		panic(err.Error())
	}

	return validateStruct
}
