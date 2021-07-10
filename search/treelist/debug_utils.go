package listtree

import (
	"fmt"
	"log"
	"runtime"

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

	*str += "(" + spew.Sprint(string(node.key.([]byte))) + "->" + spew.Sprint(string(node.value.([]byte))) + ")" + "\n"

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

var debugCheck map[string]int

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
		parentv = spew.Sprint(string(node.parent.key.([]byte)))
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
	k := string(node.key.([]byte))
	if _, ok := debugCheck[k]; !ok {
		debugCheck[k] = 1
	} else {
		count := debugCheck[k]
		count++
		debugCheck[k] = count
		if count >= 3 {
			runtime.Breakpoint()
			log.Println(node, node.key, node.children)
		}
	}

	*str += spew.Sprint(k) + suffix + "\n"

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

func outputfordebugNoSuffix(node *Node, prefix string, isTail bool, str *string) {

	if node.children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "\033[34m│   \033[0m"
		} else {
			newPrefix += "    "
		}
		outputfordebugNoSuffix(node.children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "\033[34m└── \033[0m"
	} else {
		*str += "\033[31m┌── \033[0m"
	}

	k := string(node.key.([]byte))
	if _, ok := debugCheck[k]; !ok {
		debugCheck[k] = 1
	} else {
		count := debugCheck[k]
		count++
		debugCheck[k] = count
		if count >= 4 {
			runtime.Breakpoint()
		}
	}

	*str += spew.Sprint(k) + "\n"

	if node.children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "\033[31m│   \033[0m"
		}
		outputfordebugNoSuffix(node.children[0], newPrefix, true, str)
	}
}

func (tree *ListTree) debugString(isSuffix bool) string {
	str := "BinarayList\n"
	root := tree.getRoot()
	if root == nil {
		return str + "nil"
	}

	debugCheck = make(map[string]int)
	defer func() { debugCheck = nil }()

	if isSuffix {
		outputfordebug(root, "", true, &str)
	} else {
		outputfordebugNoSuffix(root, "", true, &str)
	}

	var cur = root
	for cur.children[0] != nil {
		cur = cur.children[0]
	}

	// var i = 0
	// str += "\n"
	// start := cur
	// for start != nil {
	// 	str += spew.Sprint(string(start.key.([]byte))) + ","
	// 	start = start.direct[1]
	// 	i++
	// 	if i >= 1000 {
	// 		break
	// 	}
	// }
	// str = str[0:len(str)-1] + "(" + strconv.Itoa(i) + ")"
	// str += "\n" + tree.RotateLog
	return str
}

func (tree *ListTree) debugLookNode(cur *Node) {
	var temp []byte = cur.key.([]byte)
	cur.key = []byte(fmt.Sprintf("\033[32m%s\033[0m", cur.key))
	log.Println(tree.debugString(true))
	cur.key = temp
}
