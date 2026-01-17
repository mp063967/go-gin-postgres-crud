package repository

import (
	"database/sql"

	"go-gin-postgres-crud/models"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.QueryRow(
		"INSERT INTO users(name,email) VALUES($1,$2) RETURNING id",
		user.Name, user.Email,
	).Scan(&user.ID)
}

func (r *userRepository) FindAll() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id,name,email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepository) FindByID(id int) (*models.User, error) {
	var u models.User
	err := r.db.QueryRow(
		"SELECT id,name,email FROM users WHERE id=$1", id,
	).Scan(&u.ID, &u.Name, &u.Email)

	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *userRepository) Update(id int, user *models.User) error {
	_, err := r.db.Exec(
		"UPDATE users SET name=$1,email=$2 WHERE id=$3",
		user.Name, user.Email, id,
	)
	return err
}

func (r *userRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id=$1", id)
	return err
}
