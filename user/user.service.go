package user

import (
	"net/http"

	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/helpers"
	userDB "github.com/martzing/simple-wallet/user/db"
)

func getTokens() []GetTokenRes {
	dbTxn := db.NewTransaction()

	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()

	dbTxn.Begin()
	tokens := userDB.GetTokens(dbTxn)
	dbTxn.Commit()

	result := []GetTokenRes{}
	for _, token := range *tokens {
		result = append(result, GetTokenRes{
			ID:     token.ID,
			Name:   token.Name,
			Symbol: token.Symbol,
			Image:  token.Image,
			Value:  token.Value,
		})
	}
	return result
}

func getToken(tokenId int) GetTokenRes {
	dbTxn := db.NewTransaction()

	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()

	dbTxn.Begin()
	token := userDB.GetToken(dbTxn, tokenId)
	if token == nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "Token not found",
			StatusCode: http.StatusNotFound,
		}
		panic(ce)
	}

	dbTxn.Commit()

	return GetTokenRes{
		ID:     token.ID,
		Name:   token.Name,
		Symbol: token.Symbol,
		Image:  token.Image,
		Value:  token.Value,
	}
}
