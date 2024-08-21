package response

const (
	ErrorCodeSuccess      = 20001 // Success
	ErrorCodeParamInvalid = 20003 // Email is invalid
	ErrorTokenInvalid     = 30001 // Token is invalid

	//Register Code
	ErrorCodeUserHasExists = 50001 // User has exist
)

// message map
var msg = map[int]string{
	ErrorCodeSuccess:       "Success",
	ErrorCodeParamInvalid:  "Email is invalid",
	ErrorTokenInvalid:      "Token is invalid",
	ErrorCodeUserHasExists: "User has exist",
}
