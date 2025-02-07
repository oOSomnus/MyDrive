package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/redis/go-redis/v9"
	"strconv"
)

type FileRepository interface {
	CreateBaseFolder(userId string) (folderId string, error error)
	FetchRegisteredUploadId(userId, folderId, filename, fingerprint string) (
		UploadId string, exists bool, error error,
	)
	FetchFolderOwner(folderId string) (ownerId string, error error)
}

type FileRepositoryImpl struct {
	db          *sql.DB
	redisClient *redis.Client
}

func NewFileRepository(db *sql.DB, redisClient *redis.Client) FileRepository {
	return &FileRepositoryImpl{db: db, redisClient: redisClient}
}

func (f *FileRepositoryImpl) CreateBaseFolder(userId string) (string, error) {
	query := "INSERT INTO folders (owner_id, folder_name, parent_folder_id ) VALUES (?, ?, ?)"
	result, err := f.db.Exec(query, userId, "root", "-1")
	if err != nil {
		return "", fmt.Errorf("failed to create base folder: %v", err)
	}
	folderId, err := result.LastInsertId()
	if err != nil {
		return "", fmt.Errorf("failed to create base folder: %v", err)
	}
	return strconv.FormatInt(folderId, 10), nil
}

func (f *FileRepositoryImpl) FetchRegisteredUploadId(userId, folderId, filename, fingerprint string) (
	string, bool, error,
) {
	key := "uploadId" + "/" + userId + "/" + folderId + "/" + filename + "/" + fingerprint
	exists, err := f.redisClient.Exists(context.Background(), key).Result()
	if err != nil {
		return "", false, fmt.Errorf("failed to fetch registered upload id: %w", err)
	}
	if exists > 0 {
		val, err := f.redisClient.Get(context.Background(), key).Result()
		if err != nil {
			return "", false, fmt.Errorf("failed to fetch registered upload id: %w", err)
		}
		return val, true, nil
	} else {
		return "", false, nil
	}
}

func (f *FileRepositoryImpl) FetchFolderOwner(folderId string) (string, error) {
	query := "SELECT owner_id FROM folders WHERE folder_id = ?"
	var userId string
	err := f.db.QueryRow(query, folderId).Scan(&userId)
	if err != nil {
		return "", fmt.Errorf("failed to fetch owner: %w", err)
	}
	return userId, nil
}
