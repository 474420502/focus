package treelist

import (
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

func output(node *Node, prefix string, isTail bool, str *string) {

	if node.children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "\033[34m│   \033[0m"
		} else {
			newPrefix += "    "
		}
		output(node.children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "\033[34;40m└── \033[0m"
	} else {
		*str += "\033[31;40m┌── \033[0m"
	}

	*str += "(" + spew.Sprint(string(node.key)) + "->" + spew.Sprint(string(node.value)) + ")" + "\n"

	if node.children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "\033[31m│   \033[0m"
		}
		output(node.children[0], newPrefix, true, str)
	}

}

func outputfordebug(node *Node, prefix string, isTail bool, str *string) {

	if node.children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "\033[34m│   \033[0m"
		} else {
			newPrefix += "    "
		}
		outputfordebug(node.children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "\033[34m└── \033[0m"
	} else {
		*str += "\033[31m┌── \033[0m"
	}

	suffix := "("
	parentv := ""
	if node.parent == nil {
		parentv = "nil"
	} else {
		parentv = spew.Sprint(string(node.parent.key))
	}

	// var ldirect, rdirect string
	// if node.direct[0] != nil {
	// 	ldirect = spew.Sprint(string(node.direct[0].value))
	// } else {
	// 	ldirect = "nil"
	// }

	// if node.direct[1] != nil {
	// 	rdirect = spew.Sprint(string(node.direct[1].value))
	// } else {
	// 	rdirect = "nil"
	// }

	// suffix += parentv + "|" + spew.Sprint(node.size) + " " + ldirect + "<->" + rdirect + ")"
	suffix += parentv + "|" + spew.Sprint(node.size) + ")"
	// suffix = ""
	*str += spew.Sprint(string(node.key)) + suffix + "\n"

	if node.children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "\033[31m│   \033[0m"
		}
		outputfordebug(node.children[0], newPrefix, true, str)
	}
}
func (tree *TreeList) debugString() string {
	str := "BinarayList\n"
	root := tree.root.children[0]
	if root == nil {
		return str + "nil"
	}
	outputfordebug(root, "", true, &str)

	var cur = root
	for cur.children[0] != nil {
		cur = cur.children[0]
	}

	var i = 0
	str += "\n"
	start := cur
	for start != nil {
		str += spew.Sprint(string(start.key)) + ","
		start = start.direct[1]
		i++
		if i >= 1000 {
			break
		}
	}
	str = str[0:len(str)-1] + "(" + strconv.Itoa(i) + ")"

	return str
}
