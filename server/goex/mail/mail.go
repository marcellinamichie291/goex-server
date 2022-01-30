package mail

import "gopkg.in/gomail.v2"

const (
	MAIL_HOST = "smtp.163.com"
	MAIL_PORT = 465
	MAIL_USER = "loading8425@163.com"
	MAIL_PWD  = "ARVEPCMWXDWWNLOP"
)

func SendMail(mailAddress string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "Trading Bot"+"<"+MAIL_USER+">")
	m.SetHeader("To", mailAddress)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	d := gomail.NewDialer(MAIL_HOST, MAIL_PORT, MAIL_USER, MAIL_PWD)
	err := d.DialAndSend(m)
	return err
}
