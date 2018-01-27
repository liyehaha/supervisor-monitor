package notify

import (
	"fmt"
	"os"
	"time"

	"config"
	"monitor"
)



type Notifiable interface {
	Send(monitor.Message, *config.Config) error
}

func Push(header *monitor.Header, payload *monitor.Payload, queue chan monitor.Message) {
	queue <- monitor.Message{header, payload}
}

func send(notifyHandler Notifiable, message monitor.Message, c *config.Config) {
	// 最多重试3次
	tryTimes := 3
	i := 0
	for i < tryTimes {
		err := notifyHandler.Send(message, c)
		if err == nil {
			break
		}
		fmt.Fprintln(os.Stderr, err)
		time.Sleep(30 * time.Second)
		i++
	}
}

func Start(c *config.Config, queue chan monitor.Message) {
	var message monitor.Message
	var notifyHandler Notifiable
	for {
		message = <-queue
		notifyHandler = &Mail{}
		/*switch c.NotifyType {
		case "mail":
			notifyHandler = &Mail{}
		case "slack":
			notifyHandler = &Slack{}
		case "webhook":
			notifyHandler = &WebHook{}
		}
		if notifyHandler == nil {
			continue
		}*/
		go send(notifyHandler, message, c)
		time.Sleep(1 * time.Second)
	}
}