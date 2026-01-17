package repository

import "go-gin-postgres-crud/models"

type UserRepository interface {
	Create(user *models.User) error
	FindAll() ([]models.User, error)
	FindByID(id int) (*models.User, error)
	Update(id int, user *models.User) error
	Delete(id int) error
}
