package repo

import (
	"ecom-project/global"
	"ecom-project/internal/database"
)

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
	sqlc *database.Queries
}

// GetUserByEmail implements IUserRepository.
func (ur *userRepository) GetUserByEmail(email string) bool {
	//row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	usr, err := ur.sqlc.GetUserByEmailSQLC(ctx, email)
	if err != nil {
		return false
	}
	return usr.UsrID != NumberNull
}

// GetUserByPhone implements IUserRepository.
func (ur *userRepository) GetUserByPhone(phone string) bool {
	//row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_phone = ?", phone).First(&model.GoCrmUser{}).RowsAffected
	usr, err := ur.sqlc.GetUserByPhoneSQLC(ctx, phone)
	if err != nil {
		return false
	}
	return usr.UsrID != NumberNull
}

func NewUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}
