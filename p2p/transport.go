package p2p

// Peer is an ionterface that represents a remote node
type Peer interface{}

// transport represents anything that handles the
// communication between nodes in the network. This can be
// of the form (TCP, UDP, Websockets)
type Transport interface{}
