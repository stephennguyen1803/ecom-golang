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
	Body    string       `json:"body"`
}

// Function to build the email message with proper formatting
func BuildMessage(mail Mail) []byte {
	// Format the headers and body according to email standards
	msg := fmt.Sprintf("From: %s <%s>\r\n", mail.From.Name, mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", mail.To)
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += "MIME-Version: 1.0\r\n"
	msg += "Content-Type: text/html; charset=\"UTF-8\"\r\n"
	msg += "\r\n"                           // Separates the headers from the body
	msg += fmt.Sprintf("%s\r\n", mail.Body) // Email body

	return []byte(msg) // Return as a byte slice
}

func SendTextEmail(to []string, form string, otp int) error {
	fmt.Println("To: ", to)
	contentEmail := Mail{
		From:    EmailAddress{Address: form, Name: "Ecom Golang"},
		To:      to,
		Subject: "OTP Verification",
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
		messageEmail)

	if err != nil {
		global.Logger.Error(
			fmt.Sprintf("Send Email Has Error with: form %s to %s with otp %d", form, to, otp), zap.Error(err))
		return err
	}

	return nil

}
