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

func (ur *UserRepo) GetUserByName(name string) string {
	return "Hello " + ur.UserModel.GetUserName()
}
