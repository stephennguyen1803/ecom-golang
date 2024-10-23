package impl

import (
	"context"
	"ecom-project/global"
	"ecom-project/internal/consts"
	"ecom-project/internal/database"
	"ecom-project/internal/model"
	"ecom-project/internal/utils"
	"ecom-project/internal/utils/crypto"
	"ecom-project/internal/utils/random"
	"ecom-project/pkg/response"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type sUserLogin struct {
	query *database.Queries
}

func NewUserLogin(query *database.Queries) *sUserLogin {
	return &sUserLogin{
		query: query,
	}
}

func (s *sUserLogin) Login(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	// login
	// 1. hash email
	fmt.Printf("verifyKey: %s\n", in.VerifyKey)
	fmt.Printf("verifyType: %s\n", in.VerifyType)

	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))
	fmt.Printf("hashKey: %s\n", hashKey)

	// 2. check user exist in user base
	userFound, err := s.query.CheckUserBaseExist(ctx, in.VerifyKey)
	if err != nil {
		return response.ErrorCodeUserHasExists, err
	}

	if userFound > 0 {
		return response.ErrorCodeUserHasExists, fmt.Errorf("user has already registered")
	}

	// 3. Create OTP
	userKey := utils.GetUserKey(hashKey)
	otpFound, err := global.Redis.Get(ctx, userKey).Result()
	switch {
	case err == redis.Nil:
		fmt.Println("Key does not exist")
	case err != nil:
		fmt.Println("Get failed:", err)
		return response.ErrorInvalidOTP, err
	case otpFound != "":
		return response.ErrorCodeUserHasExists, fmt.Errorf("user has already registered")
	}

	//4. Generate OTP
	otpNew := random.GenSixDigitalOTP()
	if in.VerifyPurpose == "TEST_USER" {
		otpNew = 123456
	}

	fmt.Printf("OTP is: %d\n", otpNew)
	//5. Save OTP to Redis with expire time
	err = global.Redis.Set(ctx, userKey, strconv.Itoa(otpNew),
		time.Duration(consts.TIME_OTP_REGISTERS*int(time.Minute))).Err()

	if err != nil {
		return response.ErrorInvalidOTP, err
	}

	//6. Send OTP to user

	return nil
}

func (s *sUserLogin) VerifyOTP(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context) error {
	return nil
}
