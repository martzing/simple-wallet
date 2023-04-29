package adminDB

import (
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/models"
)

func CreateToken(dbTxn db.DatabaseTransaction, token *models.Token) *models.Token {
	db := dbTxn.Get()

	err := db.Create(&token).Error

	if err != nil {
		panic(err)
	}

	return token
}

func CreateWallet(dbTxn db.DatabaseTransaction, wallet *models.Wallet) *models.Wallet {
	db := dbTxn.Get()

	err := db.Create(&wallet).Error

	if err != nil {
		panic(err)
	}

	return wallet
}
