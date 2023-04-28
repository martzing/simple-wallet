package route

//import  "github.com/martzing/simple-wallet/auth/controller/auth"

var healthRoutes = []Route{
	{
		Method:  "GET",
		Path:    "/ready",
		Handler: authController.Register,
	},
}
