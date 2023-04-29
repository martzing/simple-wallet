package route

import (
	"github.com/gin-gonic/gin"
	"github.com/martzing/simple-wallet/admin"
	"github.com/martzing/simple-wallet/middleware"
)

var adminRoutes = []Route{
	{
		Method: "POST",
		Path:   "/token",
		Middleware: []gin.HandlerFunc{
			middleware.AdminMiddleware,
		},
		Handler: admin.CreateToken,
	},
	{
		Method: "PATCH",
		Path:   "/token",
		Middleware: []gin.HandlerFunc{
			middleware.AdminMiddleware,
		},
		Handler: admin.UpdateToken,
	},
	{
		Method: "DELETE",
		Path:   "/token/:token_id",
		Middleware: []gin.HandlerFunc{
			middleware.AdminMiddleware,
		},
		Handler: admin.DeleteToken,
	},
	{
		Method: "POST",
		Path:   "/token/increase/balance",
		Middleware: []gin.HandlerFunc{
			middleware.AdminMiddleware,
		},
		Handler: admin.AddTokenBalance,
	},
	{
		Method: "POST",
		Path:   "/token/decrease/balance",
		Middleware: []gin.HandlerFunc{
			middleware.AdminMiddleware,
		},
		Handler: admin.MinusTokenBalance,
	},
	{
		Method: "GET",
		Path:   "/token/balance",
		Middleware: []gin.HandlerFunc{
			middleware.AdminMiddleware,
		},
		Handler: admin.GetTokenBalance,
	},
}
