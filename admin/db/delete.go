package adminDB

import (
	"github.com/martzing/simple-wallet/models"
	"gorm.io/gorm"
)

func DeleteToken(db *gorm.DB, token *models.Token) error {
	err := db.Delete(token).Error
	return err
}
