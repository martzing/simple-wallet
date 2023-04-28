package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/martzing/simple-wallet/auth"
	"github.com/martzing/simple-wallet/configs"
)

func UserMiddleware(c *gin.Context) {
	authorization := c.Request.Header["Authorization"]
	if len(authorization) < 1 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization is missing",
		})
		return
	}

	tokenString := strings.Replace(authorization[0], "Bearer ", "", 1)

	claims := &auth.Claims{}

	token, _ := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(*configs.JwtSecret), nil
	}, jwt.WithLeeway(5*time.Second))

	claims, ok := token.Claims.(*auth.Claims)

	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token is expired",
		})
		return
	} else if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}
	c.Set("userId", claims.UserID)
	c.Next()
}
