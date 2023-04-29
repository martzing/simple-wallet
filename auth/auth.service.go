package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	authDB "github.com/martzing/simple-wallet/auth/db"
	"github.com/martzing/simple-wallet/configs"
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/helpers"
	"github.com/martzing/simple-wallet/models"
	"golang.org/x/crypto/bcrypt"
)

func register(data *RegisterParams) RegisterRes {
	dbTxn := db.NewTransaction()

	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()

	dbTxn.Begin()

	duplicateUser := authDB.GetUser(dbTxn, data.Username)
	if duplicateUser != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "User already register",
			StatusCode: http.StatusConflict,
		}
		panic(ce)
	}
	hashed, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 8)

	passHash := string(hashed)

	params := authDB.CreateUserParams{
		Username: data.Username,
		Password: passHash,
		Email:    data.Email,
		Role:     "user",
		IsActive: true,
	}
	user := authDB.CreateUser(dbTxn, &params)

	tokens := authDB.GetTokens(dbTxn)

	wallets := []*models.Wallet{}
	for _, token := range *tokens {
		wallets = append(wallets, &models.Wallet{
			Balance: 0,
			TokenID: token.ID,
			UserID:  user.ID,
		})
	}
	authDB.CreateWallets(dbTxn, wallets)
	dbTxn.Commit()

	return RegisterRes{
		Username: user.Username,
		Email:    user.Email,
	}
}

func login(data *LoginParams) LoginRes {
	dbTxn := db.NewTransaction()

	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()

	dbTxn.Begin()

	user := authDB.GetUser(dbTxn, data.Username)
	if user == nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "User not found",
			StatusCode: http.StatusNotFound,
		}
		panic(ce)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "Password is incorrect",
			StatusCode: http.StatusUnauthorized,
		}
		panic(ce)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)),
		},
	})

	tokenString, err := token.SignedString([]byte(*configs.JwtSecret))

	if err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    err.Error(),
			StatusCode: http.StatusUnauthorized,
		}
		panic(ce)
	}

	dbTxn.Commit()
	return LoginRes{Token: tokenString}
}
