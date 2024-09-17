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
	return global.Redis.Set(ctx, key, otp, time.Duration(expirationTime)).Err()
}
