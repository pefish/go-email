package go_email

import "gopkg.in/gomail.v2"

func SendToBy163(
	user string,
	password string,
	to string,
	subject string,
	body string,
) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", user)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	//msg.SetAddressHeader("Cc", "dan@example.com", "Dan")
	//msg.Attach("/home/Alex/lolcat.jpg")

	n := gomail.NewDialer("smtp.163.com", 465, user, password)

	return n.DialAndSend(msg)
}
