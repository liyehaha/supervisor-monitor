package main

import (
	"bufio"
	"os"
	"log"

	"config"
	"notify"
	"monitor"
)

func listen(queue chan monitor.Message) {
	reader := bufio.NewReader(os.Stdin)
	for {
		header, err := monitor.ReadHeader(reader)
		if err != nil {
			log.Print(err)
			continue
		}
		payload, err := monitor.ReadPayload(reader, header.Len)
		if err != nil {
			log.Print(err)
			continue
		}
		// 只处理进程异常退出事件
		if header.EventName == "PROCESS_STATE_EXITED" {
			notify.Push(header, payload, queue)
		}
	}
}

func main() {
	c, _ := config.Load("./config.yaml")

	var queue chan monitor.Message
	
	queue = make(chan monitor.Message, 10)

	go notify.Start(c, queue)
	defer func() {
		if err := recover(); err != nil {
			log.Print("panic", err)
		}
	}()
	listen(queue)
}