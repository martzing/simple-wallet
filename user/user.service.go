package user

import (
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/helpers"
	"github.com/martzing/simple-wallet/models"
	userDB "github.com/martzing/simple-wallet/user/db"
	"github.com/shopspring/decimal"
)

func getTokens() []GetTokenRes {
	dbTxn := db.NewTransaction()

	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()

	dbTxn.Begin(db.REPEATABLE_READ)
	tokens := userDB.GetTokens(dbTxn)
	dbTxn.Commit()

	result := []GetTokenRes{}
	for _, token := range tokens {
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

	dbTxn.Begin(db.REPEATABLE_READ)
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

func getWallet(userId int) []GetWalletRes {
	dbTxn := db.NewTransaction()

	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()

	dbTxn.Begin(db.REPEATABLE_READ)
	wallets := userDB.GetWallets(dbTxn, userId)
	dbTxn.Commit()

	result := []GetWalletRes{}
	for _, wallet := range wallets {
		result = append(result, GetWalletRes{
			ID:      wallet.ID,
			Balance: wallet.Balance,
			Token:   wallet.Token.Name,
			Symbol:  wallet.Token.Symbol,
			Image:   wallet.Token.Image,
		})
	}
	return result
}

func transferToken(data *TransferTokenParams) TransferTokenRes {
	dbTxn := db.NewTransaction()

	defer func() {
		if err := recover(); err != nil {
			dbTxn.Rollback()
			panic(err)
		}
	}()
	if data.FromUserId == data.ToUserId && data.FromToken == data.ToToken {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "Cannot transfer token to same wallet",
			StatusCode: http.StatusBadRequest,
		}
		panic(ce)
	}

	dbTxn.Begin(db.REPEATABLE_READ)
	user := userDB.GetUser(dbTxn, data.ToUserId)
	if user == nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "Receiver user not found",
			StatusCode: http.StatusNotFound,
		}
		panic(ce)
	}
	var fromToken *models.Token
	var toToken *models.Token
	tokens := userDB.GetTokenBySymbols(dbTxn, []string{data.FromToken, data.ToToken})

	for _, token := range tokens {
		if token.Symbol == data.FromToken {
			fromToken = token
		}
		if token.Symbol == data.ToToken {
			toToken = token
		}
	}
	if fromToken == nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "Invalid token",
			StatusCode: http.StatusBadRequest,
		}
		panic(ce)
	}
	if toToken == nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "Invalid destination token",
			StatusCode: http.StatusBadRequest,
		}
		panic(ce)
	}

	wallet := userDB.GetWalletBy(dbTxn, data.FromUserId, fromToken.ID)

	if wallet == nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "You don't have this token",
			StatusCode: http.StatusBadRequest,
		}
		panic(ce)
	}

	balance := decimal.NewFromFloat(wallet.Balance)
	if balance.LessThanOrEqual(decimal.NewFromInt(0)) {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    "Token balance not enough",
			StatusCode: http.StatusBadRequest,
		}
		panic(ce)
	}
	transferAmount := decimal.NewFromFloat(data.Amount)
	if fromToken.ID != toToken.ID {
		transferAmount = decimal.NewFromFloat(fromToken.Value).Div(decimal.NewFromFloat(toToken.Value)).Mul(decimal.NewFromFloat(data.Amount))
	}

	toWallet := userDB.GetWalletBy(dbTxn, user.ID, toToken.ID)

	if toWallet == nil {
		newToWalletBalance, _ := transferAmount.Float64()
		userDB.CreateWallet(dbTxn, &models.Wallet{
			Balance: newToWalletBalance,
			UserID:  user.ID,
			TokenID: toToken.ID,
		})
	} else {
		newToWalletBalance, _ := transferAmount.Add(decimal.NewFromFloat(toWallet.Balance)).Float64()
		toWallet.Balance = newToWalletBalance
		userDB.UpdateWallet(dbTxn, toWallet)
	}
	newWalletBalance, _ := decimal.NewFromFloat(wallet.Balance).Sub(decimal.NewFromFloat(data.Amount)).Float64()
	wallet.Balance = newWalletBalance
	userDB.UpdateWallet(dbTxn, wallet)

	toTokenAmount, _ := transferAmount.Float64()
	tid, _ := hex.DecodeString(strings.ReplaceAll(uuid.New().String(), "-", ""))
	userDB.CreateTransferTransaction(dbTxn, &models.TransferTransaction{
		ID:              tid,
		FromUserID:      data.FromUserId,
		ToUserID:        user.ID,
		FromTokenID:     fromToken.ID,
		ToTokenID:       toToken.ID,
		FromTokenAmount: data.Amount,
		ToTokenAmount:   toTokenAmount,
	})
	dbTxn.Commit()
	return TransferTokenRes{
		Message: "Transfer token success",
	}
}
