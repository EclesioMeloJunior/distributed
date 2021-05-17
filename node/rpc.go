package node

import (
	"log"
)

type (
	RPC struct {
		node *Node
	}

	PingRequest struct {
		IP     string
		Port   string
		NodeID string
	}

	PingReply struct {
		Node *Node
	}
)

func (r *RPC) Ping(payload PingRequest, reply *PingReply) error {
	log.Printf("receiving a ping request from %s:%s - %s", payload.IP, payload.Port, payload.NodeID)
	reply.Node = r.node
	return nil
}

func NewNodeRPC(n *Node) *RPC {
	return &RPC{
		node: n,
	}
}
