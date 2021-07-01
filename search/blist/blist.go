package blist

import "github.com/davecgh/go-spew/spew"

func assertImplementation() {

}

type BinaryList struct {
	compartor Compare
	root      *Node
}

type sNode struct {
	N *Node
}

func (bl *BinaryList) Put(key, value []byte) bool {
	if bl.root == nil {
		bl.root = &Node{key: key, value: value, size: 1}
		return true
	}

	cur := bl.root
	c := bl.compartor(key, cur.key)
	var paths []*Node = []*Node{cur}

	for {
		switch {
		case c < 0:

			cur = cur.children[0]
			if cur != nil {

			}
			paths = append(paths, cur)

		case c > 0:

			cur = cur.children[1]

		default:
			cur.value = value
			return false
		}

	}

	return false
}

func (bl *BinaryList) String() string {
	str := "BinaryList:\n"
	if bl.root == nil {
		return str + "nil"
	}
	output(bl.root, "", true, &str)
	return str
}

func output(node *Node, prefix string, isTail bool, str *string) {

	if node.children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}

	*str += "(" + spew.Sprint(string(node.key)) + "->" + spew.Sprint(string(node.value)) + ")" + "\n"

	if node.children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.children[0], newPrefix, true, str)
	}

}

func outputfordebug(node *Node, prefix string, isTail bool, str *string) {

	if node.children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		outputfordebug(node.children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}

	suffix := "("
	parentv := ""
	if node.parent == nil {
		parentv = "nil"
	} else {
		parentv = spew.Sprint(string(node.parent.key))
	}
	suffix += parentv + "|" + spew.Sprint(node.size) + ")"
	*str += spew.Sprint(string(node.key)) + suffix + "\n"

	if node.children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		outputfordebug(node.children[0], newPrefix, true, str)
	}
}

func (bl *BinaryList) debugString() string {
	str := "VTree\n"
	if bl.root == nil {
		return str + "nil"
	}
	outputfordebug(bl.root, "", true, &str)
	return str
}
