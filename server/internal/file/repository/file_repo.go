package repository

import (
	"database/sql"
	"fmt"
	"strconv"
)

type FileRepository interface {
	CreateBaseFolder(userId string) (folderId string, error error)
}

type FileRepositoryImpl struct {
	db *sql.DB
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
