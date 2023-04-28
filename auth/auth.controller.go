package auth

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
		}
	}()

	c.Header("Content-Type", "application/json")
	data := registerValidate(c)

	res := register(data)

	c.JSON(http.StatusOK, res)
}

func Login(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("%v", err),
			})
		}
	}()

	c.Header("Content-Type", "application/json")
	data := loginValidate(c)

	res := login(data)

	c.JSON(http.StatusOK, res)
}
