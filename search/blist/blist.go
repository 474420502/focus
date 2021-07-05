package blist

import (
	"fmt"
	"log"
)

func assertImplementation() {

}

type BinaryList struct {
	compartor Compare
	root      *Node
	IsDebug   int
}

func New() *BinaryList {
	return &BinaryList{
		compartor: CompatorMath,
		IsDebug:   -1,
	}
}

func (bl *BinaryList) Put(key, value []byte) bool {
	if bl.root == nil {
		bl.root = &Node{key: key, value: value, size: 1}
		return true
	}

	cur := bl.root

	var left *Node = nil
	var right *Node = nil

	const L = 0
	const R = 1

	for {
		c := bl.compartor(key, cur.key)
		switch {
		case c < 0:

			right = cur
			if cur.children[L] != nil {
				cur = cur.children[L]
			} else {

				node := &Node{parent: cur, key: key, value: value, size: 1}
				cur.children[L] = node

				if right != nil {
					right.direct[L] = node
				}

				if left != nil {
					left.direct[R] = node
				}

				node.direct[L] = left
				node.direct[R] = right

				bl.fixSize(cur)
				if bl.IsDebug >= 0 {
					var temp []byte = node.key
					node.key = []byte(fmt.Sprintf("\033[36m%s\033[0m", node.key))
					defer func() {
						node.key = temp
					}()
				}
				bl.fixBalance(cur.parent)

				return true
			}
		case c > 0:

			left = cur
			if cur.children[R] != nil {
				cur = cur.children[R]
			} else {

				node := &Node{parent: cur, key: key, value: value, size: 1}
				cur.children[R] = node

				if right != nil {
					right.direct[L] = node
				}
				if left != nil {
					left.direct[R] = node
				}

				node.direct[L] = left
				node.direct[R] = right

				bl.fixSize(cur)
				if bl.IsDebug >= 0 {
					var temp []byte = node.key
					node.key = []byte(fmt.Sprintf("\033[36m%s\033[0m", node.key))
					defer func() {
						node.key = temp
					}()
				}
				bl.fixBalance(cur.parent)

				return true
			}
		default:
			cur.value = value
			return false
		}
	}

}

func (bl *BinaryList) fixBalance(cur *Node) {

	if cur == nil {
		return
	}

	bl.blanceSize3(cur)

	var hight = 2
	for {
		if cur.size <= (1<<(hight) - 1) {
			bl.balanceNode(cur)
			break
		}
		cur = cur.parent
		if cur == nil {
			break
		}
		hight++
	}

}

func (bl *BinaryList) balanceNode(cur *Node) {

	if bl.IsDebug >= 0 {
		var temp []byte = cur.key
		cur.key = []byte(fmt.Sprintf("\033[35m%s\033[0m", cur.key))
		log.Println(bl.debugString())
		cur.key = temp
		bl.IsDebug++
	}

	const L = 0
	const R = 1

	var mov *Node = cur

	if cur.size <= 3 {
		bl.blanceSize3(cur)
		return
	}

	lsize, rsize := getChildrenSize(cur)
	if lsize > rsize {

		diff := lsize - rsize
		if diff >= 2 {
			for i := int64(0); i < (diff / 2); i++ {
				mov = mov.direct[L]
			}
			bl.up(cur, mov)
		}

	} else {
		diff := rsize - lsize
		if diff >= 2 {
			for i := int64(0); i < (diff / 2); i++ {
				mov = mov.direct[R]
			}
			bl.up(cur, mov)
		}
	}

	if mov.children[L] != nil {
		bl.balanceNode(mov.children[L])
	}
	if mov.children[R] != nil {
		bl.balanceNode(mov.children[R])
	}

}

func (bl *BinaryList) up(cur, mov *Node) {

	const L = 0
	const R = 1

	// log.Println("cur:", cur, "mov:", mov)
	parent := cur.parent

	for mov.parent != parent {
		if mov.parent.children[L] == mov {
			bl.rrotate(mov.parent)
		} else {
			bl.lrotate(mov.parent)
		}
	}

}

func (bl *BinaryList) fixSize(cur *Node) {

	for cur != nil {
		cur.size++
		cur = cur.parent

	}

}

func (bl *BinaryList) String() string {
	str := "BinaryList:\n"
	if bl.root == nil {
		return str + "nil"
	}
	output(bl.root, "", true, &str)
	return str
}

