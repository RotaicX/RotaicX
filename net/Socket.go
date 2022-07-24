package net

import (
	"github.com/xtaci/kcp-go/v5"
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

// GetData () Read the data from net.Conn and get the body content according to the protocol format
func GetData(conn *kcp.UDPSession) (string, error) {
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

func MakeData(Data string) []byte {
	bData := []byte(Data)
	iDataLen := len(bData)
	sDataLen := strconv.Itoa(iDataLen)
	if len(sDataLen) != 8 {
		complement := 8 - len(sDataLen)
		for l := 0; l != complement; l++ {
			sDataLen = "0" + sDataLen
		}
	}

	bAllData := append([]byte(sDataLen), bData...)
	return bAllData
}
