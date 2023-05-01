package userDB

import (
	"errors"

	"github.com/martzing/simple-wallet/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetTokens(db *gorm.DB) ([]*models.Token, error) {

	tokens := []*models.Token{}

	err := db.Find(&tokens).Error

	return tokens, err
}

func GetToken(db *gorm.DB, tokenId int) (*models.Token, error) {

	token := models.Token{}

	db = db.Where("id = ?", tokenId)
	err := db.First(&token).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &token, err
}

func GetWallets(db *gorm.DB, userId int) ([]*models.Wallet, error) {

	wallets := []*models.Wallet{}

	db = db.Where("user_id = ?", userId)
	err := db.Preload("Token").Find(&wallets).Error

	return wallets, err
}

func GetUser(db *gorm.DB, userId int) (*models.User, error) {

	user := models.User{}

	db = db.Where("id = ?", userId)
	err := db.First(&user).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func GetTokenBySymbols(db *gorm.DB, symbols []string) ([]*models.Token, error) {

	tokens := []*models.Token{}

	db = db.Where("symbol IN ?", symbols)
	err := db.Find(&tokens).Error

	return tokens, err
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

func GetTransferTokens(db *gorm.DB, userId int) ([]*models.TransferTransaction, error) {

	txns := []*models.TransferTransaction{}

	db = db.Where("from_user_id = ?", userId)
	err := db.Preload("ToUser").Preload("FromToken").Preload("ToToken").Find(&txns).Error

	return txns, err
}
