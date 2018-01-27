package monitor

import (
	"bufio"
)

func ReadHeader(reader *bufio.Reader) (*Header, error) {
	// 读取Header
	data, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	// 解析Header
	header, err := ParseHeader(data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

// 读取payload
func ReadPayload(reader *bufio.Reader, payloadLen int) (*Payload, error) {
	// 读取payload
	buf := make([]byte, payloadLen)
	length, err := reader.Read(buf)
	if err != nil {
		return nil, err
	}
	if payloadLen != length {
		return nil, nil
	}
	// 解析payload
	payload, err := ParsePayload(string(buf))
	if err != nil {
		return nil, err
	}

	return payload, nil
}