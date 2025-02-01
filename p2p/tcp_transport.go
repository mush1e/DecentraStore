package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPTransport struct {
	listenAddress string
	listener      net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error
	t.listener, err = net.Listen("tcp", t.listenAddress)
	defer t.listener.Close()
	if err != nil {
		return err
	}
	return nil
}

func (t *TCPTransport) acceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("Could not accept connection from %q\n", conn.RemoteAddr().String())
		}
	}
}
