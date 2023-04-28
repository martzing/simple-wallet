package route

import (
	"github.com/gin-gonic/gin"
	"github.com/martzing/simple-wallet/auth"
	"github.com/martzing/simple-wallet/middleware"
)

var authRoutes = []Route{
	{
		Method: "POST",
		Path:   "/register",
		Middleware: []gin.HandlerFunc{
			middleware.UserMiddleware,
		},
		Handler: auth.Register,
	},
	{
		Method: "POST",
		Path:   "/login",
		Middleware: []gin.HandlerFunc{
			middleware.UserMiddleware,
		},
		Handler: auth.Login,
	},
}
