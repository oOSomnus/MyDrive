package service

import (
	"fmt"
	"github.com/oOSomnus/MyDrive/internal/user/repository"
	"github.com/oOSomnus/MyDrive/pkg/auth"
)

type UserService interface {
	CreateUser(email, password string) error
	Authenticate(email, password string) (authenticated bool, userId string, error)
}
type UserServiceImpl struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository: userRepository}
}

func (u *UserServiceImpl) CreateUser(email, password string) error {
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	err = u.userRepository.InsertUserIfNotExists(email, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

func (u *UserServiceImpl) Authenticate(email, password string) (bool, string, error) {
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return false, "", fmt.Errorf("failed to hash password: %w", err)
	}
	dbPassword, userId, err := u.userRepository.FetchPasswordAndUserId(email)
	if err != nil {
		return false, "", fmt.Errorf("failed to fetch password: %w", err)
	}
	return dbPassword == hashedPassword, userId, nil
}
