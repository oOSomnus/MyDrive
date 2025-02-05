package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/oOSomnus/MyDrive/internal/user/service"
)

type UserHandler interface {
	CreateUser(c *gin.Context) error
	Authenticate(c *gin.Context) bool
}

type UserHandlerImpl struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{
		UserService: userService,
	}
}

func (u *UserHandlerImpl) CreateUser(email, password string) error {

}
