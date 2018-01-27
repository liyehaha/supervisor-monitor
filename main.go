package main

import (
	"fmt"
	"bufio"
	"os"
	"log"

	"config"
	"notify"
	"monitor"
)

func listen() {
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
			notify.Push(header, payload)
		}
	}
}

func main() {
	c, _ := config.Load("./config.yaml")

	fmt.Println(c)

	notify.Init(c)

	defer func() {
		if err := recover(); err != nil {
			log.Print("panic", err)
		}
	}()
	listen()
}