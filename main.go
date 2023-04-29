package main

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/martzing/simple-wallet/configs"
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/route"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.BootConfig()
	db.Connect(*configs.DbConfig)

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(helmet.Default())
	route.Init(r)
	r.Run("localhost:9000")
}
