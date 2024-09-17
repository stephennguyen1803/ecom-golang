package vo

type UserRegistratorRequest struct {
	Email   string `json:"email,omitempty"`
	Phone   string `json:"phone,omitempty"`
	Purpose string `json:"purpose" binding:"required"`
}
