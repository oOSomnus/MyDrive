package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"strconv"
)

type UserRepository interface {
	InsertUserIfNotExists(email, hashedPassword string) (userId string, error error)
	FetchPasswordAndUserId(email string) (password string, userId string, error error)
	FetchUserIdWithEmail(email string) (userId string, err error)
	UpdateUserHomeFolder(userId string, homeFolder string) (error error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) InsertUserIfNotExists(email, hashedPassword string) (string, error) {
	result, err := u.db.Exec(
		"INSERT INTO users(email, password, username) VALUES (?, ?, ?)",
		email, hashedPassword, email,
	)
	if err != nil {
		if isDuplicateKeyError(err) {
			return "", fmt.Errorf("user already exists")
		}
		return "", fmt.Errorf("failed to insert user: %w", err)
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return "", fmt.Errorf("failed to get user ID: %w", err)
	}

	return strconv.FormatInt(userId, 10), nil
}

func isDuplicateKeyError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return true
	}
	return false
}

func (u *UserRepositoryImpl) FetchPasswordAndUserId(email string) (string, string, error) {
	var password, userId string
	query := "SELECT password, id FROM users WHERE email = ?"
	if err := u.db.QueryRow(query, email).Scan(&password, &userId); err != nil {
		return "", "", fmt.Errorf("query failed: %w", err)
	}
	return password, userId, nil
}

func (u *UserRepositoryImpl) FetchUserIdWithEmail(email string) (string, error) {
	var userId string
	query := "SELECT id FROM users WHERE email = ?"
	if err := u.db.QueryRow(query, email).Scan(&userId); err != nil {
		return "", fmt.Errorf("query failed: %w", err)
	}
	return userId, nil
}

func (u *UserRepositoryImpl) UpdateUserHomeFolder(userId string, homeFolder string) error {
	query := "UPDATE users SET home_folder_id = ? WHERE id = ?"
	_, err := u.db.Exec(query, homeFolder, userId)
	if err != nil {
		return fmt.Errorf("query failed: %w", err)
	}
	return nil
}
