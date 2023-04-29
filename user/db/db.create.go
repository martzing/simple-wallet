package userDB

import (
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/models"
)

func CreateWallet(dbTxn db.DatabaseTransaction, wallet *models.Wallet) *models.Wallet {
	db := dbTxn.Get()

	err := db.Create(&wallet).Error

	if err != nil {
		panic(err)
	}

	return wallet
}
