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

func CreateTransferTransaction(dbTxn db.DatabaseTransaction, transaction *models.TransferTransaction) *models.TransferTransaction {
	db := dbTxn.Get()

	err := db.Create(&transaction).Error

	if err != nil {
		panic(err)
	}

	return transaction
}
