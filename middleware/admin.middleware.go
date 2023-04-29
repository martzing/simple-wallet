package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/martzing/simple-wallet/auth"
	"github.com/martzing/simple-wallet/configs"
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/helpers"
	userDB "github.com/martzing/simple-wallet/user/db"
)

func AdminMiddleware(c *gin.Context) {
	dbTxn := db.NewTransaction()
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
	dbTxn.Begin(db.REPEATABLE_READ)
	user := userDB.GetUser(dbTxn, claims.UserID)

	if user == nil {
		dbTxn.Rollback()
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "User not found",
			StatusCode: http.StatusNotFound,
		}
		panic(ce)
	}
	if user.Role != "admin" {
		dbTxn.Rollback()
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "You don't have permission for this API",
			StatusCode: http.StatusForbidden,
		}
		panic(ce)
	}
	c.Set("userId", claims.UserID)
	dbTxn.Commit()
	c.Next()
}
