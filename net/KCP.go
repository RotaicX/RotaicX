package net

import (
	"github.com/xtaci/kcp-go/v5"
)

type KCP struct {
	Addr string
}

func (receiver *KCP) Listen() (*kcp.Listener, error) {
	listen, err := kcp.ListenWithOptions(receiver.Addr, nil, 10, 3)
	return listen, err
}
