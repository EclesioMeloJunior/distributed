package node

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"math/big"
	"math/rand"
	"time"
)

const (
	IDLength = 8
)

type (
	UID [IDLength]byte
)

// String just format the UID content
func (n *UID) String() string {
	return hex.EncodeToString(n[:])
}

// Equals verify if the current UID is equals other
func (n *UID) Equals(other UID) bool {
	return bytes.Equal(n[:], other[:])
}

// Less verify if the current node is less than other
func (n *UID) Less(other UID) bool {
	return bytes.Compare(n[:], other[:]) == -1
}

// XOR return the xor distance between 2 UIDs
func (n *UID) XOR(other UID) UID {
	var ret UID

	for i := 0; i < IDLength; i++ {
		ret[i] = n[i] ^ other[i]
	}

	return ret
}

// PrefixLen returns the amount of leading 0's of a UID
func (n UID) PrefixLen() int {
	for i := 0; i < IDLength; i++ {
		for bit := 0; bit < 8; bit++ {
			if (n[i]>>uint8(7-bit))&0x1 != 0 {
				return i*8 - bit
			}
		}
	}

	return IDLength*8 - 1
}

// NewNodeUID returns an unique id for a node
func NewNodeUID(ip, port string) UID {
	cs := make([]byte, 10)
	rand.Read(cs)

	h := sha1.New()
	h.Write(bytes.Join([][]byte{cs, []byte(ip + port)}, []byte{}))

	var id UID
	copy(id[:], h.Sum(nil)[:IDLength])

	return id
}

func NewFileUID(name string) UID {
	ts := time.Now().UTC().Unix()

	rnd := make([]byte, 10)
	rand.Read(rnd)

	h := sha1.New()
	h.Write(bytes.Join([][]byte{rnd, big.NewInt(ts).Bytes()}, []byte{}))

	var id UID
	copy(id[:], h.Sum(nil)[:IDLength])

	return id
}
