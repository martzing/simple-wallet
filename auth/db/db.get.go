package authDB

import (
	"errors"

	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/models"
	"gorm.io/gorm"
)

func GetUser(dbTxn db.DatabaseTransaction, username string) *models.User {
	db := dbTxn.Get()

	user := models.User{}

	db = db.Where("username = ?", username)
	err := db.First(&user).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			panic(err)
		}
	}
	return &user
}

func GetTokens(dbTxn db.DatabaseTransaction) []*models.Token {
	db := dbTxn.Get()

	tokens := []*models.Token{}

	err := db.Find(&tokens).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			panic(err)
		}
	}
	return tokens
}
