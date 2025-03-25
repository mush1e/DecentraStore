package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTRansport(t *testing.T) {
	listenAddr := ":4000"
	tr := NewTCPTransport(listenAddr)
	assert.Equal(t, listenAddr, tr.listenAddress)

	assert.Nil(t, tr.ListenAndAccept())

}
