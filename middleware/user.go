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
	authorization := c.Request.Header["Authorization"]
	if len(authorization) < 1 {
		msg := "Authorization is missing"
		code := http.StatusUnauthorized
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    &msg,
			StatusCode: &code,
		}
		helpers.AbortError(c, ce)
		return
	}

	tokenString := strings.Replace(authorization[0], "Bearer ", "", 1)

	claims := &auth.Claims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(*configs.JwtSecret), nil
	}, jwt.WithLeeway(5*time.Second))

	if err != nil {
		msg := strings.TrimSpace(strings.Split(err.Error(), ":")[1])
		code := http.StatusUnauthorized
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    &msg,
			StatusCode: &code,
		}
		helpers.AbortError(c, ce)
		return
	}

	c.Set("userId", claims.UserID)
	c.Next()
}
