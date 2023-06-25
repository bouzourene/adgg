package tools

import (
	"crypto/tls"
	"fmt"
	"strconv"

	"github.com/wneessen/go-mail"
)

// Fromat mail before sending
func FormatMail(key, added, removed string) (string, string) {
	var subject string
	var body string

	// Create mail subject
	subject = fmt.Sprintf("[ADGG] Change in AD group: %s", key)

	// Create mail body, indentation is super important!!!
	body = fmt.Sprintf(`
>>> IMPORTANT <<<

Changes detected in AD group "%s" :

- Members added: %s
- Members removed: %s

This mail was sent by ADGG (Active Directory Groups Guard)`,
		key, added, removed)

	return subject, body
}

func SendMail(subject, body string) {

	// Get logger
	logger := GetLogger()

	// Log that we're sending an email
	logger.Info(fmt.Sprintf(
		"Sending mail with subject %s",
		subject,
	))

	// Get SMTP port from config and convert to int
	port, err := strconv.Atoi(
		ConfigValue("SMTP_PORT"),
	)
	if err != nil {
		logger.Fatal("SMTP port has to be interger")
	}

	// Create mail client with settings from .env
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

	// If user is empty, do not auth
	if ConfigValue("SMTP_USER") == "" {
		c, err = mail.NewClient(
			ConfigValue("SMTP_ADDR"),
			mail.WithPort(port),
		)

		if err != nil {
			logger.Fatal(err.Error())
		}
	}

	c.SetTLSConfig(
		&tls.Config{InsecureSkipVerify: true},
	)

	if err != nil {
		logger.Fatal(err.Error())
	}

	// Create new mail with settings from input vars
	// as well as from .env
	m := mail.NewMsg()
	m.From(ConfigValue("MAIL_FROM"))
	m.To(ConfigValue("MAIL_TO"))
	m.Subject(subject)
	m.SetBodyString(
		mail.TypeTextPlain,
		body,
	)
	m.SetImportance(mail.ImportanceHigh)

	// Send email or die with fatal
	if err := c.DialAndSend(m); err != nil {
		logger.Fatal(err.Error())
	}
}
