package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martzing/simple-wallet/helpers"
)

func GetTokens(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			helpers.AbortError(c, err)
			return
		}
	}()
	res := getTokens()

	c.JSON(http.StatusOK, res)
}

func GetToken(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			helpers.AbortError(c, err)
			return
		}
	}()
	tokenId := getTokenValidate(c)
	res := getToken(tokenId)

	c.JSON(http.StatusOK, res)
}
