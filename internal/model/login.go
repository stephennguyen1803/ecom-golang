package model

type RegisterInput struct {
	VerifyKey     string `json:"verify_key"`
	VerifyType    string `json:"verify_type"`
	VerifyPurpose string `json:"verify_purpose"`
}
