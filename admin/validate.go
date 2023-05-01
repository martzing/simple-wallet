package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/martzing/simple-wallet/helpers"
)

func createTokenValidate(c *gin.Context) (*CreateTokenParams, helpers.CustomError) {
	validate := validator.New()
	validateStruct := new(CreateTokenParams)

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

func updateTokenValidate(c *gin.Context) (*UpdateTokenParams, helpers.CustomError) {
	validate := validator.New()
	validateStruct := new(UpdateTokenParams)
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

func deleteTokenValidate(c *gin.Context) (*int, helpers.CustomError) {
	validate := validator.New()

	tokenId, err := strconv.Atoi(c.Param("token_id"))

	if err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	if err := validate.Var(tokenId, "required,numeric,min=1"); err != nil {
		msg := err.Error()
		code := http.StatusBadRequest
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	return &tokenId, nil
}

func updateTokenBalanceValidate(c *gin.Context) (*UpdateTokenBalanceParams, helpers.CustomError) {
	validate := validator.New()
	validateStruct := new(UpdateTokenBalanceParams)

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
