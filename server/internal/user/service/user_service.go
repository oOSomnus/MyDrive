package service

import (
	"fmt"
	fr "github.com/oOSomnus/MyDrive/internal/file/repository"
	ur "github.com/oOSomnus/MyDrive/internal/user/repository"
	"github.com/oOSomnus/MyDrive/pkg/auth"
)

type UserService interface {
	CreateUser(email, password string) error
	Authenticate(email, password string) (authenticated bool, userId string, error error)
}
type UserServiceImpl struct {
	userRepository ur.UserRepository
	fileRepository fr.FileRepository
}

func NewUserService(userRepository ur.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository: userRepository}
}

func (u *UserServiceImpl) CreateUser(email, password string) error {
	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	userId, err := u.userRepository.InsertUserIfNotExists(email, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	folderId, err := u.fileRepository.CreateBaseFolder(userId)
	if err != nil {
		return fmt.Errorf("failed to create folder: %w", err)
	}
	err = u.userRepository.UpdateUserHomeFolder(userId, folderId)
	if err != nil {
		return fmt.Errorf("failed to update user home: %w", err)
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
