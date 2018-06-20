package main

import (
	"github.com/liyehaha/supervisor-event-listener/listener"
)

func main() {
	for {
		listener.Start()
	}
}
