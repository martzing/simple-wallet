package admin

import (
	"fmt"
	"net/http"

	adminDB "github.com/martzing/simple-wallet/admin/db"
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/helpers"
	"github.com/martzing/simple-wallet/models"
	"github.com/shopspring/decimal"
)

func createToken(data *CreateTokenParams) CreateTokenRes {
	dbTxn := db.NewTransaction()

	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()

	dbTxn.Begin(db.REPEATABLE_READ)

	token := adminDB.CreateToken(dbTxn, &models.Token{
		Name:   data.Name,
		Symbol: data.Symbol,
		Image:  data.Image,
		Value:  data.Value,
	})
	dbTxn.Commit()
	return CreateTokenRes{
		Name:   token.Name,
		Symbol: token.Symbol,
		Image:  token.Image,
		Value:  token.Value,
	}
}

func updateToken(data *UpdateTokenParams) UpdateTokenRes {
	fmt.Printf("updateToken\n")
	dbTxn := db.NewTransaction()

	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()

	dbTxn.Begin(db.REPEATABLE_READ)

	token := adminDB.GetToken(dbTxn, data.ID)

	if token == nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "Token not found",
			StatusCode: http.StatusNotFound,
		}
		panic(ce)
	}

	if data.Name != nil {
		token.Name = *data.Name
	}
	if data.Symbol != nil {
		token.Symbol = *data.Symbol
	}
	if data.Image != nil {
		token.Image = *data.Image
	}
	if data.Value != nil {
		token.Value = *data.Value
	}

	adminDB.UpdateToken(dbTxn, token)

	dbTxn.Commit()

	return UpdateTokenRes{
		Message: "Update token success",
	}
}

func deleteToken(tokenId int) DeleteTokenRes {
	dbTxn := db.NewTransaction()

	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()

	dbTxn.Begin(db.REPEATABLE_READ)

	token := adminDB.GetToken(dbTxn, tokenId)

	if token == nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "Token not found",
			StatusCode: http.StatusNotFound,
		}
		panic(ce)
	}

	adminDB.DeleteToken(dbTxn, token)

	dbTxn.Commit()
	return DeleteTokenRes{
		Message: "Delete token success",
	}
}

func updateTokenBalance(data *UpdateTokenBalanceParams, action string) UpdateBalanceRes {
	dbTxn := db.NewTransaction()

	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()

	dbTxn.Begin(db.REPEATABLE_READ)

	token := adminDB.GetTokenBySymbol(dbTxn, data.TokenSymbol)

	if token == nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "Token not found",
			StatusCode: http.StatusNotFound,
		}
		panic(ce)
	}

	wallet := adminDB.GetWalletBy(dbTxn, data.UserID, token.ID)

	if wallet == nil {
		if action == "minus" {
			var ce helpers.CustomError
			ce = &helpers.Error{
				Message:    "Balance cannot be negative",
				StatusCode: http.StatusBadRequest,
			}
			panic(ce)
		}
		adminDB.CreateWallet(dbTxn, &models.Wallet{
			UserID:  data.UserID,
			TokenID: token.ID,
			Balance: data.Amount,
		})
	} else {
		balance := decimal.NewFromFloat(wallet.Balance)
		if action == "add" {
			balance = balance.Add(decimal.NewFromFloat(data.Amount))
		} else if action == "minus" {
			balance = balance.Sub(decimal.NewFromFloat(data.Amount))
		} else {
			var ce helpers.CustomError
			ce = &helpers.Error{
				Message:    "Something went wrong, Please contact support",
				StatusCode: http.StatusInternalServerError,
			}
			panic(ce)
		}
		newBalance, _ := balance.Float64()
		wallet.Balance = newBalance
		adminDB.UpdateWallet(dbTxn, wallet)
	}

	dbTxn.Commit()
	return UpdateBalanceRes{
		Message: "Update token balance success",
	}
}

func getTokenBalance() []GetTokenBalanceRes {
	dbTxn := db.NewTransaction()

	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()

	dbTxn.Begin(db.REPEATABLE_READ)

	tokens := adminDB.GetTokens(dbTxn)
	wallets := adminDB.SumWalletBalance(dbTxn)

	mapBalance := make(map[string]*GetTokenBalanceRes)
	for _, token := range tokens {
		mapBalance[token.Symbol] = &GetTokenBalanceRes{
			Name:    token.Name,
			Symbol:  token.Symbol,
			Image:   token.Image,
			Balance: 0,
		}
	}

	for _, wallet := range wallets {
		mapBalance[wallet.Token.Symbol].Balance = wallet.Total
	}

	result := []GetTokenBalanceRes{}
	for _, balance := range mapBalance {
		result = append(result, *balance)
	}

	dbTxn.Commit()

	return result
}
