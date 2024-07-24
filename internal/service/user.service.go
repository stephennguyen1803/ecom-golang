package service

type UserService struct {
	userRepo UserRepoInterface
}

type UserRepoInterface interface {
	GetUserByName(name string) string
}

func NewUserService(userRepo UserRepoInterface) *UserService {
	return &UserService{userRepo}
}

func (us *UserService) GetUserName(name string) string {
	return "Hello " + us.userRepo.GetUserByName(name)
}
