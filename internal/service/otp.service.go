package service

import (
	"ecom-project/internal/repo"
	"ecom-project/internal/utils/sendto"
	"errors"
)

// OTPService defines the interface for sending OTP.
type OTPService interface {
	SendOTP(userIdentify string, otp int) error
	CheckIfUserExists(userIdentify string) error
}

// EmailOTPService is the strategy for sending OTP via email.
type EmailOTPService struct {
	userRepo repo.IUserRepository
}

// Constructor for EmailOTPService
func NewEmailOTPService(userRepo repo.IUserRepository) *EmailOTPService {
	return &EmailOTPService{userRepo: userRepo}
}

// SendOTP implements OTP sending via email.
func (e *EmailOTPService) SendOTP(email string, otp int) error {
	err := sendto.SendTextEmail([]string{email}, "anhdung.phc@gmail.com", otp)
	return err
}

// CheckIfUserExists checks if the email exists in the database.
func (e *EmailOTPService) CheckIfUserExists(email string) error {
	if e.userRepo.GetUserByEmail(email) {
		return errors.New("email already exists")
	}
	return nil
}

// PhoneOTPService is the strategy for sending OTP via phone.
type PhoneOTPService struct {
	userRepo repo.IUserRepository
}

// Constructor for PhoneOTPService
func NewPhoneOTPService(userRepo repo.IUserRepository) *PhoneOTPService {
	return &PhoneOTPService{userRepo: userRepo}
}

func (p *PhoneOTPService) SendOTP(phone string, otp int) error {
	//TODO Logic to send OTP via phone
	return nil
}

// CheckIfUserExists checks if the phone exists in the database.
func (p *PhoneOTPService) CheckIfUserExists(phone string) error {
	if p.userRepo.GetUserByPhone(phone) {
		return errors.New("phone number already exists")
	}
	return nil
}
