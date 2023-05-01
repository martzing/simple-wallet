package adminDB

import (
	"github.com/martzing/simple-wallet/models"
	"gorm.io/gorm"
)

func UpdateToken(db *gorm.DB, token *models.Token) (*models.Token, error) {
	err := db.Save(token).Error

	return token, err
}

func UpdateWallet(db *gorm.DB, wallet *models.Wallet) (*models.Wallet, error) {
	err := db.Save(wallet).Error

	return wallet, err
}
