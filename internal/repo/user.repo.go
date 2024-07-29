package repo

type UserRepo struct {
	UserModel UserModelInterface
}

type UserModelInterface interface {
	GetUserName() string
}

func NewUserRepo(User UserModelInterface) *UserRepo {
	return &UserRepo{UserModel: User}
}

func (ur *UserRepo) GetUser() string {
	return "Hello " + ur.UserModel.GetUserName()
}
