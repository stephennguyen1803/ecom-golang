package initialize

import (
	"ecom-project/global"
	"ecom-project/internal/database"
	"ecom-project/internal/service"
	"ecom-project/internal/service/impl"
)

func InitServicesInteface() {
	queries := database.New(global.Mdbc)
	// User Service Interface
	service.InitUserLogin(impl.NewUserLogin(queries))
}
