package userDB

import (
	"github.com/martzing/simple-wallet/models"
	"gorm.io/gorm"
)

func CreateWallet(db *gorm.DB, wallet *models.Wallet) (*models.Wallet, error) {

	err := db.Create(wallet).Error

	return wallet, err
}

func CreateTransferTransaction(db *gorm.DB, transaction *models.TransferTransaction) (*models.TransferTransaction, error) {

	err := db.Create(transaction).Error

	return transaction, err
}
