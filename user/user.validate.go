package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/martzing/simple-wallet/helpers"
)

func getTokenValidate(c *gin.Context) int {
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

func getWalletValidate(c *gin.Context) int {
	validate := validator.New()
	userId := c.GetInt("userId")

	if err := validate.Var(userId, "required,numeric,min=1"); err != nil {
		var ce helpers.CustomError
		ce = &helpers.ValidateError{
			Message:    err.Error(),
			StatusCode: http.StatusBadRequest,
		}
		panic(ce)
	}

	return userId
}
