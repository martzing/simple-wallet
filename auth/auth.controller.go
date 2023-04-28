package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/martzing/simple-wallet/helpers"
)

func Register(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			helpers.AbortError(c, err)
			return
		}
	}()
	c.Header("Content-Type", "application/json")
	data := registerValidate(c)

	res := register(data)

	c.JSON(http.StatusCreated, res)
}

func Login(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			helpers.AbortError(c, err)
			return
		}
	}()
	c.Header("Content-Type", "application/json")
	data := loginValidate(c)

	res := login(data)

	c.JSON(http.StatusCreated, res)
}
