package tools

import (
	"fmt"
	"strconv"

	"github.com/wneessen/go-mail"
)

func FormatMail(key, added, removed string) (string, string) {
	var subject string
	var body string

	subject = fmt.Sprintf("[ADGG] Change in AD group: %s", key)
	body = fmt.Sprintf(`
Changes detected in AD group [%s]
- Members added: %s
- Members removed: %s

This mail was sent by ADGG (Active Directory Groups Guard)`,
		key, added, removed)

	return subject, body
}

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
	m.SetImportance(mail.ImportanceHigh)

	if err := c.DialAndSend(m); err != nil {
		logger.Fatal(err.Error())
	}
}
