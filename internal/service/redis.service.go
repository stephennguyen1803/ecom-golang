package service

import (
	"ecom-project/internal/repo"
	"time"
)

type RedisService struct {
	userAuthRepo repo.IUserAuthRepository // Interface for Redis operations
}

func NewRedisService(userAuthRepo repo.IUserAuthRepository) *RedisService {
	return &RedisService{userAuthRepo: userAuthRepo}
}

// SaveOTP saves the OTP into Redis with a TTL.
func (s *RedisService) SaveOTP(destinationHash string, otp int, ttl time.Duration) error {
	return s.userAuthRepo.AddOtp(destinationHash, otp, int64(ttl.Seconds()))
}
