package handlers

import (
	"go-gin-postgres-crud/models"
	"go-gin-postgres-crud/repository"

	"github.com/gin-gonic/gin"
)

const (
	usersRoute  = "/users"
	userIDRoute = "/users/:id"

	usersPath = "/users"
	userPath  = "/users/1"

	manishName  = "Manish"
	manishEmail = "manish@example.com"

	alexName  = "Alex"
	alexEmail = "alex@example.com"

	updatedName  = "Updated"
	updatedEmail = "updated@example.com"
)

func setupTestRouter(repo repository.UserRepository) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	h := NewUserHandler(repo)

	r.POST(usersRoute, h.CreateUser)
	r.GET(usersRoute, h.GetUsers)
	r.GET(userIDRoute, h.GetUserByID)
	r.PUT(userIDRoute, h.UpdateUser)
	r.DELETE(userIDRoute, h.DeleteUser)

	return r
}

// Optional helper to seed data
func seedUser(repo repository.UserRepository, name, email string) models.User {
	u := models.User{
		Name:  name,
		Email: email,
	}
	repo.Create(&u)
	return u
}
