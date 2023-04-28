package auth

import (
	authDB "github.com/martzing/simple-wallet/auth/db"
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/models"
	"golang.org/x/crypto/bcrypt"
)

func register(data *RegisterData) RegisterRes {
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
		panic("User already register")
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
