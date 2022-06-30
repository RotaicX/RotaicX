package net

import (
	"github.com/xtaci/kcp-go/v5"
	"net"
	"strconv"
)

func InitKCP(Addr string) (*kcp.Listener, error) {
	kcpObject := KCP{Addr: Addr}
	listener, err := kcpObject.Listen()

	if err != nil {
		return nil, err
	}

	return listener, nil
}

func GetData(conn net.Conn) (string, error) {
	dataLength := make([]byte, 8)
	_, err := conn.Read(dataLength)
	if err != err {
		return "", err
	}
	for _, v := range dataLength {
		if v == '0' {
			dataLength = dataLength[1:]
		} else {
			break
		}
	}
	sDataLen := string(dataLength)
	iDataLen, err := strconv.Atoi(sDataLen)
	if err != nil {
		return "", err
	}
	data := make([]byte, iDataLen)
	_, err = conn.Read(data)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
