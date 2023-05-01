package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type RegisterParams struct {
	Username string `form:"username" json:"username" binding:"required,alphanum"`
	Password string `form:"password" json:"password" binding:"required,alphanum"`
	Email    string `form:"email" json:"email" binding:"required,email"`
}

type RegisterRes struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type LoginParams struct {
	Username string `form:"username" json:"username" binding:"required,alphanum"`
	Password string `form:"password" json:"password" binding:"required,alphanum"`
}

type LoginRes struct {
	Token string `json:"token"`
}

type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}
