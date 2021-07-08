package listtree

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Llongfile)
}

type Node struct {
	parent   *Node
	children [2]*Node
	direct   [2]*Node

	size  int64
	key   []byte
	value []byte
}

type ListTree struct {
	root    *Node
	compare Compare

	Count int64
}

func New() *ListTree {
	return &ListTree{compare: CompatorMath, root: &Node{}}
}

func (tree *ListTree) getRoot() *Node {
	return tree.root.children[0]
}

func (tree *ListTree) Size() int64 {
	return tree.root.children[0].size
}

func (tree *ListTree) Get(key []byte) ([]byte, bool) {
	const L = 0
	const R = 1

	cur := tree.getRoot()
	for cur != nil {
		c := tree.compare(key, cur.key)
		switch {
		case c < 0:
			cur = cur.children[L]
		case c > 0:
			cur = cur.children[R]
		default:
			return cur.value, true
		}
	}
	return nil, false
}

func (tree *ListTree) Put(key, value []byte) bool {

	cur := tree.getRoot()
	if cur == nil {
		tree.root.children[0] = &Node{key: key, value: value, size: 1, parent: tree.root}
		return true
	}

	var left *Node = nil
	var right *Node = nil

	const L = 0
	const R = 1

	for {
		c := tree.compare(key, cur.key)
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

				tree.fixSize(cur)
				tree.fix(cur.parent)
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

				tree.fixSize(cur)
				tree.fix(cur.parent)
				return true
			}
		default:
			return false
		}
	}
}

func (tree *ListTree) fixSize(cur *Node) {
	for cur != tree.root {
		cur.size++
		cur = cur.parent
	}
}

func (tree *ListTree) fix(cur *Node) {

	const L = 0
	const R = 1

	var temp []byte = cur.key
	cur.key = []byte(fmt.Sprintf("\033[35m%s\033[0m", cur.key))
	log.Println(tree.debugString(false))
	cur.key = temp

	var height int64 = 2

	// var childLimitSize int64 = 1 // 1 << (height - 1) - 1

	for cur != tree.root {

		limitsize := ((int64(1) << height) - 1)
		// (1<< height) -1 允许的最大size　超过证明高度超1
		if cur.size <= limitsize {
			lsize, rsize := getChildrenSize(cur)
			if lsize < rsize {
				diff := (rsize - lsize) / 2
				up := cur.direct[R]
				// 寻找缩小差距的点
				tree.fn0(up, cur, diff, L, R)
			} else {
				diff := (lsize - rsize) / 2
				up := cur.direct[L]
				tree.fn0(up, cur, diff, R, L)
			}
			return
		}

		height++
		cur = cur.parent
	}
}

func (tree *ListTree) fn0(up *Node, cur *Node, diff int64, L int, R int) {

	minDiff := diff

	for up.parent != cur {
		ndiff := diff - up.size
		if ndiff >= 0 {
			if minDiff >= ndiff {
				break
			}
			minDiff = ndiff
		} else {
			if minDiff >= -ndiff {
				break
			}
			minDiff = -ndiff
		}
	}

	tree.debugLookNode(up)

	upLeft := up.children[L]
	upRight := up.children[R]

	var upNewRight *Node
	if up == cur.children[R] {
		upNewRight = cur.children[R].children[R] //符合规律
	} else {
		upNewRight = up.parent
	}

	// 链接当前节点的父节点
	if cur.parent.children[L] == cur {
		cur.parent.children[L] = up
	} else {
		cur.parent.children[R] = up
	}
	up.parent = cur.parent

	// cur的父节点释放, 接下来　关联上up
	up.children[L] = cur
	cur.parent = up

	up.children[R] = upNewRight
	if upNewRight != nil {
		upNewRight.parent = up
		upNewRight.children[L] = upRight
		upNewRight.size = getChildrenSumSize(upNewRight) + 1
		if upRight != nil {
			upRight.parent = upNewRight
		}
	}

	cur.children[R] = upLeft
	if upLeft != nil {
		upLeft.parent = cur
	}

	cur.size = getChildrenSumSize(cur) + 1
	up.size = getChildrenSumSize(up) + 1
}

func (tree *ListTree) lrotate(cur *Node) *Node {

	tree.Count++

	const L = 1
	const R = 0
	// 1 right 0 left
	mov := cur.children[L]
	movright := mov.children[R]

	if cur.parent.children[L] == cur {
		cur.parent.children[L] = mov
	} else {
		cur.parent.children[R] = mov
	}
	mov.parent = cur.parent

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

func (tree *ListTree) rrotate(cur *Node) *Node {

	tree.Count++

	const L = 0
	const R = 1
	// 1 right 0 left
	mov := cur.children[L]
	movright := mov.children[R]

	if cur.parent.children[L] == cur {
		cur.parent.children[L] = mov
	} else {
		cur.parent.children[R] = mov
	}
	mov.parent = cur.parent

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
