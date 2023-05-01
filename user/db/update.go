package userDB

import (
	"github.com/martzing/simple-wallet/models"
	"gorm.io/gorm"
)

func UpdateWallet(db *gorm.DB, wallet *models.Wallet) error {
	err := db.Save(wallet).Error
	return err
}
