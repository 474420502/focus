package blist

import "log"

func assertImplementation() {

}

type BinaryList struct {
	compartor Compare
	root      *Node
}

func New() *BinaryList {
	return &BinaryList{
		compartor: CompatorMath,
	}
}

func checkNil(n *Node) string {
	if n == nil {
		return "nil"
	}
	return string(n.key)
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

				// log.Println("now left", "left:", checkNil(right), "right:", checkNil(left))
				node := &Node{parent: cur, key: key, value: value, size: 1}
				cur.children[L] = node
				// node.relation = L

				if right != nil {
					right.direct[L] = node
				}

				if left != nil {
					left.direct[R] = node
				}

				node.direct[R] = right
				node.direct[L] = left

				bl.fixSize(cur)
				if cur.parent != nil && cur.parent.size >= 3 {
					bl.fixBalance(cur.parent)
				}
				return true
			}
		case c > 0:

			left = cur
			if cur.children[R] != nil {
				cur = cur.children[R]
			} else {

				// log.Println("now right", "left:", checkNil(right), "right:", checkNil(left))
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
				if cur.parent != nil && cur.parent.size >= 3 {
					bl.fixBalance(cur.parent)
				}

				return true
			}
		default:
			cur.value = value
			return false
		}
	}

}

func getLeftSize(cur *Node) int64 {
	if cur.children[0] == nil {
		return 0
	}
	return cur.children[0].size
}

func getRightSize(cur *Node) int64 {
	if cur.children[1] == nil {
		return 0
	}
	return cur.children[1].size
}

func (bl *BinaryList) fixBalance(cur *Node) {

	const L = 0
	const R = 1

	for cur != nil {
		lsize, rsize := getChildrenSize(cur)
		parant := cur.parent

		if lsize > rsize {
			diff := lsize - rsize
			if diff >= 2 {
				var mov *Node
				for i := int64(0); i < (diff / 2); i++ {
					mov = cur.direct[L]
				}
				bl.leftUp(cur, mov)
			}

		} else {
			diff := rsize - lsize
			if diff >= 2 {
				var mov *Node
				for i := int64(0); i < (diff / 2); i++ {
					mov = cur.direct[R]
				}
				bl.leftUp(cur, mov)
			}
		}

		cur = parant
	}
}

func (bl *BinaryList) leftUp(cur, mov *Node) {

	const L = 0
	const R = 1

	log.Println("cur:", cur, "mov:", mov)
	parent := cur.parent

	for mov.parent != parent {
		if mov.parent.children[L] == mov {
			bl.rrotate(mov.parent)
		} else {
			bl.lrotate(mov.parent)
		}
	}

	// left := mov
	// for left.children[L] != nil {
	// 	left = left.children[L]
	// }

	// lgroup := left.direct[L]
	// for lgroup.parent == lgroup.direct[L] {
	// 	lgroup = lgroup.parent
	// }

	// right := mov
	// for right.children[R] != nil {
	// 	right = right.children[R]
	// }
	// rgroup := right.direct[R]
	// for rgroup.parent == rgroup.direct[R] {
	// 	rgroup = rgroup.parent
	// }

	// left.direct[L].children[R] = mov.children[L]
	// mov.children[L].parent = left.direct[L].children[R]

	// right.direct[R].children[L] = mov.children[R]
	// mov.children[R].parent = right.direct[R].children[L]

}

func (bl *BinaryList) rightUp(cur, mov *Node) {
	return

	const L = 0
	const R = 1

	log.Println(cur, mov)

	cparent := cur.parent

	leftGroup := mov.children[L]
	if leftGroup == nil {
		leftGroup = mov.direct[L]
	}
	for {

		lchild := leftGroup.children[L]
		var newGroup *Node
		if lchild == nil {
			newGroup = leftGroup.direct[L]
		} else {
			for lchild.children[L] != nil {
				lchild = lchild.children[L]
			}
			newGroup = lchild.direct[L]
		}

		if newGroup == nil || newGroup.parent == cur {
			break
		}

		newGroup.children[R] = leftGroup
		leftGroup.parent = newGroup //缺少size计算

		newGroup.size = getChildrenSumSize(newGroup)

		leftGroup = newGroup

	}

	rightGroup := mov.children[R]
	if rightGroup == nil {
		rightGroup = mov.direct[R]
	}

	for {
		rchild := rightGroup.children[R]
		var newGroup *Node
		if rchild == nil {
			newGroup = rightGroup.direct[R]
		} else {
			for rchild.children[R] != nil {
				rchild = rchild.children[R]
			}
			newGroup = rchild.direct[R]
		}

		if newGroup == nil || newGroup.parent == cur {
			break
		}

		newGroup.children[R] = rightGroup
		rightGroup.parent = newGroup //缺少size计算
		newGroup.size = getChildrenSumSize(newGroup)

		rightGroup = newGroup
	}

	if cparent == nil {
		bl.root = mov
	} else {
		if cparent.children[L] == cur {
			cparent.children[L] = mov
		} else {
			cparent.children[R] = mov
		}
	}
	mov.parent = cparent

	mov.children[L] = leftGroup
	leftGroup.parent = mov
	mov.children[R] = rightGroup
	rightGroup.parent = mov
	mov.size = getChildrenSumSize(mov)
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

	const l = 1
	const r = 0
	// 1 right 0 left
	mov := cur.children[l]
	movright := mov.children[r]

	if cur.parent == nil {
		tree.root = mov
		mov.parent = nil

	} else {
		if cur.parent.children[l] == cur {
			cur.parent.children[l] = mov
		} else {
			cur.parent.children[r] = mov
		}
		mov.parent = cur.parent
	}

	if movright != nil {
		cur.children[l] = movright
		movright.parent = cur

	} else {
		cur.children[l] = nil
	}

	mov.children[r] = cur
	cur.parent = mov

	cur.size = getChildrenSumSize(cur) + 1
	mov.size = getChildrenSumSize(mov) + 1

	return mov
}

func (tree *BinaryList) rrotate(cur *Node) *Node {

	const l = 0
	const r = 1
	// 1 right 0 left
	mov := cur.children[l]
	movright := mov.children[r]

	if cur.parent == nil {
		tree.root = mov
		mov.parent = nil

	} else {
		if cur.parent.children[l] == cur {
			cur.parent.children[l] = mov
		} else {
			cur.parent.children[r] = mov
		}
		mov.parent = cur.parent
	}

	if movright != nil {
		cur.children[l] = movright
		movright.parent = cur

	} else {
		cur.children[l] = nil
	}

	mov.children[r] = cur
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
