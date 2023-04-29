package admin

type CreateTokenParams struct {
	Name   string  `form:"name" json:"name" binding:"required"`
	Symbol string  `form:"symbol" json:"symbol" binding:"required,alpha,uppercase"`
	Image  string  `form:"image" json:"image" binding:"required,url"`
	Value  float64 `form:"value" json:"value" binding:"required,number,gt=0"`
}

type CreateTokenRes struct {
	Name   string  `form:"name" json:"name"`
	Symbol string  `form:"symbol" json:"symbol"`
	Image  string  `form:"image" json:"image"`
	Value  float64 `form:"value" json:"value"`
}

type UpdateTokenParams struct {
	ID     int      `form:"id" json:"id" binding:"required,numeric,min=1"`
	Name   *string  `form:"name" json:"name" binding:"required_without_all=Symbol Image Value,omitempty"`
	Symbol *string  `form:"symbol" json:"symbol" binding:"required_without_all=Name Image Value,omitempty,alpha,uppercase"`
	Image  *string  `form:"image" json:"image" binding:"required_without_all=Name Symbol Value,omitempty,url"`
	Value  *float64 `form:"value" json:"value" binding:"required_without_all=Name Symbol Image,omitempty,number,gt=0"`
}

type UpdateTokenRes struct {
	Message string `form:"message" json:"message"`
}

type DeleteTokenRes struct {
	Message string `form:"message" json:"message"`
}

type UpdateTokenBalanceParams struct {
	UserID      int     `form:"user_id" json:"user_id" binding:"required,numeric,min=1"`
	TokenSymbol string  `form:"token" json:"token" binding:"required,alpha,uppercase"`
	Amount      float64 `form:"amount" json:"amount" binding:"required,number,gt=0"`
}

type UpdateBalanceRes struct {
	Message string `form:"message" json:"message"`
}
