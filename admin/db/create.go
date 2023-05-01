package adminDB

import (
	"github.com/martzing/simple-wallet/models"
	"gorm.io/gorm"
)

func CreateToken(db *gorm.DB, token *models.Token) (*models.Token, error) {

	err := db.Create(token).Error

	return token, err
}

func CreateWallet(db *gorm.DB, wallet *models.Wallet) (*models.Wallet, error) {

	err := db.Create(&wallet).Error

	return wallet, err
}
