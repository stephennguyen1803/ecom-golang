package service

import (
	"ecom-project/global"
	"ecom-project/internal/consts"
	"errors"
	"fmt"

	"go.uber.org/zap"
)

type OTPFactory struct {
	//userRepo repo.IUserRepository
}

func NewOTPFactory() *OTPFactory {
	return &OTPFactory{}
}

// GetOTPService returns the appropriate strategy (email or phone) based on input.
func (f *OTPFactory) GetOTPService(typeService int) (OTPService, error) {
	if typeService == consts.EMAIL {
		return NewEmailOTPService(), nil
	} else if typeService == consts.MOBILE {
		return NewPhoneOTPService(), nil
	}
	err := errors.New("invalid destination: must provide a valid email or phone number")
	global.Logger.Error(fmt.Sprintf("invalid destination: must provide a valid email or phone number. But the destination is: %d", typeService), zap.Error(err))
	return nil, err
}