func (tree *BinaryList) lrotate(cur *Node) *Node {

	const L = 1
	const R = 0
	// 1 right 0 left
	mov := cur.children[L]
	movright := mov.children[R]

	if cur.parent == nil {
		tree.root = mov
		mov.parent = nil

	} else {
		if cur.parent.children[L] == cur {
			cur.parent.children[L] = mov
		} else {
			cur.parent.children[R] = mov
		}
		mov.parent = cur.parent
	}

	if movright != nil {
		cur.children[L] = movright
		movright.parent = cur
	} else {
		cur.children[L] = nil
	}

	mov.children[R] = cur
	cur.parent = mov

	cur.size = getChildrenSumSize(cur) + 1
	mov.size = getChildrenSumSize(mov) + 1

	return mov
}

func (tree *BinaryList) rrotate(cur *Node) *Node {

	const L = 0
	const R = 1
	// 1 right 0 left
	mov := cur.children[L]
	movright := mov.children[R]

	if cur.parent == nil {
		tree.root = mov
		mov.parent = nil

	} else {
		if cur.parent.children[L] == cur {
			cur.parent.children[L] = mov
		} else {
			cur.parent.children[R] = mov
		}
		mov.parent = cur.parent
	}

	if movright != nil {
		cur.children[L] = movright
		movright.parent = cur

	} else {
		cur.children[L] = nil
	}

	mov.children[R] = cur
	cur.parent = mov

	cur.size = getChildrenSumSize(cur) + 1
	mov.size = getChildrenSumSize(mov) + 1

	return mov
}

func getChildrenSumSize(cur *Node) int64 {
	return getSize(cur.children[0]) + getSize(cur.children[1])
}

func getChildrenSize(cur *Node) (int64, int64) {
	return getSize(cur.children[0]), getSize(cur.children[1])
}

func getSize(cur *Node) int64 {
	if cur == nil {
		return 0
	}
	return cur.size
}

func (bl *BinaryList) lrrotate3(cur *Node) {
	const l = 0
	const r = 1

	movleft := cur.children[l]
	mov := movleft.children[r]
	movleft.children[r] = nil

	if cur.parent == nil {
		bl.root = mov
		mov.parent = nil
	} else {
		if cur.parent.children[l] == cur {
			cur.parent.children[l] = mov
		} else {
			cur.parent.children[r] = mov
		}
		mov.parent = cur.parent
	}

	mov.children[l] = movleft
	movleft.parent = mov

	cur.children[l] = nil

	mov.children[r] = cur
	cur.parent = mov

	mov.size = 3
	cur.size = 1
}

func (bl *BinaryList) rlrotate3(cur *Node) {
	const l = 1
	const r = 0

	movleft := cur.children[l]
	mov := movleft.children[r]
	movleft.children[r] = nil

	if cur.parent == nil {
		bl.root = mov
		mov.parent = nil
	} else {
		if cur.parent.children[l] == cur {
			cur.parent.children[l] = mov
		} else {
			cur.parent.children[r] = mov
		}
		mov.parent = cur.parent
	}

	mov.children[l] = movleft
	movleft.parent = mov

	cur.children[l] = nil

	mov.children[r] = cur
	cur.parent = mov

	mov.size = 3
	cur.size = 1
}

func (bl *BinaryList) rrotate3(cur *Node) {
	const l = 0
	const r = 1
	// 1 right 0 left

	mov := cur.children[l]
	if cur.parent == nil {
		bl.root = mov
		mov.parent = nil
	} else {
		if cur.parent.children[l] == cur {
			cur.parent.children[l] = mov
		} else {
			cur.parent.children[r] = mov
		}
		mov.parent = cur.parent
	}
	cur.children[l] = nil

	mov.children[r] = cur
	cur.parent = mov

	mov.size = 3
	cur.size = 1
}

func (bl *BinaryList) lrotate3(cur *Node) {
	const l = 1
	const r = 0

	mov := cur.children[l]
	if cur.parent == nil {
		bl.root = mov
		mov.parent = nil
	} else {
		if cur.parent.children[l] == cur {
			cur.parent.children[l] = mov
		} else {
			cur.parent.children[r] = mov
		}
		mov.parent = cur.parent
	}

	cur.children[l] = nil

	mov.children[r] = cur
	cur.parent = mov

	mov.size = 3
	cur.size = 1
}

func (bl *BinaryList) blanceSize3(cur *Node) {
	const L = 0
	const R = 1

	if cur.size == 3 {
		if cur.children[R] == nil {
			if cur.children[L].children[R] == nil {
				bl.rrotate3(cur)
			} else {
				bl.lrrotate3(cur)
			}
			return
		} else if cur.children[L] == nil {
			if cur.children[R].children[L] == nil {
				bl.lrotate3(cur)
			} else {
				bl.rlrotate3(cur)
			}
			return
		}
	}
}
