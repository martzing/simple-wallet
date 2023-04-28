package main

import (
	"github.com/martzing/simple-wallet/configs"
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/route"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.BootConfig()
	db.Connect(*configs.DbConfig)
	r := gin.Default()
	route.Init(r)
	r.Run("localhost:9000")
}
