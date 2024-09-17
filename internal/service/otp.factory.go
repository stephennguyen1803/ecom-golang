package service

import (
	"ecom-project/global"
	"ecom-project/internal/repo"
	"errors"
	"fmt"
	"regexp"

	"go.uber.org/zap"
)

type OTPFactory struct {
	userRepo repo.IUserRepository
}

func NewOTPFactory(userRepo repo.IUserRepository) *OTPFactory {
	return &OTPFactory{userRepo: userRepo}
}

// GetOTPService returns the appropriate strategy (email or phone) based on input.
func (f *OTPFactory) GetOTPService(destination string) (OTPService, error) {
	if isValidEmail(destination) {
		return NewEmailOTPService(f.userRepo), nil
	} else if isValidPhone(destination) {
		return NewPhoneOTPService(f.userRepo), nil
	}
	err := errors.New("invalid destination: must provide a valid email or phone number")
	global.Logger.Error(fmt.Sprintf("invalid destination: must provide a valid email or phone number. But the destination is: %s", destination),
		zap.Error(err))
	return nil, err
}

// Helper function to validate email format.
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

// Helper function to validate phone format (e.g., 10-digit number).
func isValidPhone(phone string) bool {
	re := regexp.MustCompile(`^\d{10}$`)
	return re.MatchString(phone)
}
