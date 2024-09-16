package request

type UserRequestBody struct {
	Email   string `json:"email"`
	Purpose string `json:"purpose"`
}
