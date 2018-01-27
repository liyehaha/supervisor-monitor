package monitor

import "fmt"

type Header struct {
	Ver        string
	Server     string
	Serial     int
	Pool       string
	PoolSerial int
	EventName  string
	Len        int
}

type Payload struct {
	ProcessName string
	GroupName   string
	FromState   string
	Expected    int
	Time		string
}

type Message struct {
	Header  *Header
	Payload *Payload
}

type Fields map[string]string


func (msg *Message) String() string {
	return fmt.Sprintf("Process: %s\nEXITED FROM state: %s\nAT_TIME: %s\n", msg.Payload.ProcessName, msg.Payload.FromState, msg.Payload.Time)

}