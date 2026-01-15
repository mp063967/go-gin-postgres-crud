package main

import (
	"go-gin-postgres-crud/config"
	"go-gin-postgres-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	r := gin.Default()
	routes.RegisterRoutes(r, db)

	r.Run(":8082")
}
