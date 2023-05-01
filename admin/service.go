package admin

import (
	"net/http"

	adminDB "github.com/martzing/simple-wallet/admin/db"
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/helpers"
	"github.com/martzing/simple-wallet/models"
	"github.com/shopspring/decimal"
)

func createToken(data *CreateTokenParams) (*CreateTokenRes, error) {
	dbTxn := db.NewTransaction()
	DB, err := dbTxn.Begin(db.REPEATABLE_READ)

	if err != nil {
		return nil, err
	}

	token, err := adminDB.CreateToken(DB, &models.Token{
		Name:   data.Name,
		Symbol: data.Symbol,
		Image:  data.Image,
		Value:  data.Value,
	})

	if err != nil {
		dbTxn.Rollback()
		return nil, err
	}

	dbTxn.Commit()

	return &CreateTokenRes{
		Name:   token.Name,
		Symbol: token.Symbol,
		Image:  token.Image,
		Value:  token.Value,
	}, nil
}

func updateToken(data *UpdateTokenParams) (*UpdateTokenRes, helpers.CustomError) {

	token, err := adminDB.GetToken(db.DB, data.ID)

	if err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	if token == nil {
		msg := "Token not found"
		code := http.StatusNotFound
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
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

	dbTxn := db.NewTransaction()
	DB, err := dbTxn.Begin(db.REPEATABLE_READ)

	if err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	_, _err := adminDB.UpdateToken(DB, token)

	if _err != nil {
		dbTxn.Rollback()
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: _err,
		}
		return nil, ce
	}

	dbTxn.Commit()

	return &UpdateTokenRes{
		Message: "Update token success",
	}, nil
}

func deleteToken(tokenId int) (*DeleteTokenRes, helpers.CustomError) {

	token, err := adminDB.GetToken(db.DB, tokenId)

	if err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	if token == nil {
		msg := "Token not found"
		code := http.StatusNotFound
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	dbTxn := db.NewTransaction()
	DB, err := dbTxn.Begin(db.REPEATABLE_READ)

	if err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	_err := adminDB.DeleteToken(DB, token)

	if _err != nil {
		dbTxn.Rollback()
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: _err,
		}
		return nil, ce
	}

	dbTxn.Commit()

	return &DeleteTokenRes{
		Message: "Delete token success",
	}, nil
}

func updateTokenBalance(data *UpdateTokenBalanceParams, action string) (*UpdateBalanceRes, helpers.CustomError) {

	token, err := adminDB.GetTokenBySymbol(db.DB, data.TokenSymbol)

	if err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	if token == nil {
		msg := "Token not found"
		code := http.StatusNotFound
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	dbTxn := db.NewTransaction()
	DB, err := dbTxn.Begin(db.REPEATABLE_READ)

	wallet, err := adminDB.GetWalletBy(DB, data.UserID, token.ID, true)

	if err != nil {
		dbTxn.Rollback()
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	if wallet == nil {
		if action == "minus" {
			dbTxn.Rollback()
			msg := "Balance cannot be negative"
			code := http.StatusBadRequest
			var ce helpers.CustomError
			ce = &helpers.Error{
				Message:    &msg,
				StatusCode: &code,
			}
			return nil, ce
		}
		_, err := adminDB.CreateWallet(DB, &models.Wallet{
			UserID:  data.UserID,
			TokenID: token.ID,
			Balance: data.Amount,
		})
		if err != nil {
			dbTxn.Rollback()
			var ce helpers.CustomError
			ce = &helpers.Error{
				Err: err,
			}
			return nil, ce
		}
	} else {
		balance := decimal.NewFromFloat(wallet.Balance)
		if action == "add" {
			balance = balance.Add(decimal.NewFromFloat(data.Amount))
		} else if action == "minus" {
			balance = balance.Sub(decimal.NewFromFloat(data.Amount))
		} else {
			dbTxn.Rollback()
			msg := "Something went wrong, Please contact support"
			code := http.StatusInternalServerError
			var ce helpers.CustomError
			ce = &helpers.Error{
				Message:    &msg,
				StatusCode: &code,
			}
			return nil, ce
		}

		newBalance, _ := balance.Float64()
		wallet.Balance = newBalance
		_, err := adminDB.UpdateWallet(DB, wallet)

		if err != nil {
			dbTxn.Rollback()
			var ce helpers.CustomError
			ce = &helpers.Error{
				Err: err,
			}
			return nil, ce
		}
	}

	dbTxn.Commit()

	return &UpdateBalanceRes{
		Message: "Update token balance success",
	}, nil
}

func getTokenBalance() (*[]GetTokenBalanceRes, error) {

	tokens, err := adminDB.GetTokens(db.DB)

	if err != nil {
		return nil, err
	}

	wallets, _err := adminDB.SumWalletBalance(db.DB)

	if _err != nil {
		return nil, _err
	}

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

	return &result, nil
}
