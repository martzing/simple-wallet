package user

type GetTokenParams struct {
	TokenID int `form:"token_id" json:"token_id" binding:"required"`
}

type GetTokenRes struct {
	ID     int     `form:"id" json:"id"`
	Name   string  `form:"name" json:"name"`
	Symbol string  `form:"symbol" json:"symbol"`
	Image  string  `form:"image" json:"image"`
	Value  float64 `form:"value" json:"value"`
}

type GetWalletRes struct {
	ID      int     `form:"id" json:"id"`
	Balance float64 `form:"balance" json:"balance"`
	Token   string  `form:"token" json:"token"`
	Symbol  string  `form:"symbol" json:"symbol"`
	Image   string  `form:"image" json:"image"`
}

type TransferTokenParams struct {
	FromUserId int     `form:"from_user_id" json:"from_user_id" binding:"required,numeric,min=1"`
	ToUserId   int     `form:"to_user_id" json:"to_user_id" binding:"required,numeric,min=1"`
	FromToken  string  `form:"from_token" json:"from_token" binding:"required,alpha,uppercase"`
	ToToken    string  `form:"to_token" json:"to_token" binding:"required,alpha,uppercase"`
	Amount     float64 `form:"amount" json:"amount" binding:"required,number,gt=0"`
}
