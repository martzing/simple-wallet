package adminDB

import (
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/models"
)

func DeleteToken(dbTxn db.DatabaseTransaction, token *models.Token) {
	db := dbTxn.Get()
	err := db.Delete(token).Error
	if err != nil {
		panic(err)
	}
}
