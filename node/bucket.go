package node

import (
	"errors"
	"fmt"
)

const (
	kbucket = 8
)

// sendToTail sends a peer to the tail of the array
func (n *Node) updatePeers(newpeer *Peer) error {
	n.mu.Lock()
	defer n.mu.Unlock()

	plen := newpeer.ID.XOR(n.ID).PrefixLen()

	peers := n.peers[plen]
	peerPos, ok := n.getPeer(plen, newpeer)

	// if peers does not exists
	if !ok {
		// there is free space
		if len(peers) < kbucket {
			n.peers[plen][len(peers)-1] = newpeer
			return nil
		} else {
			_, err := n.PingPeer(peers[0])

			// if the peer at head does not respond
			// then remove it from head and add the new at tail
			if errors.Is(err, ErrUnreachablePeer) {
				peersSlice := append(peers[1:], newpeer)
				copy(n.peers[plen][:], peersSlice)
				return nil
			}

			if err != nil {
				return fmt.Errorf("could not ping peer: %w", err)
			}

			// just move the peer from the head to tail
			// and ignore the toKeep peer
			peersSlice := append(peers[1:], peers[0])
			copy(n.peers[plen][:], peersSlice)
			return nil
		}
	}

	// if there's 0 or 1 elem in array then nothing happens
	if len(n.peers) <= 1 {
		return nil
	}

	peersSlice := append(peers[:peerPos], peers[peerPos+1:]...)
	peersSlice = append(peersSlice, newpeer)

	copy(n.peers[plen][:], peersSlice)
	return nil
}

// getPeer return the position and true if peer exists
// return -1 and false if it does not
func (n *Node) getPeer(plen int, p *Peer) (int, bool) {
	peers := n.peers[plen]

	for i, peer := range peers {
		if peer.ID == p.ID {
			return i, true
		}
	}

	return -1, false
}
