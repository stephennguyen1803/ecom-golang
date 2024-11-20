package impl

import (
	"context"
	"database/sql"
	"ecom-project/global"
	"ecom-project/internal/consts"
	"ecom-project/internal/database"
	"ecom-project/internal/model"
	"ecom-project/internal/service"
	"ecom-project/internal/utils"
	"ecom-project/internal/utils/crypto"
	"ecom-project/internal/utils/random"
	"ecom-project/pkg/response"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

type sUserLogin struct {
	query      *database.Queries
	otpFactory *service.OTPFactory
}

func NewUserLogin(query *database.Queries, otpFactory *service.OTPFactory) *sUserLogin {
	return &sUserLogin{
		query:      query,
		otpFactory: otpFactory,
	}
}

func (s *sUserLogin) Login(ctx context.Context) error {
	return nil
}

func (s *sUserLogin) Register(ctx context.Context, in *model.RegisterInput) (codeResult int, err error) {
	// login
	strategy, err := s.otpFactory.GetOTPService(in.VerifyType)
	if err != nil {
		return response.ErrorCodeParamInvalid, err
	}
	// 1. hash email
	fmt.Printf("verifyKey: %s\n", in.VerifyKey)
	fmt.Printf("verifyType: %d\n", in.VerifyType)

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

	// 3.1 Check key has already existed in redis
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

	// 6. Send the OTP via the appropriate strategy (email or phone)
	err = strategy.SendOTP(*in, otpNew)
	if err != nil {
		return response.ErrorSendOTP, err
	}

	// 7. Save OTP to database
	results, err := s.query.InsertOTPVerfiy(ctx, database.InsertOTPVerfiyParams{
		VerifyOtp:     strconv.Itoa(otpNew),
		VerifyKey:     in.VerifyKey,
		VerifyKeyHash: hashKey,
		VerifyType:    sql.NullInt32{Int32: int32(in.VerifyType), Valid: true},
		IsVerfified:   sql.NullBool{Bool: false, Valid: true},
		IsDeleted:     sql.NullBool{Bool: false, Valid: true},
	})

	if err != nil {
		return response.ErrorSendEmail, err
	}

	//8. Get LastId
	lastId, err := results.LastInsertId()
	if err != nil {
		return response.ErrorSendEmail, err
	}
	log.Println("lastIdVerifyUser: ", lastId)

	return response.ErrorCodeSuccess, nil
}

func (s *sUserLogin) VerifyOTP(ctx context.Context, in *model.VerifyOTPInput) (out model.VerifyOTPOutput, err error) {
	// logic
	hashKey := crypto.GetHash(strings.ToLower(in.VerifyKey))

	// 1. Get OTP from Redis
	otpFound, err := global.Redis.Get(ctx, utils.GetUserKey(hashKey)).Result()
	if err != nil {
		return out, err
	}

	// 2. Check OTP
	if in.VerifyCode != otpFound {
		// if OTP is invalid 3 times in minutes ??? => send email to user to verify
		return out, fmt.Errorf("OTP is invalid")
	}

	infoOtp, err := s.query.GetInfoOTP(ctx, hashKey)
	if err != nil {
		return out, err
	}

	// 3. Update status verified OTP
	err = s.query.UpdateUserVerificationStatus(ctx, hashKey)
	if err != nil {
		return out, err
	}

	// 4. output - should use key_secret to generate token
	out.Token = infoOtp.VerifyKeyHash
	out.Message = "Success"

	return out, nil
}

func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context) error {
	return nil
}
