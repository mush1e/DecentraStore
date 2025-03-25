package p2p

// Peer represents a remote node in the network.
type Peer interface{}

// Transport handles communication between nodes in the network.
// This can be TCP, UDP, Websockets etc
type Transport interface {
	ListenAndAccept() error
}
