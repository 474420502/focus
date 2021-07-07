package listtree

type Node struct {
	parent   *Node
	children [2]*Node
	direct   [2]*Node

	size  int64
	key   []byte
	value []byte
}

type TreeList struct {
	root    *Node
	compare Compare

	Count int64
}

func New() *TreeList {
	return &TreeList{compare: CompatorMath, root: &Node{}}
}

func (tree *TreeList) Put(key, value []byte) bool {

	cur := tree.root.children[0]
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

				tree.fix(cur.parent, [2]int{0, L})
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

func (tree *TreeList) fix(cur *Node, relations [2]int) {

	// s := cur
	// defer func() {
	// 	if err := recover(); err != nil {

	// 		var temp []byte = cur.key
	// 		var temp2 []byte = s.key
	// 		cur.key = []byte(fmt.Sprintf("\033[35m%s\033[0m", cur.key))
	// 		s.key = []byte(fmt.Sprintf("\033[32m%s\033[0m", s.key))
	// 		log.Println(tree.debugString())

	// 		cur.key = temp
	// 		s.key = temp2
	// 		log.Panic(err)
	// 	}
	// }()

	const L = 0
	const R = 1

	cur.size++

	if cur.parent.children[L] == cur {
		if relations[1] != L {
			relations[0] = R
		} else {
			relations[0] = L
		}
	} else {
		if relations[1] != R {
			relations[0] = L
		} else {
			relations[0] = R
		}
	}
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
						tree.rrotate(cur.children[R].children[L])

					}
					cur = tree.lrotate(cur)

				}

			} else {

				rsize := getSize(cur.children[R])
				if rsize <= childLimitSize {

					if relations[1] == R {
						tree.lrotate(cur.children[L].children[R])
					}
					// if cur.children[L] == nil {
					// 	log.Println(tree.root.children[0].size)
					// 	str := "BinarayList\n"

					// 	outputfordebug(cur, "", true, &str)
					// 	log.Println(str)
					// 	// log.Println(tree.debugString())
					// }
					cur = tree.rrotate(cur)

				}

			}

		} else {
			height++
			childLimitSize = limitsize
		}

		relations[1] = relations[0]
		if cur.parent.children[L] == cur {
			if relations[1] != L {
				relations[0] = R
			} else {
				relations[0] = L
			}
		} else {
			if relations[1] != R {
				relations[0] = L
			} else {
				relations[0] = R
			}
		}
		cur = cur.parent

	}
}

func (tree *TreeList) lrotate(cur *Node) *Node {

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

func (tree *TreeList) rrotate(cur *Node) *Node {

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
