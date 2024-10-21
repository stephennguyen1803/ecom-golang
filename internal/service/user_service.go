package service

import "context"

type (
	// ... define list user interface in here
	IUserLogin interface {
		Login(ctx context.Context) error
		Register(ctx context.Context) error
		VerifyOTP(ctx context.Context) error
		UpdatePasswordRegister(ctx context.Context) error
	}

	IUserInfo interface {
		GetUserByUserID(ctx context.Context) error
		GetAllUsers(ctx context.Context) error
	}

	IUserAdmin interface {
		RemoveUser(ctx context.Context) error
		FindOneUser(ctx context.Context) error
	}
)

var (
	localUserLogin  IUserLogin
	localIUserInfo  IUserInfo
	localIUserAdmin IUserAdmin
)

func UserAdmin() IUserAdmin {
	if localIUserAdmin == nil {
		panic("implement localUserAdmin not found for interface IUserAdmin")
	}

	return localIUserAdmin
}

func InitUserAdmin(userAdmin IUserAdmin) {
	localIUserAdmin = userAdmin
}

func UserInfo() IUserInfo {
	if localIUserInfo == nil {
		panic("implement localUserInfo not found for interface IUserInfo")
	}

	return localIUserInfo
}

func InitUserInfo(userInfo IUserInfo) {
	localIUserInfo = userInfo
}

func UserLogin() IUserLogin {
	if localUserLogin == nil {
		panic("implement localUserLogin not found for interface IUserLogin")
	}

	return localUserLogin
}

func InitUserLogin(userLogin IUserLogin) {
	localUserLogin = userLogin
}
