package sendto

import (
	"ecom-project/global"
	"fmt"
	"net/smtp"
	"strconv"

	"go.uber.org/zap"
)

const (
	SMTPHOST     = "127.0.0.1"
	SMTPPORT     = "1025"
	SMTPUsername = ""
	SMTPPassword = ""
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
	fmt.Println("To: ", to)
	contentEmail := Mail{
		From:    EmailAddress{Address: form, Name: "Ecom Golang"},
		To:      to,
		Subject: "OTP Verification",
		Object:  "OTP",
		Body:    "Your OTP is " + strconv.Itoa(otp),
	}

	messageEmail := BuildMessage(contentEmail)

	//send email
	auth := smtp.PlainAuth("", SMTPUsername, SMTPPassword, SMTPHOST)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%s", SMTPHOST, SMTPPORT),
		auth,
		form,
		to,
		[]byte(messageEmail))

	if err != nil {
		global.Logger.Error(
			fmt.Sprintf("Send Email Has Error with: form %s to %s with otp %d", form, to, otp), zap.Error(err))
		return err
	}

	return nil

}
