package services

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
	"strings"
)

const SUBJECT = "Invitation for Placement Drive at IIT Dharwad"
const BrouchrePath = "config/brochure.pdf"

func setHeaders(msg *gomail.Message, from string, to string) {
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", SUBJECT)
	msg.SetHeader("Body")
	msg.Attach(BrouchrePath)
}

func SendMail(template string, to string) error {
	var FROM = os.Getenv("SENDER_MAIL")
	var PASSWORD = os.Getenv("GOOGLE_APP_PASS")

	newMessage := gomail.NewMessage()
	setHeaders(newMessage, FROM, to)
	newMessage.SetBody("text/html", strings.Replace(template, "\n", "<br>", -1))

	dial := gomail.NewDialer("smtp.gmail.com", 587, FROM, PASSWORD)

	err := dial.DialAndSend(newMessage)
	if err != nil {
		fmt.Println("Error in sending mail to", to)
		fmt.Println("[ERROR] ", err)
		return err
	}
	return nil
}
