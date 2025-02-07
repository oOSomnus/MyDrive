package service

import "github.com/oOSomnus/MyDrive/internal/file/repository"

type FileService interface {
}

type FileServiceImpl struct {
	fr repository.FileRepository
}

func NewFileService(fr repository.FileRepository) FileService {
	return &FileServiceImpl{
		fr: fr,
	}
}

func (fs *FileServiceImpl) GetFileStatus(userId, filename, folderId, fingerprint string) {
	
}
