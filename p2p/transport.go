package p2p

// Peer is an interface that
// represents a remote node
type Peer interface {
	Close() error
}

// Handles communication between
// 2 nodes in a network
// TCP, UDP, WebSockets etc
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
