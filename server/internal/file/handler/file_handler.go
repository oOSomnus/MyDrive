package handler

import "github.com/gin-gonic/gin"

type FileHandler interface {
	UploadFile(c *gin.Context)
}

type FileHandlerImpl struct {
}

func NewFileHandler() FileHandler {
	return &FileHandlerImpl{}
}

func (*FileHandlerImpl) UploadFile(c *gin.Context) {
	
}
