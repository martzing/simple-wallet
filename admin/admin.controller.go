package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martzing/simple-wallet/helpers"
)

func CreateToken(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			helpers.AbortError(c, err)
			return
		}
	}()
	c.Header("Content-Type", "application/json")
	data := createTokenValidate(c)
	res := createToken(data)

	c.JSON(http.StatusCreated, res)
}

func UpdateToken(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			helpers.AbortError(c, err)
			return
		}
	}()
	c.Header("Content-Type", "application/json")
	data := updateTokenValidate(c)
	res := updateToken(data)

	c.JSON(http.StatusOK, res)
}

func DeleteToken(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			helpers.AbortError(c, err)
			return
		}
	}()
	c.Header("Content-Type", "application/json")
	tokenId := deleteTokenValidate(c)
	res := deleteToken(tokenId)

	c.JSON(http.StatusOK, res)
}

func AddTokenBalance(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			helpers.AbortError(c, err)
			return
		}
	}()
	c.Header("Content-Type", "application/json")
	data := updateTokenBalanceValidate(c)
	res := updateTokenBalance(data, "add")

	c.JSON(http.StatusOK, res)
}

func MinusTokenBalance(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			helpers.AbortError(c, err)
			return
		}
	}()
	c.Header("Content-Type", "application/json")
	data := updateTokenBalanceValidate(c)
	res := updateTokenBalance(data, "minus")

	c.JSON(http.StatusOK, res)
}

func GetTokenBalance(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			helpers.AbortError(c, err)
			return
		}
	}()
	res := getTokenBalance()

	c.JSON(http.StatusOK, res)
}
