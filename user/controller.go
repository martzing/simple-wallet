package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martzing/simple-wallet/helpers"
)

func GetTokens(c *gin.Context) {
	res, err := getTokens()

	if err != nil {
		helpers.AbortError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetToken(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	tokenId, err := getTokenValidate(c)

	if err != nil {
		helpers.AbortError(c, err)
		return
	}

	res, _err := getToken(*tokenId)

	if _err != nil {
		helpers.AbortError(c, _err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetWallet(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	userId, err := getWalletValidate(c)

	if err != nil {
		helpers.AbortError(c, err)
		return
	}

	res, _err := getWallet(*userId)

	if _err != nil {
		helpers.AbortError(c, _err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func TransferToken(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	data, err := transferTokenValidate(c)

	if err != nil {
		helpers.AbortError(c, err)
		return
	}

	res, _err := transferToken(data)

	if _err != nil {
		helpers.AbortError(c, _err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func GetTransferTokens(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	userId, err := getTransferTokensValidate(c)

	if err != nil {
		helpers.AbortError(c, err)
		return
	}

	res, _err := getTransferTokens(*userId)

	if _err != nil {
		helpers.AbortError(c, _err)
		return
	}

	c.JSON(http.StatusOK, res)
}
