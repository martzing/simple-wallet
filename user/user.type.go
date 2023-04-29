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
