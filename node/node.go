package node

import (
	"errors"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/boltdb/bolt"
)

var (
	ErrUnreachablePeer = errors.New("peer could not be reached")
)

type (
	Peer struct {
		ID UID

		IPAddr net.IP
		Port   int
	}

	Node struct {
		ID UID

		IPAddr net.IP
		Port   int

		database *bolt.DB
		basepath string
		files    []File
		mu       *sync.Mutex
		peers    [IDLength][kbucket]*Peer
	}
)

// PingPeer will send a ping message to the peer
func (n *Node) PingPeer(peer *Peer) ([]byte, error) {
	return nil, errors.New("not implemented yet")
}

func (n *Node) FindClosest(peer *Node) ([]*Node, error) {
	return nil, nil
}

func NewNode(ip net.IP, port int, basepath string) (*Node, error) {
	database, err := bolt.Open(filepath.Join(basepath, "register.db"), os.ModePerm, nil)
	if err != nil {
		return nil, err
	}

	UID := NewNodeUID(ip.String(), strconv.Itoa(port))

	return &Node{
		IPAddr:   ip,
		Port:     port,
		mu:       &sync.Mutex{},
		ID:       UID,
		basepath: basepath,
		database: database,
	}, nil
}
