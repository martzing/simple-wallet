package userDB

import (
	"errors"

	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/models"
	"gorm.io/gorm"
)

func GetTokens(dbTxn db.DatabaseTransaction) *[]models.Token {
	db := dbTxn.Get()

	tokens := []models.Token{}

	err := db.Find(&tokens).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			panic(err)
		}
	}
	return &tokens
}

func GetToken(dbTxn db.DatabaseTransaction, tokenId int) *models.Token {
	db := dbTxn.Get()

	token := models.Token{}

	db = db.Where("id = ?", tokenId)
	err := db.First(&token).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			panic(err)
		}
	}
	return &token
}

func GetWallets(dbTxn db.DatabaseTransaction, userId int) *[]models.Wallet {
	db := dbTxn.Get()

	wallets := []models.Wallet{}

	db = db.Where("user_id = ?", userId)
	err := db.Preload("Token").Find(&wallets).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			panic(err)
		}
	}
	return &wallets
}
