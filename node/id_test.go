package node_test

import (
	"testing"

	"github.com/EclesioMeloJunior/distributed/node"
	"github.com/stretchr/testify/assert"
)

func TestPrefixLen(t *testing.T) {
	testcases := []struct {
		id       node.UID
		expected int
	}{
		{
			id:       node.UID([node.IDLength]byte{0, 0, 0, 0, 0, 0, 0, 0}),
			expected: node.IDLength*8 - 1,
		},
		{
			id:       node.UID([node.IDLength]byte{1, 0, 0, 0, 0, 0, 0, 0}),
			expected: 0*8 - 7,
		},
	}

	for _, test := range testcases {
		plen := test.id.PrefixLen()
		assert.Equal(t, test.expected, plen)
	}

}
