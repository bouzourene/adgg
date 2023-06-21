package tools

import (
	"strconv"

	"github.com/wneessen/go-mail"
)

func SendMail(subject, body string) {
	logger := GetLogger()

	port, err := strconv.Atoi(
		ConfigValue("SMTP_PORT"),
	)
	if err != nil {
		logger.Fatal("SMTP port has to be interger")
	}

	c, err := mail.NewClient(
		ConfigValue("SMTP_ADDR"),
		mail.WithPort(port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(ConfigValue("SMTP_USER")),
		mail.WithPassword(ConfigValue("SMTP_PASS")),
	)

	if err != nil {
		logger.Fatal(err.Error())
	}

	m := mail.NewMsg()
	m.From(ConfigValue("MAIL_FROM"))
	m.To(ConfigValue("MAIL_TO"))
	m.Subject(subject)
	m.SetBodyString(
		mail.TypeTextPlain,
		body,
	)

	if err := c.DialAndSend(m); err != nil {
		logger.Fatal(err.Error())
	}
}
