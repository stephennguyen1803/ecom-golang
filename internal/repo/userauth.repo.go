package repo

import (
	"ecom-project/global"
	"fmt"
	"time"
)

type IUserAuthRepository interface {
	AddOtp(email string, otp int, expirationTime int64) error
}

type userAuthRepository struct{}

func NewUserAuthRepository() IUserAuthRepository {
	return &userAuthRepository{}
}

func (uar *userAuthRepository) AddOtp(email string, otp int, expirationTime int64) error {
	key := fmt.Sprintf("user:%s:otp", email) // key = user:email:otp
	return global.Redis.Set(ctx, key, otp, time.Duration(expirationTime)).Err()
}
