package repository

import (
	"database/sql"
	"fmt"
)

type UserRepository interface {
	InsertUserIfNotExists(email, hashedPassword string) error
	FetchPasswordAndUserId(email string) (password string, userId string, error error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) InsertUserIfNotExists(email, hashedPassword string) error {
	tx, err := u.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	var exists int
	query := "SELECT COUNT(*) FROM users WHERE email = ? FOR UPDATE"
	if err := tx.QueryRow(query, email).Scan(&exists); err != nil {
		return fmt.Errorf("query failed: %w", err)
	}

	if exists == 0 {
		insertQuery := "INSERT INTO users(email, password) VALUES (?, ?)"
		if _, err := tx.Exec(insertQuery, email, hashedPassword); err != nil {
			return fmt.Errorf("insert failed: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("transaction commit failed: %w", err)
	}
	return nil
}

func (u *UserRepositoryImpl) FetchPasswordAndUserId(email string) (string, string, error) {
	var password, userId string
	query := "SELECT password, id FROM users WHERE email = ?"
	if err := u.db.QueryRow(query, email).Scan(&password, &userId); err != nil {
		return "", "", fmt.Errorf("query failed: %w", err)
	}
	return password, userId, nil
}
