package headerlist

import "github.com/btgsuite/btgd/wire"

// Chain is an interface that stores a list of Nodes. Each node represents a
// header in the main chain and also includes a height along with it. This is
// meant to serve as a replacement to list.List which provides similar
// functionality, but allows implementations to use custom storage backends and
// semantics.
type Chain interface {
	// ResetHeaderState resets the state of all nodes. After this method, it will
	// be as if the chain was just newly created.
	ResetHeaderState(Node)

	// Back returns the end of the chain. If the chain is empty, then this
	// return a pointer to a nil node.
	Back() *Node

	// Front returns the head of the chain. If the chain is empty, then
	// this returns a  pointer to a nil node.
	Front() *Node

	// PushBack will push a new entry to the end of the chain. The entry
	// added to the chain is also returned in place.
	PushBack(Node) *Node
}

// Node is a node within the Chain. Each node stores a header as well as a
// height. Nodes can also be used to traverse the chain backwards via their
// Prev() method.
type Node struct {
	// Height is the height of this node within the main chain.
	Height int32

	// Header is the header that this node represents.
	Header wire.BlockHeader

	prev *Node
}

// Prev attempts to access the prior node within the header chain relative to
// this node. If this is the start of the chain, then this method will return
// nil.
func (n *Node) Prev() *Node {
	return n.prev
}

// Equals compares two Nodes if they have the same properties.
// A structure is comparable only if it contains no slice/array.
// wire.BlockHeader has Solution byte[]
func (n *Node) Equals(o Node) bool {
	if n == nil {
		return false
	}

	if n.Height != o.Height {
		return false
	}

	return n.prev == o.prev && Equals(n.Header, o.Header)
}

// Equals is comparing two BlockHeaders if they have the same properties
func Equals(h1 wire.BlockHeader, h2 wire.BlockHeader) bool {
	if h1.Version != h2.Version {
		return false
	}

	if h1.PrevBlock != h2.PrevBlock {
		return false
	}

	if h1.Height != h2.Height {
		return false
	}

	if h1.Reserved != h2.Reserved {
		return false
	}

	if h1.Timestamp != h2.Timestamp {
		return false
	}

	if h1.Bits != h2.Bits {
		return false
	}

	if h1.Nonce != h2.Nonce {
		return false
	}

	if len(h1.Solution) != len(h2.Solution) {
		return false
	}

	for i, v := range h1.Solution {
		if v != h2.Solution[i] {
			return false
		}
	}

	return true
}
