package authDB

import (
	"github.com/martzing/simple-wallet/models"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) (*models.User, error) {

	err := db.Create(user).Error

	return user, err
}

func CreateWallets(db *gorm.DB, wallets []*models.Wallet) ([]*models.Wallet, error) {

	err := db.Create(wallets).Error

	if err != nil {
		panic(err)
	}

	return wallets, err
}
