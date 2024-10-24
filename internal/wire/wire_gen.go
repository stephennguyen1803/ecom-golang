// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"ecom-project/internal/controller"
	"ecom-project/internal/repo"
	"ecom-project/internal/service"
)

// Injectors from user.wire.go:

func InitUserRouterHandler() (*controller.UserController, error) {
	iUserAuthRepository := repo.NewUserAuthRepository()
	redisService := service.NewRedisService(iUserAuthRepository)
	otpFactory := service.NewOTPFactory()
	iUserService := service.NewUserService(redisService, otpFactory)
	userController := controller.NewUserController(iUserService)
	return userController, nil
}
