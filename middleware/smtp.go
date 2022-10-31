package middleware

import (
	"fmt"
	"net/smtp"
	"os"
	"time"
)

func SendEmail(to []string, request_type string, requestor string, request_date time.Time) (bool, error) {
	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")
	smtp_host := "smtp.gmail.com"
	smtp_port := "587"
	msg := []byte("Subject: New " + request_type + " Request\r\n" +
		"\r\n" +
		"There has been a new " + request_type + " generated in the financial request platform.\n" +
		requestor + " submitted the request on " + request_date.Local().String() + "\n" +
		`Please login to the finance request hub at https://finance-requests.vercel.app/ for more information.`)
	auth := smtp.PlainAuth("", from, password, smtp_host)
	email_err := smtp.SendMail(smtp_host+":"+smtp_port, auth, from, to, msg)
	fmt.Println("emailing err", email_err)
	if email_err != nil {
		return false, email_err
	}
	return true, nil
}
