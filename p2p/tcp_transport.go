package p2p

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// TCPPeer Represents a remote node
// over an established TCP Connection
type TCPPeer struct {
	// this is the underlying connection of the peer
	conn net.Conn

	// if we dial a connection outbound (true)
	// but if we accept and retrieve a connection
	// it is inbound (false)
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandshakeFunc
	decoder       Decoder

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

// Constructor to generate new TCP Transport
func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		shakeHands:    NOPHandshakeFunc,
		listenAddress: listenAddr,
		peers:         make(map[net.Addr]Peer),
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}
	t.listener = ln

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			log.Println("TCP Accept error : ", err)
		}
		go t.handleConnection(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConnection(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.shakeHands(peer); err != nil {
		// Do some shi
	}

	msg := &Temp{}
	// read loop
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("tcp error : %s\n", err)
			continue
		}
	}

	fmt.Printf("connection established with peer : %+v\n", peer)
}
