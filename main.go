package main

import (
	"github.com/martzing/simple-wallet/configs"
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	configs.BootConfig()
	db.Connect(*configs.DbConfig)
	route.Init(r)
	r.Run("localhost:9000")
}
