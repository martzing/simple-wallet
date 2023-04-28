package authDB

import (
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/models"
)

func CreateUser(dbTxn db.DatabaseTransaction, data *CreateUserParams) *models.User {

	db := dbTxn.Get()

	user := models.User{
		Username: data.Username,
		Password: data.Password,
		Email:    data.Email,
		Role:     data.Role,
		IsActive: data.IsActive,
	}

	err := db.Create(&user).Error

	if err != nil {
		panic(err)
	}

	return &user
}

func CreateWallets(dbTxn db.DatabaseTransaction, wallets []*models.Wallet) []*models.Wallet {

	db := dbTxn.Get()

	err := db.Create(&wallets).Error

	if err != nil {
		panic(err)
	}

	return wallets
}
