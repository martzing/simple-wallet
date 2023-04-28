package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/martzing/simple-wallet/auth"
	"github.com/martzing/simple-wallet/configs"
	"github.com/martzing/simple-wallet/helpers"
)

func UserMiddleware(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			helpers.AbortError(c, err)
			return
		}
	}()

	authorization := c.Request.Header["Authorization"]
	if len(authorization) < 1 {

		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "Authorization is missing",
			StatusCode: http.StatusUnauthorized,
		}
		panic(ce)
	}

	tokenString := strings.Replace(authorization[0], "Bearer ", "", 1)

	claims := &auth.Claims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(*configs.JwtSecret), nil
	}, jwt.WithLeeway(5*time.Second))

	if err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    strings.TrimSpace(strings.Split(err.Error(), ":")[1]),
			StatusCode: http.StatusUnauthorized,
		}
		panic(ce)
	}

	c.Set("userId", claims.UserID)
	c.Next()
}
