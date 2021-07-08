package listtree

import (
	"fmt"
	"log"
)

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

				tree.fix(cur, [2]int{0, L})
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

				tree.fix(cur, [2]int{0, R})
				return true
			}
		default:
			return false
		}
	}
}

func (tree *ListTree) fix(cur *Node, relations [2]int) {

	const L = 0
	const R = 1

	cur.size++

	if cur.parent.children[L] == cur {
		relations[0] = L
	} else {
		relations[0] = R
	}

	node := cur.children[relations[1]]
	var temp []byte = node.key
	node.key = []byte(fmt.Sprintf("\033[35m%s\033[0m", node.key))
	log.Println(tree.debugString(false))
	node.key = temp

	cur = cur.parent

	var height int64 = 2

	var childLimitSize int64 = 1 // 1 << (height - 1) - 1

	for cur != tree.root {
		cur.size++
		limitsize := ((int64(1) << height) - 1)

		// (1<< height) -1 允许的最大size　超过证明高度超1
		if cur.size <= limitsize {
			// lsize, rsize := getChildrenSize(cur)

			if relations[0] == R {

				lsize := getSize(cur.children[L])
				if lsize <= childLimitSize {
					if relations[1] == L {
						tree.rrotate(cur.children[R])
					}
					cur = tree.lrotate(cur)
				}

			} else {

				rsize := getSize(cur.children[R])
				if rsize <= childLimitSize {
					if relations[1] == R {
						tree.lrotate(cur.children[L])
					}
					cur = tree.rrotate(cur)
				}
			}

		} else {
			height++
			childLimitSize = limitsize
		}

		relations[1] = relations[0]
		if cur.parent.children[L] == cur {
			relations[0] = L
		} else {
			relations[0] = R
		}
		cur = cur.parent

	}
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