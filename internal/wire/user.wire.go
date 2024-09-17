//go:build wireinject

package wire

import (
	"ecom-project/internal/controller"
	"ecom-project/internal/repo"
	"ecom-project/internal/service"

	"github.com/google/wire"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repo.NewUserRepository,
		repo.NewUserAuthRepository,
		service.NewUserService,
		service.NewOTPFactory,
		service.NewRedisService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
