package main

import (
	"github.com/martzing/simple-wallet/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	route.Init(r)
	r.Run()
}
