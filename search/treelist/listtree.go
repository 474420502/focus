package listtree

import (
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

	Count     int64
	RotateLog string
}

func New() *ListTree {
	return &ListTree{compare: CompatorMath, root: &Node{}}
}

func (tree *ListTree) getRoot() *Node {
	return tree.root.children[0]
}

func (tree *ListTree) Size() int64 {
	if root := tree.getRoot(); root != nil {
		return root.size
	}
	return 0
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
	tree.RotateLog = ""
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
				tree.fixPut(cur)
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
				tree.fixPut(cur)
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

func (tree *ListTree) fixPut(cur *Node) {

	const L = 0
	const R = 1

	// var temp []byte = cur.key
	// cur.key = []byte(fmt.Sprintf("\033[35m%s\033[0m", cur.key))
	// log.Println(tree.debugString(false))
	// cur.key = temp

	var height int64 = 2
	// var childLimitSize int64 = 1 // 1 << (height - 1) - 1

	var relations int
	if cur.parent.children[R] == cur {
		relations = R
	}
	cur = cur.parent

	for cur != tree.root {

		limitsize := ((int64(1) << height) - 1)
		// (1<< height) -1 允许的最大size　超过证明高度超1, 并且有最少１size的空缺
		if cur.size <= limitsize {

			childlimit := (int64(1) << (height - 2)) - 1

			// 右就检测左边
			if relations == R {
				lsize := getSize(cur.children[L])
				if lsize <= childlimit { // 3
					// tree.debugLookNode(cur)
					if lsize == childlimit {
						if checkEqual(cur.children[L]) {
							tree.avlrrotate(cur)
							return
						}
					} else {
						tree.avlrrotate(cur)
						return
					}

				}

			} else {

				rsize := getSize(cur.children[R])
				if rsize <= childlimit { // 3
					// tree.debugLookNode(cur)

					if rsize == childlimit {
						if checkEqual(cur.children[R]) {
							tree.avllrotate(cur)
							return
						}
					} else {
						tree.avllrotate(cur)
						return
					}

				}
			}
		}

		height++

		if cur.parent.children[R] == cur {
			relations = R
		} else {
			relations = L
		}

		cur = cur.parent
	}
}

func checkEqual(check *Node) bool {

	if check == nil {
		return true
	}

	lsize, rsize := getChildrenSize(check)
	if lsize > rsize && lsize-rsize < 3 {
		return true
	}

	if lsize < rsize && rsize-lsize < 3 {
		return true
	}

	return false
}

func (tree *ListTree) avlrrotate(cur *Node) {
	const R = 1
	llsize, lrsize := getChildrenSize(cur.children[R])
	if llsize > lrsize {
		tree.rrotate(cur.children[R])
	}
	tree.lrotate(cur)
}

func (tree *ListTree) avllrotate(cur *Node) {
	const L = 0
	llsize, lrsize := getChildrenSize(cur.children[L])
	if llsize < lrsize {
		tree.lrotate(cur.children[L])
	}
	tree.rrotate(cur)
}

func (tree *ListTree) lrotate(cur *Node) *Node {

	tree.Count++
	tree.RotateLog += " lrotate "

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
	tree.RotateLog += "rrotate"

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
