package main

import (
	"github.com/martzing/simple-wallet/db"
	"github.com/martzing/simple-wallet/route"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	config := db.DBConfig{
		Host:     "localhost",
		Port:     "3306",
		Username: "root",
		Password: "root",
		DBName:   "db",
	}
	db.Connect(config)
	route.Init(r)
	r.Run("localhost:9000")
}
