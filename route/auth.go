package route

import (
	"github.com/martzing/simple-wallet/auth"
)

var authRoutes = []Route{
	{
		Method: "POST",
		Path:   "/register",
		// Middleware: []gin.HandlerFunc{
		// 	middleware.UserMiddleware,
		// },
		Handler: auth.Register,
	},
	{
		Method:  "POST",
		Path:    "/login",
		Handler: auth.Login,
	},
}
