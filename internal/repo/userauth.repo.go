package repo

import (
	"ecom-project/global"
	"fmt"
	"time"
)

type IUserAuthRepository interface {
	AddOtp(uniqueData string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}

func (uar *userAuthRepository) AddOtp(uniqueData string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("user:%s:otp", uniqueData) // key = user:email:otp or user:phone:otp
	fmt.Println("key: ", key)
	err := global.Redis.Set(ctx, key, otp, time.Minute*10).Err()
	if err == nil {
		fmt.Printf(global.Redis.Get(ctx, key).Val())
	}

	return err
}
