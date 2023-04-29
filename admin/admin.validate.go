package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/martzing/simple-wallet/helpers"
)

func createTokenValidate(c *gin.Context) *CreateTokenParams {
	validate := validator.New()
	validateStruct := new(CreateTokenParams)

	if err := c.Bind(validateStruct); err != nil {
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		panic(ce)
	}

	if err := validate.Struct(validateStruct); err != nil {
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		panic(ce)
	}

	return validateStruct
}

func updateTokenValidate(c *gin.Context) *UpdateTokenParams {
	validate := validator.New()
	validateStruct := new(UpdateTokenParams)
	if err := c.Bind(validateStruct); err != nil {
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		panic(ce)
	}

	if err := validate.Struct(validateStruct); err != nil {
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		panic(ce)
	}

	return validateStruct
}

func deleteTokenValidate(c *gin.Context) int {
	validate := validator.New()

	tokenId, err := strconv.Atoi(c.Param("token_id"))

	if err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		panic(ce)
	}

	if err := validate.Var(tokenId, "required,numeric,min=1"); err != nil {
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		panic(ce)
	}

	return tokenId
}

func updateTokenBalanceValidate(c *gin.Context) *UpdateTokenBalanceParams {
	validate := validator.New()
	validateStruct := new(UpdateTokenBalanceParams)

	if err := c.Bind(validateStruct); err != nil {
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		panic(ce)
	}

	if err := validate.Struct(validateStruct); err != nil {
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		panic(ce)
	}

	return validateStruct
}
