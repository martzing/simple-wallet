package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/martzing/simple-wallet/helpers"
)

func getTokenValidate(c *gin.Context) (*int, helpers.CustomError) {
	validate := validator.New()

	tokenId, err := strconv.Atoi(c.Param("token_id"))

	if err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	err = validate.Var(tokenId, "required,numeric,min=1")

	if err != nil {
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

func getWalletValidate(c *gin.Context) (*int, helpers.CustomError) {
	validate := validator.New()
	userId := c.GetInt("userId")

	err := validate.Var(userId, "required,numeric,min=1")

	if err != nil {
		msg := err.Error()
		code := http.StatusBadRequest
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	return &userId, nil
}

func transferTokenValidate(c *gin.Context) (*TransferTokenParams, helpers.CustomError) {
	validate := validator.New()
	validateStruct := new(TransferTokenParams)

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

	err := validate.Struct(validateStruct)

	if err != nil {
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

func getTransferTokensValidate(c *gin.Context) (*int, helpers.CustomError) {
	validate := validator.New()
	userId := c.GetInt("userId")

	err := validate.Var(userId, "required,numeric,min=1")

	if err != nil {
		msg := err.Error()
		code := http.StatusBadRequest
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	return &userId, nil
}
