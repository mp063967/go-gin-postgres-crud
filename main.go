package main

import (
	"go-gin-postgres-crud/config"
	"go-gin-postgres-crud/handlers"
	"go-gin-postgres-crud/repository"
	"go-gin-postgres-crud/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepo)

	r := gin.Default()
	routes.RegisterRoutes(r, userHandler)

	r.Run(":8082")
}
