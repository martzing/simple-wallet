package route

import (
	"github.com/gin-gonic/gin"
	"github.com/martzing/simple-wallet/middleware"
	"github.com/martzing/simple-wallet/user"
)

var userRoutes = []Route{
	{
		Method: "GET",
		Path:   "/token",
		Middleware: []gin.HandlerFunc{
			middleware.UserMiddleware,
		},
		Handler: user.GetTokens,
	},
	{
		Method: "GET",
		Path:   "/wallet",
		Middleware: []gin.HandlerFunc{
			middleware.UserMiddleware,
		},
		Handler: user.GetWallet,
	},
	{
		Method: "GET",
		Path:   "/token/:token_id",
		Middleware: []gin.HandlerFunc{
			middleware.UserMiddleware,
		},
		Handler: user.GetToken,
	},
}
