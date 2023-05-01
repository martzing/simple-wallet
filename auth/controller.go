package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martzing/simple-wallet/helpers"
)

func Register(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	data, err := registerValidate(c)

	if err != nil {
		helpers.AbortError(c, err)
		return
	}

	res, _err := register(data)

	if _err != nil {
		helpers.AbortError(c, _err)
		return
	}

	c.JSON(http.StatusCreated, res)
}

func Login(c *gin.Context) {

	c.Header("Content-Type", "application/json")

	data, err := loginValidate(c)

	if err != nil {
		helpers.AbortError(c, err)
		return
	}

	res, _err := login(data)

	if _err != nil {
		helpers.AbortError(c, _err)
		return
	}

	c.JSON(http.StatusCreated, res)
}
