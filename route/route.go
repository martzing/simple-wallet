package route

import (
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	auth := r.Group("/auth")
	for _, route := range authRoutes {
		route.Middleware = append(route.Middleware, route.Handler)
		auth.Handle(route.Method, route.Path, route.Middleware...)
	}
}
