package node_test

import (
	"net"
	"testing"

	"github.com/EclesioMeloJunior/distributed/node"
	"github.com/stretchr/testify/assert"
)

func TestNewNode(t *testing.T) {
	ipaddr := net.IP("192.168.0.0")
	port := 90009

	firstNode, err := node.NewNode(ipaddr, port, "")
	assert.Nil(t, err)
	assert.NotEmpty(t, firstNode.ID)
	assert.Len(t, firstNode.ID, node.IDLength)

	ipaddr = net.IP("192.168.1.0")
	port = 90009

	secondNode, err := node.NewNode(ipaddr, port, "")
	assert.Nil(t, err)
	assert.NotEmpty(t, secondNode.ID)
	assert.Len(t, secondNode.ID, node.IDLength)

	assert.NotEqual(t, firstNode.ID, secondNode.ID)
}
