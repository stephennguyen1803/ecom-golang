package service

import (
	"ecom-project/internal/repo"
	"ecom-project/internal/utils/crypto"
	"ecom-project/internal/utils/random"
	"ecom-project/pkg/response"
	"fmt"
	"time"
)

type IUserService interface {
	RegisterUser(email string, password string) (string, error)
	Register(email string, purpose string) int
}

type userService struct {
	userRepo     repo.IUserRepository
	userAuthRepo repo.IUserAuthRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	//0 - hash email - security email - save info into redis
	emailHash := crypto.GetHash(email)
	fmt.Printf("Email hash: %s\n", emailHash)

	//5 - check OTP exist in redis

	//6 - process OTP invalid

	//1 - check if email is already exist in DB
	if us.userRepo.GetUserByEmail(email) {
		return response.ErrorCodeUserHasExists
	}
	//2 - create new OTP
	otp := random.GenSixDigitalOTP()
	if purpose == "TEST-DEV" {
		otp = 123456
	}

	fmt.Printf("OTP: %d\n", otp)

	//3 - save OTP into redis with ttl 5 minutes
	err := us.userAuthRepo.AddOtp(emailHash, otp, int64(10*time.Minute))
	if err != nil {
		return response.ErrorInvalidOTP
	}

	//4 - send OTP to email

	return response.ErrorCodeSuccess
}

// RegisterUser implements IUserService.
func (us *userService) RegisterUser(email string, password string) (string, error) {
	panic("unimplemented")
}

func NewUserService(userRepo repo.IUserRepository, userAuthenRepo repo.IUserAuthRepository) IUserService {
	return &userService{userRepo, userAuthenRepo}
}
