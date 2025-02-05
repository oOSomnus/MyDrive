package repository

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrStartTransaction   = errors.New("failed to start transaction")
	ErrProceedTransaction = errors.New("failed to proceed transaction")
)

type UserRepository interface {
	InsertUserIfNotExists(email, password string) error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u *UserRepositoryImpl) InsertUserIfNotExists(email, password string) error {
	tx, err := u.db.Begin()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrStartTransaction, err)
	}

	defer func() {
		if err != nil {
			err := tx.Rollback()
			if err != nil {
				return
			}
		}
	}()

	var exists int
	query := "SELECT COUNT(*) FROM users WHERE email = ? FOR UPDATE"
	err = tx.QueryRow(query, email).Scan(&exists)
	if err != nil {
		return fmt.Errorf("%w: query failed (%s) - %v", ErrProceedTransaction, query, err)
	}

	if exists == 0 {
		insertQuery := "INSERT INTO users(email, password) VALUES (?, ?)"
		_, err = tx.Exec(insertQuery, email, password)
		if err != nil {
			return fmt.Errorf("%w: insert failed (%s) - %v", ErrProceedTransaction, insertQuery, err)
		}
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("transaction commit failed: %v", err)
	}
	return nil
}
