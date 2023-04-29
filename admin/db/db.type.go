package adminDB

import "github.com/martzing/simple-wallet/models"

type SumWallet struct {
	TokenID int
	Total   float64
	Token   models.Token
}
