package adminDB

import (
	"errors"

	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/models"
	"gorm.io/gorm"
)

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

func GetTokenBySymbol(dbTxn db.DatabaseTransaction, symbol string) *models.Token {
	db := dbTxn.Get()

	token := models.Token{}

	db = db.Where("symbol = ?", symbol)
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

func GetWalletBy(dbTxn db.DatabaseTransaction, userId int, tokenId int) *models.Wallet {
	db := dbTxn.Get()

	wallet := models.Wallet{}

	db = db.Where("user_id = ? AND token_id = ?", userId, tokenId)
	err := db.First(&wallet).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			panic(err)
		}
	}
	return &wallet
}
