package auth

type RegisterData struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
}

type RegisterRes struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
