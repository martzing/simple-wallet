package authController

import (
	"simple-wallet/constants"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.JSON(constants.HttpStatus["CREATED"], gin.H{
		"message": "Register success",
	})
}
