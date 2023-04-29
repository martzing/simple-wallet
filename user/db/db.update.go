package userDB

import (
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/models"
)

func UpdateWallet(dbTxn db.DatabaseTransaction, wallet *models.Wallet) {
	db := dbTxn.Get()
	err := db.Save(wallet).Error
	if err != nil {
		panic(err)
	}
}
