package service

import (
	"ecom-project/internal/repo"
	"ecom-project/pkg/response"
)

// type userService struct {
// 	userRepo UserRepoInterface
// }

// type UserRepoInterface interface {
// 	GetUser() string
// }

// func NewUserService(userRepo UserRepoInterface) *userService {
// 	return &userService{userRepo}
// }

// func (us *userService) GetUserSerivce() string {
// 	return "Hello " + us.userRepo.GetUser()
// }

type IUserService interface {
	RegisterUser(email string, password string) (string, error)
	Register(email string, purpose string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	//1 - check if email is already exist

	if us.userRepo.GetUserByEmail(email) {
		return response.ErrorCodeUserHasExists
	}
	//2 - if not exist, send email

	return response.ErrorCodeSuccess
}

// RegisterUser implements IUserService.
func (us *userService) RegisterUser(email string, password string) (string, error) {
	panic("unimplemented")
}

func NewUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{userRepo}
}
