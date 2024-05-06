package repositories

import (
	"api/src/models"
	"database/sql"
)

type usersRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *usersRepository {
	return &usersRepository{db}
}

func (repository usersRepository) Create(user models.User) (uint64, error) {

	statement, error := repository.db.Prepare(
		"INSERT INTO users (name, username, email, password) values(?, ?, ?, ?)",
	)

	if error != nil {
		return 0, error
	}

	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Username, user.Email, user.Password)

	if error != nil {
		return 0, error
	}

	ID, error := result.LastInsertId()

	if error != nil {
		return 0, error
	}

	return uint64(ID), nil
}