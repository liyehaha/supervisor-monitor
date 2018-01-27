package notify

import (
	"fmt"
	"strings"
	"errors"
	
	"config"
	"monitor"

	gomail "github.com/go-mail/gomail"
)

type Mail struct{}

func (mail *Mail) Send(message monitor.Message, conf *config.Config) error {
	body := message.String()
	body = strings.Replace(body, "\n", "<br>", -1)
	gomailMessage := gomail.NewMessage()
	gomailMessage.SetHeader("From", conf.Mail.User)
	gomailMessage.SetHeader("To", config.SplitRepcients(conf.Mail.Repcients)...)
	gomailMessage.SetHeader("Subject", "Supervisor事件通知")
	gomailMessage.SetBody("text/html", body)
	mailer := gomail.NewPlainDialer(
		conf.Mail.Host,
		conf.Mail.Port,
		conf.Mail.User,
		conf.Mail.Password,
	)
	err := mailer.DialAndSend(gomailMessage)
	if err == nil {
		return nil
	}

	return errors.New(fmt.Sprintf("邮件发送失败#%s", err.Error()))
}