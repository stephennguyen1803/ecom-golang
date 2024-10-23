package response

const (
	ErrorCodeSuccess      = 20001 // Success
	ErrorCodeParamInvalid = 20003 // Email is invalid
	ErrorTokenInvalid     = 30001 // Token is invalid
	ErrorSendEmail        = 30002 // Send email error
	ErrorSendOTP          = 30003 // Send OTP error

	//Register Code
	ErrorMissingDestinationType = 50000 // Missing destination in headers
	ErrorCodeUserHasExists      = 50001 // User has exist
	ErrorInvalidOTP             = 50002 // Invalid OTP
	ErrorUserBadRequest         = 50003 // Bad request

	//Error Code Login
	ErrorCodeOTPExisted = 60001 // OTP existed
)

// message map
var msg = map[int]string{
	ErrorCodeSuccess:            "Success",
	ErrorCodeParamInvalid:       "Email is invalid",
	ErrorTokenInvalid:           "Token is invalid",
	ErrorCodeUserHasExists:      "User has exist",
	ErrorInvalidOTP:             "Invalid OTP",
	ErrorSendEmail:              "Send email OTP has error",
	ErrorUserBadRequest:         "User is missing information",
	ErrorMissingDestinationType: "Missing destination in headers",
	ErrorSendOTP:                "Send OTP error",
	ErrorCodeOTPExisted:         "OTP existed but not registry",
}
