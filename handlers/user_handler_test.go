package handlers

import (
	"bytes"
	"encoding/json"
	"go-gin-postgres-crud/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	repo := NewMockUserRepository()
	router := setupTestRouter(repo)

	user := models.User{
		Name:  manishName,
		Email: manishEmail,
	}

	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, usersPath, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.User
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, 1, response.ID)
	assert.Equal(t, manishName, response.Name)
	assert.Equal(t, manishEmail, response.Email)
}

func TestGetUsers(t *testing.T) {
	repo := NewMockUserRepository()
	seedUser(repo, manishName, manishEmail)
	seedUser(repo, alexName, alexEmail)

	router := setupTestRouter(repo)

	req := httptest.NewRequest(http.MethodGet, usersPath, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users []models.User
	json.Unmarshal(w.Body.Bytes(), &users)

	assert.Len(t, users, 2)
}
