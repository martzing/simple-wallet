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

func getTokens() (*[]GetTokenRes, error) {

	tokens, err := userDB.GetTokens(db.DB)

	if err != nil {
		return nil, err
	}

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
	return &result, nil
}

func getToken(tokenId int) (*GetTokenRes, helpers.CustomError) {

	token, err := userDB.GetToken(db.DB, tokenId)

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

	if err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	return &GetTokenRes{
		ID:     token.ID,
		Name:   token.Name,
		Symbol: token.Symbol,
		Image:  token.Image,
		Value:  token.Value,
	}, nil
}

func getWallet(userId int) (*[]GetWalletRes, error) {

	wallets, err := userDB.GetWallets(db.DB, userId)

	if err != nil {
		return nil, err
	}

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
	return &result, nil
}

func transferToken(data *TransferTokenParams) (*TransferTokenRes, helpers.CustomError) {

	if data.FromUserId == data.ToUserId && data.FromToken == data.ToToken {
		msg := "Cannot transfer token to same wallet"
		code := http.StatusBadRequest
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	user, err := userDB.GetUser(db.DB, data.ToUserId)
	if err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}
	if user == nil {
		msg := "Receiver user not found"
		code := http.StatusNotFound
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	var fromToken *models.Token
	var toToken *models.Token

	tokens, err := userDB.GetTokenBySymbols(db.DB, []string{data.FromToken, data.ToToken})

	if err != nil {
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	for _, token := range tokens {
		if token.Symbol == data.FromToken {
			fromToken = token
		}
		if token.Symbol == data.ToToken {
			toToken = token
		}
	}

	if fromToken == nil {
		msg := "Invalid token"
		code := http.StatusBadRequest
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	if toToken == nil {
		msg := "Invalid destination token"
		code := http.StatusBadRequest
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

	wallet, err := userDB.GetWalletBy(DB, data.FromUserId, fromToken.ID, true)

	if err != nil {
		dbTxn.Rollback()
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	if wallet == nil {
		dbTxn.Rollback()
		msg := "You don't have this token"
		code := http.StatusBadRequest
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	balance := decimal.NewFromFloat(wallet.Balance)
	if balance.LessThanOrEqual(decimal.NewFromInt(0)) {
		dbTxn.Rollback()
		msg := "Token balance not enough"
		code := http.StatusBadRequest
		var ce helpers.CustomError
		ce = &helpers.Error{
			Message:    &msg,
			StatusCode: &code,
		}
		return nil, ce
	}

	transferAmount := decimal.NewFromFloat(data.Amount)
	if fromToken.ID != toToken.ID {
		transferAmount = decimal.NewFromFloat(fromToken.Value).Div(decimal.NewFromFloat(toToken.Value)).Mul(decimal.NewFromFloat(data.Amount))
	}

	toWallet, err := userDB.GetWalletBy(DB, user.ID, toToken.ID, true)

	if err != nil {
		dbTxn.Rollback()
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	if toWallet == nil {
		newToWalletBalance, _ := transferAmount.Float64()
		_, err := userDB.CreateWallet(DB, &models.Wallet{
			Balance: newToWalletBalance,
			UserID:  user.ID,
			TokenID: toToken.ID,
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
		newToWalletBalance, _ := transferAmount.Add(decimal.NewFromFloat(toWallet.Balance)).Float64()
		toWallet.Balance = newToWalletBalance
		err := userDB.UpdateWallet(DB, toWallet)
		if err != nil {
			dbTxn.Rollback()
			var ce helpers.CustomError
			ce = &helpers.Error{
				Err: err,
			}
			return nil, ce
		}
	}

	newWalletBalance, _ := decimal.NewFromFloat(wallet.Balance).Sub(decimal.NewFromFloat(data.Amount)).Float64()
	wallet.Balance = newWalletBalance

	err = userDB.UpdateWallet(DB, wallet)
	if err != nil {
		dbTxn.Rollback()
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	toTokenAmount, _ := transferAmount.Float64()
	tid, _ := hex.DecodeString(strings.ReplaceAll(uuid.New().String(), "-", ""))
	_, err = userDB.CreateTransferTransaction(DB, &models.TransferTransaction{
		ID:              tid,
		FromUserID:      data.FromUserId,
		ToUserID:        user.ID,
		FromTokenID:     fromToken.ID,
		ToTokenID:       toToken.ID,
		FromTokenAmount: data.Amount,
		ToTokenAmount:   toTokenAmount,
	})

	if err != nil {
		dbTxn.Rollback()
		var ce helpers.CustomError
		ce = &helpers.Error{
			Err: err,
		}
		return nil, ce
	}

	dbTxn.Commit()

	return &TransferTokenRes{
		Message: "Transfer token success",
	}, nil
}

func getTransferTokens(userId int) ([]GetTransferTokenRes, error) {

	txns, err := userDB.GetTransferTokens(db.DB, userId)

	if err != nil {
		return nil, err
	}

	result := []GetTransferTokenRes{}
	for _, txn := range txns {
		result = append(result, GetTransferTokenRes{
			ID:              hex.EncodeToString(txn.ID),
			ToUser:          txn.ToUser.Username,
			FromToken:       txn.FromToken.Name,
			ToToken:         txn.ToToken.Name,
			FromTokenAmount: txn.FromTokenAmount,
			ToTokenAmount:   txn.ToTokenAmount,
			TransactionDate: txn.CreatedAt,
		})
	}
	return result, nil
}
