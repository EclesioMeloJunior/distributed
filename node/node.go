package node

import (
	"errors"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
	dbpath, err := expandHomeDirectory(filepath.Join(basepath, "register.db"))
	if err != nil {
		return nil, err
	}

	_, err = os.Stat(dbpath)
	if err != nil {
		if err = createDefaultDirectory(basepath); err != nil {
			return nil, err
		}
	}

	database, err := bolt.Open(dbpath, os.ModePerm, nil)
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

func createDefaultDirectory(dir string) error {
	expandedDir, err := expandHomeDirectory(dir)
	if err != nil {
		return err
	}

	return os.MkdirAll(expandedDir, os.ModeDir)
}

func expandHomeDirectory(dir string) (string, error) {
	if strings.HasPrefix(dir, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}

		dir = strings.Replace(dir, "~", home, -1)
	}

	return dir, nil
}
