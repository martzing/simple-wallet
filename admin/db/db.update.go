package adminDB

import (
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/models"
)

func UpdateToken(dbTxn db.DatabaseTransaction, token *models.Token) {
	db := dbTxn.Get()
	err := db.Save(token).Error
	if err != nil {
		panic(err)
	}
}

func UpdateWallet(dbTxn db.DatabaseTransaction, wallet *models.Wallet) {
	db := dbTxn.Get()
	err := db.Save(wallet).Error
	if err != nil {
		panic(err)
	}
}
