package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (repository usersRepository) GetUsers(nameOrUsername string) ([]models.User, error) {
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername)
	
	rows, error := repository.db.Query("SELECT id, name, username, email, created_at FROM users WHERE name LIKE ? or username LIKE ?", nameOrUsername, nameOrUsername)

	if error != nil {
		return nil, error
	}

	var users []models.User

	for rows.Next() {
		var user models.User

		if error := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.CreatedAt); error != nil {
			return nil, error
		}

		users = append(users, user)
	}
	return users, nil
}

func (repository usersRepository) GetUserById(userId uint64) (models.User, error) {
	
	row, error := repository.db.Query("SELECT id, name, username, email, created_at FROM users WHERE id = ?", userId)

	if error != nil {
		return models.User{}, error
	}

	var user models.User

	if row.Next() {
		if error := row.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.CreatedAt); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}

func (repository usersRepository) UpdateUser(userId uint64, user models.User) error {
	
	statement, error := repository.db.Prepare("UPDATE users set name = ?, username = ?, email = ? WHERE id = ?")

	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error := statement.Exec(user.Name, user.Username, user.Email, userId); error != nil {
		return error
	}

	return nil
}


func (repository usersRepository) DeleteUser(userId uint64) error {
	
	statement, error := repository.db.Prepare("DELETE FROM users WHERE id = ?")

	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error := statement.Exec(userId); error != nil {
		return error
	}

	return nil
}

func (repository usersRepository) GetUserByEmail(userEmail string) (models.User, error) {
	fmt.Println(userEmail)
	row, error := repository.db.Query("SELECT id, email, password FROM users WHERE email = ?", userEmail)

	if error != nil {
		return models.User{}, error
	}

	var user models.User

	if row.Next() {
		if error := row.Scan(&user.ID, &user.Email, &user.Password); error != nil {
			return models.User{}, error
		}
	}

	return user, nil
}

func (repository usersRepository) Follow (userId uint64, followerId uint64) error {


	statement, error := repository.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)")

	if error != nil {
		return error
	}

	if _, error := statement.Exec(userId, followerId); error !=nil {
		return error
	}

	defer statement.Close()

	return nil
}

func (repository usersRepository) Unfollow (userId uint64, followerId uint64) error {


	statement, error := repository.db.Prepare("DELETE from followers WHERE user_id = ? AND follower_id = ?")

	if error != nil {
		return error
	}

	if _, error := statement.Exec(userId, followerId); error !=nil {
		return error
	}

	defer statement.Close()

	return nil
}

func (repository usersRepository) GetFollowers(userId uint64) ([]models.User, error) {

	rows, error := repository.db.Query(`
		SELECT u.id, u.name, u.username, u.email, u.created_at
		FROM users u INNER JOIN followers f ON u.id = f.follower_id WHERE f.user_id = ? 
	`, userId)

	if error != nil {
		return nil, error
	}

	var users []models.User

	for rows.Next() {
		var user models.User

		if error := rows.Scan(
			&user.ID, 
			&user.Name,
			&user.Username,
			&user.Email,
			&user.CreatedAt,
		); error !=nil {
			return nil, error
		}
		users = append(users, user)
	}

	return users, nil
}