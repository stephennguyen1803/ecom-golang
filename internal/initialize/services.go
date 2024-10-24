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
	otpFactory := service.NewOTPFactory() // Assuming you have a function to create an OTPFactory instance
	service.InitUserLogin(impl.NewUserLogin(queries, otpFactory))
}
