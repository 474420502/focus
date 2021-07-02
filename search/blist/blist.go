package blist

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
					bl.fixBalance(cur)
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
					bl.fixBalance(cur)
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

	lszie, rsize := getChildrenSize(cur)
	if lszie != rsize {
		// var diff int64 = 0
		if lszie > rsize {
			// diff = lszie - rsize
			mid := (cur.size - 1) / 2
			if lszie == mid {
				// right rotate
				bl.rrotate(cur)
			} else if lszie == cur.children[L].children[R].size {
				// (left chilid left rotate) + right rotate

				bl.lrotate(cur.children[L])
				bl.rrotate(cur)
			}
		} else {
			//TODO:

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
