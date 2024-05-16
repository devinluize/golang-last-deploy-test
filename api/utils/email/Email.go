package email

import (
	"bytes"
	"crypto/tls"
	"errors"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"user-services/api/config"
	"user-services/api/exceptions"
	"user-services/api/utils"

	"github.com/k3a/html2text"
	"gopkg.in/gomail.v2"
)

type EmailData struct {
	URL     string
	Subject string
	Remark  string
	Date    time.Time
}

// ? Email template parser

func ParseTemplateDir(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return template.ParseFiles(paths...)
}

func SendEmail(email string, data *EmailData, emailTemp string) (bool, error) {
	config.InitEnvConfigs(true, "")

	// Sender data.
	from := config.EnvConfigs.SmtpEmailFrom
	smtpPass := config.EnvConfigs.SmtpPass
	smtpUser := config.EnvConfigs.SmtpUser
	to := email
	smtpHost := config.EnvConfigs.SmtpHost
	smtpPort := config.EnvConfigs.SmtpPort
	var body bytes.Buffer

	template, err := ParseTemplateDir("api/templates")
	if err != nil {
		return false, errors.New(utils.CannotSendEmail)
	}

	template.ExecuteTemplate(&body, emailTemp, &data)

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", data.Subject)
	m.SetBody("text/html", body.String())
	m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send Email
	if err := d.DialAndSend(m); err != nil {
		return false, errors.New(utils.CannotSendEmail)
	}
	return true, nil
}

func SendEmails(email []string, data *EmailData, emailTemp string) (bool, *exceptions.BaseErrorResponse) {
	config.InitEnvConfigs(true, "")

	// Sender data.
	from := config.EnvConfigs.SmtpEmailFrom
	smtpPass := config.EnvConfigs.SmtpPass
	smtpUser := config.EnvConfigs.SmtpUser
	recipients := email
	smtpHost := config.EnvConfigs.SmtpHost
	smtpPort := config.EnvConfigs.SmtpPort
	var body bytes.Buffer

	template, err := ParseTemplateDir("api/templates")
	if err != nil {
		return false, &exceptions.BaseErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    utils.CannotSendEmail,
			Err:        err,
		}
	}

	template.ExecuteTemplate(&body, emailTemp, &data)

	m := gomail.NewMessage()
	for _, recipient := range recipients {
		m.SetHeader("From", from)
		m.SetHeader("To", recipient)
		m.SetHeader("Subject", data.Subject)
		m.SetBody("text/html", body.String())
		m.AddAlternative("text/plain", html2text.HTML2Text(body.String()))

		d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

		// Send Email
		if err := d.DialAndSend(m); err != nil {
			return false, &exceptions.BaseErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Message:    utils.CannotSendEmail,
				Err:        err,
			}
		}
	}
	return true, nil
}
