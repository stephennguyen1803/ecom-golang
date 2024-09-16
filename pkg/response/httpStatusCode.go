package response

const (
	ErrorCodeSuccess      = 20001 // Success
	ErrorCodeParamInvalid = 20003 // Email is invalid
	ErrorTokenInvalid     = 30001 // Token is invalid
	ErrorSendEmail        = 30002 // Send email error

	//Register Code
	ErrorCodeUserHasExists = 50001 // User has exist
	ErrorInvalidOTP        = 50002 // Invalid OTP
	ErrorUserBadRequest    = 50003 // Bad request
)

// message map
var msg = map[int]string{
	ErrorCodeSuccess:       "Success",
	ErrorCodeParamInvalid:  "Email is invalid",
	ErrorTokenInvalid:      "Token is invalid",
	ErrorCodeUserHasExists: "User has exist",
	ErrorInvalidOTP:        "Invalid OTP",
	ErrorSendEmail:         "Send email OTP has error",
	ErrorUserBadRequest:    "User is missing information",
}
