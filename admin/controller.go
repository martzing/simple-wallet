package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martzing/simple-wallet/helpers"
)

func CreateToken(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	data, err := createTokenValidate(c)

	if err != nil {
		helpers.AbortError(c, err)
		return
	}

	res, _err := createToken(data)

	if _err != nil {
		helpers.AbortError(c, _err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func UpdateToken(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	data, err := updateTokenValidate(c)

	if err != nil {
		helpers.AbortError(c, err)
		return
	}

	res, _err := updateToken(data)

	if _err != nil {
		helpers.AbortError(c, _err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func DeleteToken(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	tokenId, err := deleteTokenValidate(c)

	if err != nil {
		helpers.AbortError(c, err)
		return
	}

	res, _err := deleteToken(*tokenId)

	if _err != nil {
		helpers.AbortError(c, _err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func AddTokenBalance(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	data, err := updateTokenBalanceValidate(c)

	if err != nil {
		helpers.AbortError(c, err)
		return
	}

	res, _err := updateTokenBalance(data, "add")

	if _err != nil {
		helpers.AbortError(c, _err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func MinusTokenBalance(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	data, err := updateTokenBalanceValidate(c)

	if err != nil {
		helpers.AbortError(c, err)
		return
	}

	res, _err := updateTokenBalance(data, "minus")

	if _err != nil {
		helpers.AbortError(c, _err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func GetTokenBalance(c *gin.Context) {
	res, err := getTokenBalance()

	if err != nil {
		helpers.AbortError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
