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

func register(data *RegisterParams) (*RegisterRes, helpers.CustomError) {

	duplicateUser, err := authDB.GetUser(db.DB, data.Username)
	if duplicateUser != nil {
		msg := "User already register"
		code := http.StatusConflict
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	tokens, err := authDB.GetTokens(db.DB)

	if err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 8)

	passHash := string(hashed)

	dbTxn := db.NewTransaction()
	DB, err := dbTxn.Begin(db.REPEATABLE_READ)

	if err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	user, err := authDB.CreateUser(DB, &models.User{
		Username: data.Username,
		Password: passHash,
		Email:    data.Email,
		Role:     "user",
		IsActive: true,
	})

	if err != nil {
		dbTxn.Rollback()
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	wallets := []*models.Wallet{}
	for _, token := range tokens {
		wallets = append(wallets, &models.Wallet{
			Balance: 0,
			TokenID: token.ID,
			UserID:  user.ID,
		})
	}

	_, _err := authDB.CreateWallets(DB, wallets)

	if _err != nil {
		dbTxn.Rollback()
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: _err,
		}
		return nil, ce
	}

	dbTxn.Commit()

	return &RegisterRes{
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func login(data *LoginParams) (*LoginRes, helpers.CustomError) {

	user, err := authDB.GetUser(db.DB, data.Username)
	if user == nil {
		msg := "User not found"
		code := http.StatusNotFound
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		msg := "Password is incorrect"
		code := http.StatusUnauthorized
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
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
			Err: err,
		}
		return nil, ce
	}

	return &LoginRes{Token: tokenString}, nil
}
