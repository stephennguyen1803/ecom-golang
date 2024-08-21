package repo

// type userRepo struct {
// 	UserModel UserModelInterface
// }

// type UserModelInterface interface {
// 	GetUserName() string
// }

// func NewUserRepo(User UserModelInterface) *userRepo {
// 	return &userRepo{UserModel: User}
// }

// func (ur *userRepo) GetUser() string {
// 	return "Hello " + ur.UserModel.GetUserName()
// }

type IUserRepository interface {
	GetUserByEmail(email string) bool
	GetUserByPhone(phone string) bool
}

type userRepository struct {
}

// GetUserByEmail implements IUserRepository.
func (ur *userRepository) GetUserByEmail(email string) bool {
	return true
}

// GetUserByPhone implements IUserRepository.
func (ur *userRepository) GetUserByPhone(phone string) bool {
	panic("unimplemented")
}

func NewUserRepository() IUserRepository {
	return &userRepository{}
}
