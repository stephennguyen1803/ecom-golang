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
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
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

func (s *sUserLogin) Login(ctx context.Context, in model.LoginInput) (Result int, out model.LoginOutput, err error) {
	// Logic login
	// 1. Get user info in User_Base
	userBase, err := s.query.GetOneUserInfo(ctx, in.UserAccount)
	if err != nil {
		return response.ErrorCodeAuthenFailed, out, err
	}

	// 2. Check password
	if !crypto.MatchingPassword(in.UserPassword, userBase.UserSalt, userBase.UserPassword) {
		return response.ErrorCodeAuthenFailed, out, fmt.Errorf("password is incorrect")
	}

	// 3. check two factor authentication
	// 4. Update User login time, login ip in User_Base
	go s.query.LoginUserBase(ctx, database.LoginUserBaseParams{
		UserLoginIp:  sql.NullString{String: "127.0.0.1", Valid: true},
		UserAccount:  in.UserAccount,
		UserPassword: userBase.UserPassword, // password not need
	})

	// 5. Create UUID
	subToken := utils.GeneralCliTokenUUID(int(userBase.UserID))
	global.Logger.Info("subToken: ", zap.String("subToken", subToken))

	// 6. Get user_info Table
	infoUser, err := s.query.GetUser(ctx, uint64(userBase.UserID))
	if err != nil {
		return response.ErrorCodeAuthenFailed, out, err
	}

	// 7. Convert to json to save in redis
	infoUserJson, err := json.Marshal(infoUser)
	if err != nil {
		return response.ErrorCodeAuthenFailed, out, fmt.Errorf("conver to json failed")
	}

	// 8. Save in Redis with key is subToken and save infoUserJson
	err = global.Redis.Set(ctx, subToken, infoUserJson, time.Duration(consts.TIME_OTP_REGISTERS*int(time.Minute))).Err()
	if err != nil {
		return response.ErrorCodeAuthenFailed, out, err
	}

	// 9. Create Token

	return 200, out, nil
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
		fmt.Printf("hashKey: %s\n", hashKey)
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

func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context, token, password string) (userId int64, err error) {
	//1 - token is already verified : user_verified table
	infoOtp, err := s.query.GetInfoOTP(ctx, token)
	if err != nil {
		return response.ErrorCodeUserOTPNotExisted, err
	}

	//2 - Check isVerified is Ok?
	if !infoOtp.IsVerfified.Bool {
		return response.ErrorCodeUserOTPNotExisted, fmt.Errorf("OTP is not verified")
	}

	//3 - Update password, status
	userBase := database.AddUserBaseParams{}
	userBase.UserAccount = infoOtp.VerifyKey
	UserSalt, err := crypto.GeneralSalt(16)
	if err != nil {
		return response.ErrorCodeUserOTPNotExisted, err
	}
	userBase.UserSalt = UserSalt
	userBase.UserPassword = crypto.HashPassword(password, UserSalt)
	//4 - Update user base
	newUserBase, err := s.query.AddUserBase(ctx, userBase)
	if err != nil {
		return response.ErrorCodeUserOTPNotExisted, err
	}
	userId, err = newUserBase.LastInsertId()
	if err != nil {
		return response.ErrorCodeUserOTPNotExisted, err
	}

	//5 - add userId to user_info
	userInfo := database.AddUserHaveUserIdParams{
		UserID:               uint64(userId),
		UserAccount:          infoOtp.VerifyKey,
		UserNickname:         sql.NullString{String: infoOtp.VerifyKey, Valid: true},
		UserAvatar:           sql.NullString{String: "", Valid: true},
		UserState:            1,
		UserMobile:           sql.NullString{String: "", Valid: true},
		UserGender:           sql.NullInt16{Int16: 0, Valid: true},
		UserBirthday:         sql.NullTime{Time: time.Time{}, Valid: false},
		UserEmail:            sql.NullString{String: infoOtp.VerifyKey, Valid: true},
		UserIsAuthentication: 1,
	}

	newUserInfo, err := s.query.AddUserHaveUserId(ctx, userInfo)
	if err != nil {
		return response.ErrorCodeUserOTPNotExisted, err
	}

	newUserInfoId, err := newUserInfo.LastInsertId()
	if err != nil {
		return response.ErrorCodeUserOTPNotExisted, err
	}

	return newUserInfoId, nil
}
