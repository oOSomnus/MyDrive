package service

type UserService interface {
}
type UserServiceImpl struct {
	userService UserService
}

func NewUserService(userService UserService) *UserServiceImpl {
	return &UserServiceImpl{userService: userService}
}
