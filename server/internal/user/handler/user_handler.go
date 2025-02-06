package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/oOSomnus/MyDrive/internal/dto"
	"github.com/oOSomnus/MyDrive/internal/user/service"
	"github.com/oOSomnus/MyDrive/pkg/auth"
	"net/http"
	"time"
)

type UserHandler interface {
	CreateUser(c *gin.Context) error
	Authenticate(c *gin.Context) bool
}

type UserHandlerImpl struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{
		userService: userService,
	}
}

func (u *UserHandlerImpl) CreateUser(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if !auth.IsValidEmail(email) {
		c.JSON(
			http.StatusBadRequest, dto.ResponseEntity{Message: "Invalid email"},
		)
		return
	}
	if len(password) < 8 || len(password) > 22 {
		c.JSON(http.StatusBadRequest, dto.ResponseEntity{Message: "Password must be between 8 and 22 characters"})
		return
	}
	if err := u.userService.CreateUser(email, password); err != nil {
		c.JSON(
			http.StatusOK, dto.ResponseEntity{
				Message: "Register succeed!",
			},
		)
		return
	}

	c.JSON(
		http.StatusBadRequest, dto.ResponseEntity{
			Message: "Failed to register user, please try again later",
		},
	)
}

func (u *UserHandlerImpl) Authenticate(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	isAuthenticated, userId, err := u.userService.Authenticate(email, password)
	if err != nil || !isAuthenticated {
		c.JSON(
			http.StatusBadRequest, dto.ResponseEntity{
				Message: "Failed to authenticate, please check the email and the password",
			},
		)
		return
	}
	token, err := auth.GenerateToken(userId, 3*time.Hour)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError, dto.ResponseEntity{
				Message: "Failed to authenticate, please try again later",
			},
		)
		return
	}
	c.JSON(
		http.StatusOK, dto.ResponseEntity{
			Message: "Successfully authenticated",
			Data: map[string]string{
				"token": token,
			},
		},
	)
}
