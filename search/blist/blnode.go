package blist

import "github.com/davecgh/go-spew/spew"

type Node struct {
	parent *Node

	children [2]*Node
	direct   [2]*Node

	size int64

	key   []byte
	value []byte
}

// Key get node key
func (n *Node) Key() []byte {
	return n.key
}

// Value get node value
func (n *Node) Value() []byte {
	return n.value
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	return spew.Sprint(string(n.key)) + ":" + spew.Sprint(string(n.value))
}

func (n *Node) debugString() string {
	if n == nil {
		return "nil"
	}

	var p string
	if n.parent != nil {
		p = spew.Sprint(string(n.parent.value))
	} else {
		p = "nil"
	}
	return spew.Sprint(string(n.value)) + "(" + p + "|" + spew.Sprint(n.size) + ")"
}
