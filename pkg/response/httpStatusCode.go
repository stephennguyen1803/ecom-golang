package response

const (
	ErrorCodeSuccess      = 20001 // Success
	ErrorCodeParamInvalid = 20003 // Param is invalid
	ErrorTokenInvalid     = 30001 // Token is invalid
	ErrorSendEmail        = 30002 // Send email error
	ErrorSendOTP          = 30003 // Send OTP error

	//User Authentication
	ErrorCodeAuthenFailed = 40005 // Authentication failed

	//Register Code
	ErrorMissingDestinationType = 50000 // Missing destination in headers
	ErrorCodeUserHasExists      = 50001 // User has exist
	ErrorInvalidOTP             = 50002 // Invalid OTP
	ErrorUserBadRequest         = 50003 // Bad request
	ErrorUpdateUserPassword     = 50004 // Update user password error

	//Error Code Login
	ErrorCodeOTPExisted        = 60001 // OTP existed
	ErrorCodeUserOTPNotExisted = 60002 // OTP existed but not registry
)

// message map
var msg = map[int]string{
	ErrorCodeSuccess:            "Success",
	ErrorCodeParamInvalid:       "Param is invalid",
	ErrorTokenInvalid:           "Token is invalid",
	ErrorCodeUserHasExists:      "User has exist",
	ErrorInvalidOTP:             "Invalid OTP",
	ErrorSendEmail:              "Send email OTP has error",
	ErrorUserBadRequest:         "User is missing information",
	ErrorMissingDestinationType: "Missing destination in headers",
	ErrorSendOTP:                "Send OTP error",
	ErrorCodeOTPExisted:         "OTP existed but not registry",
	ErrorCodeUserOTPNotExisted:  "User OTP not exists",
	ErrorUpdateUserPassword:     "Update user password error",
}
