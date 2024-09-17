package service

import (
	"ecom-project/internal/utils/crypto"
	"ecom-project/pkg/response"
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

type IUserService interface {
	RegisterUser(email string, password string) (string, error)
	Register(email string, purpose string) int
}

type userService struct {
	redisService *RedisService
	otpFactory   *OTPFactory
}

// Register implements IUserService.
func (us *userService) Register(userIdentify string, purpose string) int {
	//1. Get the appropriate strategy (email or phone)
	strategy, err := us.otpFactory.GetOTPService(userIdentify)
	if err != nil {
		return response.ErrorCodeParamInvalid
	}

	//5 - check OTP exist in redis

	//6 - process OTP invalid

	//1 - Check if the user (email or phone) already exists
	if err := strategy.CheckIfUserExists(userIdentify); err != nil {
		return response.ErrorCodeUserHasExists // Return if user already exists
	}
	// 3. Generate OTP
	otp := us.generateOTP(purpose)
	fmt.Printf("Generated OTP: %d\n", otp)

	// 4. Hash the userIdentify (email or phone) for Redis
	userIdentifyHash := crypto.GetHash(userIdentify)
	fmt.Printf("Destination hash: %s\n", userIdentifyHash)

	// 5. Save OTP into Redis with TTL (5 minutes)
	err = us.redisService.SaveOTP(userIdentifyHash, otp, 5*time.Minute)
	if err != nil {
		fmt.Println("Error saving OTP to Redis:", err)
		return response.ErrorInvalidOTP
	}

	// 6. Send the OTP via the appropriate strategy (email or phone)
	err = strategy.SendOTP(userIdentify, otp)
	if err != nil {
		return response.ErrorSendOTP
	}

	return response.ErrorCodeSuccess
}

// RegisterUser implements IUserService.
func (us *userService) RegisterUser(email string, password string) (string, error) {
	panic("unimplemented")
}

func NewUserService(redisService *RedisService, otpFactory *OTPFactory) IUserService {
	return &userService{redisService, otpFactory}
}

// Helper function to generate OTP based on the purpose
func (us *userService) generateOTP(purpose string) int {
	if purpose == "TEST-DEV" {
		return 123456
	}
	return rand.Intn(900000) + 100000 // Generates a random 6-digit OTP
}
