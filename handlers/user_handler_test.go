package handlers

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"go-gin-postgres-crud/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

const (
	usersRoute   = "/users"
	userIDRoute  = "/users/:id"
	usersPath    = "/users"
	userPath     = "/users/1"
	manishEmail  = "manish@example.com"
	manishName   = "Manish"
	alexName     = "Alex"
	alexEmail    = "alex@example.com"
	updatedName  = "Updated"
	updatedEmail = "updated@example.com"
)

func setupTestRouter(db *sql.DB) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	h := NewUserHandler(db)

	r.POST(usersRoute, h.CreateUser)
	r.GET(usersRoute, h.GetUsers)
	r.GET(userIDRoute, h.GetUserByID)
	r.PUT(userIDRoute, h.UpdateUser)
	r.DELETE(userIDRoute, h.DeleteUser)

	return r
}

func TestCreateUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	user := models.User{
		Name:  manishName,
		Email: manishEmail,
	}

	mock.ExpectQuery(`INSERT INTO users`).
		WithArgs(user.Name, user.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, usersPath, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router := setupTestRouter(db)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response models.User
	json.Unmarshal(w.Body.Bytes(), &response)

	assert.Equal(t, 1, response.ID)
	assert.Equal(t, user.Name, response.Name)
	assert.Equal(t, user.Email, response.Email)
}

func TestGetUsers(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow(1, manishName, manishEmail).
		AddRow(2, alexName, alexEmail)

	mock.ExpectQuery(`SELECT id,name,email FROM users`).
		WillReturnRows(rows)

	req := httptest.NewRequest(http.MethodGet, usersPath, nil)
	w := httptest.NewRecorder()

	router := setupTestRouter(db)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var users []models.User
	json.Unmarshal(w.Body.Bytes(), &users)

	assert.Len(t, users, 2)
}
func TestGetUserByID(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	userID := 1
	row := sqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow(1, manishName, manishEmail)

	mock.ExpectQuery(`SELECT id,name,email FROM users`).
		WithArgs(userID).
		WillReturnRows(row)

	req := httptest.NewRequest(http.MethodGet, userPath, nil)
	w := httptest.NewRecorder()

	router := setupTestRouter(db)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var user models.User
	json.Unmarshal(w.Body.Bytes(), &user)

	assert.Equal(t, 1, user.ID)
	assert.Equal(t, manishName, user.Name)
	assert.Equal(t, manishEmail, user.Email)
}

func TestUpdateUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	user := models.User{
		Name:  updatedName,
		Email: updatedEmail,
	}

	mock.ExpectExec(`UPDATE users`).
		WithArgs(user.Name, user.Email, 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPut, userPath, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router := setupTestRouter(db)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}

func TestDeleteUser(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	mock.ExpectExec(`DELETE FROM users WHERE id=\$1`).
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	req := httptest.NewRequest(http.MethodDelete, userPath, nil)
	w := httptest.NewRecorder()

	router := setupTestRouter(db)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
