package authDB

import (
	"errors"

	"github.com/martzing/simple-wallet/models"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB, username string) (*models.User, error) {

	user := models.User{}

	db = db.Where("username = ?", username)
	err := db.First(&user).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &user, err
}

func GetTokens(db *gorm.DB) ([]*models.Token, error) {

	tokens := []*models.Token{}

	err := db.Find(&tokens).Error

	return tokens, err
}
