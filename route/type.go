package route

import "github.com/gin-gonic/gin"

type Route struct {
	Method     string
	Path       string
	Middleware []gin.HandlerFunc
	Handler    gin.HandlerFunc
}
