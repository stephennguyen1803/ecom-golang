package sendto

import (
	"fmt"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}

type Mail struct {
	From    EmailAddress `json:"from"`
	To      []string     `json:"to"`
	Subject string       `json:"subject"`
	Object  string       `json:"object"`
	Body    string       `json:"body"`
}

func BuildMessage(mail Mail) string {
	msg := "MINE-Version: 1.0;\nContent-Type: text/html;\ncharset=\"UTF-8\";\n\n"
	msg += fmt.Sprintf("From: %s <%s>\n", mail.From.Name, mail.From.Address)
	msg += fmt.Sprintf("To: %s\n", mail.To)
	msg += fmt.Sprintf("Subject: %s\n", mail.Subject)
	msg += fmt.Sprintf("Object: %s\n", mail.Object)
	msg += fmt.Sprintf("Body: %s\n", mail.Body)
	return msg
}

func SendTextEmail(to []string, form string, otp int) error {
	contentEmail := Mail{
		From:    EmailAddress{Address: form, Name: "Ecom Golang"},
		To:      to,
		Subject: "OTP Verification",
		Object:  "OTP",
		Body:    "Your OTP is " + string(otp),
	}

	messageEmail := BuildMessage(contentEmail)

	//send email
	//auth := smtp.PlainAuth("",)
}
