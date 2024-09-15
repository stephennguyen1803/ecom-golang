package random

import (
	"math/rand"
	"time"
)

// GenSixDigitalOTP generates a 6-digit OTP.
func GenSixDigitalOTP() int {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := 1000000 + rd.Intn(9000000) //100000 to 999999
	return otp
}
