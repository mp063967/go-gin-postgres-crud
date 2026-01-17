package handlers

import (
	"errors"
	"go-gin-postgres-crud/models"
)

type MockUserRepository struct {
	users map[int]models.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[int]models.User),
	}
}

func (m *MockUserRepository) Create(user *models.User) error {
	user.ID = len(m.users) + 1
	m.users[user.ID] = *user
	return nil
}

func (m *MockUserRepository) FindAll() ([]models.User, error) {
	var list []models.User
	for _, u := range m.users {
		list = append(list, u)
	}
	return list, nil
}

func (m *MockUserRepository) FindByID(id int) (*models.User, error) {
	u, ok := m.users[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &u, nil
}

func (m *MockUserRepository) Update(id int, user *models.User) error {
	if _, ok := m.users[id]; !ok {
		return errors.New("not found")
	}
	user.ID = id
	m.users[id] = *user
	return nil
}

func (m *MockUserRepository) Delete(id int) error {
	delete(m.users, id)
	return nil
}
