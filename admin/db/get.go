package adminDB

import (
	"errors"

	"github.com/martzing/simple-wallet/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetToken(db *gorm.DB, tokenId int) (*models.Token, error) {

	token := models.Token{}

	db = db.Where("id = ?", tokenId)
	err := db.First(&token).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &token, err
}

func GetTokens(db *gorm.DB) ([]*models.Token, error) {

	tokens := []*models.Token{}

	err := db.Find(&tokens).Error

	return tokens, err
}

func GetTokenBySymbol(db *gorm.DB, symbol string) (*models.Token, error) {

	token := models.Token{}

	db = db.Where("symbol = ?", symbol)
	err := db.First(&token).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &token, err
}

func GetWalletBy(db *gorm.DB, userId int, tokenId int, isLock bool) (*models.Wallet, error) {

	wallet := models.Wallet{}

	if isLock {
		db = db.Clauses(clause.Locking{Strength: "UPDATE"})
	}

	db = db.Where("user_id = ? AND token_id = ?", userId, tokenId)
	err := db.First(&wallet).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &wallet, err
}

func SumWalletBalance(db *gorm.DB) ([]*SumWallet, error) {

	wallets := []*SumWallet{}

	err := db.Model(&models.Wallet{}).Select("token_id, sum(balance) as total").InnerJoins("Token").Group("token_id").Find(&wallets).Error

	return wallets, err
}
