package service

type UserService struct {
	userRepo UserRepoInterface
}

type UserRepoInterface interface {
	GetUser() string
}

func NewUserService(userRepo UserRepoInterface) *UserService {
	return &UserService{userRepo}
}

func (us *UserService) GetUserSerivce() string {
	return "Hello " + us.userRepo.GetUser()
}
