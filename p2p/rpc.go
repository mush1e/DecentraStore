package p2p

import "net"

// Message represents any arbitrary data that
// is being sent over each transport between 2 nodes
// in the network
type RPC struct {
	From    net.Addr
	Payload []byte
}
