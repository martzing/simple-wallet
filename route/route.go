package route

import (
	"github.com/martzing/simple-wallet/simple-wallet/constants"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	strings := []Route{
		{
			Method: "GET",
			Path:   "/",
			Handler: func(c *gin.Context) {
				c.JSON(constants.HttpStatus["OK"], gin.H{"message": "OK"})
			},
		},
	}
	for _, s := range strings {
		r.Handle(s.Method, s.Path, s.Handler)
	}
}
