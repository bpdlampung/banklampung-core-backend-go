package email

import (
	"crypto/tls"
	"fmt"
	"github.com/bpdlampung/banklampung-core-backend-go/enums"
	"github.com/bpdlampung/banklampung-core-backend-go/logs"
	"gopkg.in/gomail.v2"
)

const (
	EmailAttachmentDirectory = "./configs/resources/email-attachment/"
)

type EmailConfig struct {
	emailDialler *gomail.Dialer
	logger       logs.Collections
	mailSender   string
	mailTrap     string
	environment  enums.Environment
}

type SendMail struct {
	Subject string
	Content string
	From    string
	To      []string
	Cc      []string
	Bcc     []string
	Attach  []string
}

func InitConfig(mailPort int, mailHost, mailSender, username, password, mailTrap string, tlsStatus bool, logger logs.Collections, env enums.Environment) EmailConfig {
	emailDialler := gomail.NewDialer(mailHost,
		mailPort,
		username,
		password)

	emailDialler.TLSConfig = &tls.Config{InsecureSkipVerify: tlsStatus}

	return EmailConfig{
		mailSender:   mailSender,
		emailDialler: emailDialler,
		logger:       logger,
		environment:  env,
		mailTrap:     mailTrap,
	}
}

func (cfg EmailConfig) SendMail(payload SendMail) error {
	if cfg.environment != enums.Production {
		payload.To = []string{cfg.mailTrap}
	}

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", fmt.Sprintf("%s <%s>", payload.From, cfg.mailSender))
	mailer.SetHeader("To", payload.To...)
	mailer.SetHeader("Cc", payload.Cc...)
	mailer.SetHeader("Bcc", payload.Bcc...)
	mailer.SetHeader("Subject", payload.Subject)
	mailer.SetBody("text/html", payload.Content)
	if payload.Attach != nil || len(payload.Attach) > 0 {
		for _, val := range payload.Attach {
			mailer.Attach(val)
		}
	}

	err := cfg.emailDialler.DialAndSend(mailer)

	if err != nil {
		cfg.logger.Error(fmt.Sprintf("Mail Cannot Sent!! Error:: %s || Subject:: %s", err.Error(), payload.Subject))
		return err
	}

	cfg.logger.Info(fmt.Sprintf("Mail Sent!! Subject:: %s", payload.Subject))

	return nil
}
