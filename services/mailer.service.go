package services

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
)

const SUBJECT = "Invitation for Placement Drive at IIT Dharwad"
const BrouchrePath = "../config/brochure.pdf"

var FROM = os.Getenv("SENDER_MAIL")
var PASSWORD = os.Getenv("GOOGLE_APP_PASS")

func setHeaders(msg *gomail.Message, to string) {
	msg.SetHeader("From", FROM)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", SUBJECT)
	msg.Attach(BrouchrePath)
}

func SendMail(companyId uint, template string, to string) error {
	newMessage := gomail.NewMessage()
	setHeaders(newMessage, to)

	dial := gomail.NewDialer("smtp.gmail.com", 587, FROM, PASSWORD)

	err := dial.DialAndSend(newMessage)
	if err != nil {
		fmt.Println("Error in sending mail to", to)
		return err
	}
	return nil
}
