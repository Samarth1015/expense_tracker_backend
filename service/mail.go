package service

import (
	"fmt"
	"net/smtp"
	"os"

	loging "github.com/Samarth1015/expense/Loging"
)

func SendMail(toEmail string) error {
	fmt.Print("calling send mail")
	from := os.Getenv("MAIL")
	password := os.Getenv("MAIL_PASSWORD")
	host := "smtp.gmail.com"

	// Its the default port of smtp server
	port := "587"
	msg := "Welcome to Expense Tracker"
	body := []byte(msg)
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, []string{toEmail}, body)
	if err != nil {
		loging.Logger.Error("Error in sending mail")

	}
	return err

}
