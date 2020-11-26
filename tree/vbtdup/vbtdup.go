package vbtdup

import (
	"github.com/davecgh/go-spew/spew"

	"github.com/474420502/focus/compare"
	"github.com/474420502/focus/tree"
)

type Node struct {
	children [2]*Node
	parent   *Node
	size     int
	value    interface{}
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}

	p := "nil"
	if n.parent != nil {
		p = spew.Sprint(n.parent.value)
	}
	return spew.Sprint(n.value) + "(" + p + "|" + spew.Sprint(n.size) + ")"
}

type Tree struct {
	root    *Node
	Compare compare.Compare

	iter *Iterator
}

func assertImplementation() {
	var _ tree.IBSTreeDup = (*Tree)(nil)
}

func New(Compare compare.Compare) *Tree {
	return &Tree{Compare: Compare, iter: NewIteratorWithCap(nil, 16)}
}

func (tree *Tree) String() string {
	str := "VBTree-Dup\n"
	if tree.root == nil {
		return str + "nil"
	}
	output(tree.root, "", true, &str)
	return str
}

func (tree *Tree) Iterator() *Iterator {
	return initIterator(tree)
}

func (tree *Tree) Size() int {
	if tree.root == nil {
		return 0
	}
	return tree.root.size
}

func (tree *Tree) IndexNode(idx int) *Node {
	cur := tree.root
	if idx >= 0 {
		for cur != nil {
			ls := getSize(cur.children[0])
			if idx == ls {
				return cur
			} else if idx < ls {
				cur = cur.children[0]
			} else {
				idx = idx - ls - 1
				cur = cur.children[1]
			}
		}
	} else {
		idx = -idx - 1
		for cur != nil {
			rs := getSize(cur.children[1])
			if idx == rs {
				return cur
			} else if idx < rs {
				cur = cur.children[1]
			} else {
				idx = idx - rs - 1
				cur = cur.children[0]
			}
		}
	}
	return nil
}

func (tree *Tree) Index(idx int) (interface{}, bool) {
	n := tree.IndexNode(idx)
	if n != nil {
		return n.value, true
	}
	return nil, false
}

func (tree *Tree) IndexRange(idx1, idx2 int) (result []interface{}, ok bool) { // 0 -1

	if idx1^idx2 < 0 {
		if idx1 < 0 {
			idx1 = tree.root.size + idx1
		} else {
			idx2 = tree.root.size + idx2
		}
	}

	if idx1 > idx2 {
		ok = true
		if idx1 >= tree.root.size {
			idx1 = tree.root.size - 1
			ok = false
		}

		n := tree.IndexNode(idx1)
		tree.iter.SetNode(n)
		iter := tree.iter
		result = make([]interface{}, 0, idx1-idx2)
		for i := idx2; i <= idx1; i++ {
			if iter.Prev() {
				result = append(result, iter.Value())
			} else {
				ok = false
				return
			}
		}

		return

	} else {
		ok = true
		if idx2 >= tree.root.size {
			idx2 = tree.root.size - 1
			ok = false
		}

		if n := tree.IndexNode(idx1); n != nil {
			tree.iter.SetNode(n)
			iter := tree.iter
			result = make([]interface{}, 0, idx2-idx1)
			for i := idx1; i <= idx2; i++ {
				if iter.Next() {
					result = append(result, iter.Value())
				} else {
					ok = false
					return
				}
			}

			return
		}

	}

	return nil, false
}

func (tree *Tree) RemoveIndex(idx int) (interface{}, bool) {
	n := tree.IndexNode(idx)
	if n != nil {
		tree.RemoveNode(n)
		return n.value, true
	}
	return nil, false
}

func (tree *Tree) RemoveNode(n *Node) {
	if tree.root.size == 1 {
		tree.root = nil
		// return n
		return
	}

	ls, rs := getChildrenSize(n)
	if ls == 0 && rs == 0 {
		p := n.parent
		p.children[getRelationship(n)] = nil
		tree.fixSizeWithRemove(p)
		// return n
		return
	}

	var cur *Node
	if ls > rs {
		cur = n.children[0]
		for cur.children[1] != nil {
			cur = cur.children[1]
		}

		cleft := cur.children[0]
		cur.parent.children[getRelationship(cur)] = cleft
		if cleft != nil {
			cleft.parent = cur.parent
		}

	} else {
		cur = n.children[1]
		for cur.children[0] != nil {
			cur = cur.children[0]
		}

		cright := cur.children[1]
		cur.parent.children[getRelationship(cur)] = cright

		if cright != nil {
			cright.parent = cur.parent
		}
	}

	cparent := cur.parent
	// 修改为interface 交换
	n.value, cur.value = cur.value, n.value

	// 考虑到刚好替换的节点是 被替换节点的孩子节点的时候, 从自身修复高度
	tree.fixSizeWithRemove(cparent)

	// return cur
	return
}

func (tree *Tree) Remove(key interface{}) (interface{}, bool) {

	if n, ok := tree.GetNode(key); ok {
		tree.RemoveNode(n)
		return n.value, true
	}
	// return nil
	return nil, false
}

func (tree *Tree) Clear() {
	tree.root = nil
	tree.iter = NewIteratorWithCap(nil, 16)
}

// Values 返回先序遍历的值
func (tree *Tree) Values() []interface{} {
	mszie := 0
	if tree.root != nil {
		mszie = tree.root.size
	}
	result := make([]interface{}, 0, mszie)
	tree.Traversal(func(v interface{}) bool {
		result = append(result, v)
		return true
	}, LDR)
	return result
}

func (tree *Tree) GetRange(k1, k2 interface{}) (result []interface{}) {
	c := tree.Compare(k2, k1)
	switch c {
	case 1:

		var min, max *Node
		resultmin := tree.getArountNode(k1)
		resultmax := tree.getArountNode(k2)
		for i := 1; i < 3 && min == nil; i++ {
			min = resultmin[i]
		}

		for i := 1; i > -1 && max == nil; i-- {
			max = resultmax[i]
		}

		if max == nil {
			return []interface{}{}
		}

		result = make([]interface{}, 0, 16)

		tree.iter.SetNode(min)
		iter := tree.iter
		for iter.Next() {
			result = append(result, iter.Value())
			if iter.cur == max {
				break
			}
		}
	case -1:

		var min, max *Node
		resultmin := tree.getArountNode(k2)
		resultmax := tree.getArountNode(k1)
		for i := 1; i < 3 && min == nil; i++ {
			min = resultmin[i]
		}
		for i := 1; i > -1 && max == nil; i-- {
			max = resultmax[i]
		}

		if min == nil {
			return []interface{}{}
		}

		result = make([]interface{}, 0, 16)

		tree.iter.SetNode(max)
		iter := tree.iter
		for iter.Prev() {
			result = append(result, iter.Value())
			if iter.cur == min {
				break
			}
		}
	case 0:
		if n, ok := tree.GetNode(k1); ok {
			return []interface{}{n.value}
		}
		return []interface{}{}
	}

	return
}

func (tree *Tree) Get(key interface{}) (interface{}, bool) {
	n, ok := tree.GetNode(key)
	if ok {
		return n.value, true
	}
	return n, false
}

func (tree *Tree) GetAround(key interface{}) (result [3]interface{}) {
	an := tree.getArountNode(key)
	for i, n := range an {
		if n != nil {
			result[i] = n.value
		}
	}
	return
}

func (tree *Tree) getArountNode(key interface{}) (result [3]*Node) {
	var last *Node
	var lastc int

	for n := tree.root; n != nil; {
		last = n
		c := tree.Compare(key, n.value)
		switch c {
		case -1:
			n = n.children[0]
			lastc = c
		case 1:
			n = n.children[1]
			lastc = c
		case 0:
			tree.iter.SetNode(n)
			iter := tree.iter
			iter.Prev()
			for iter.Prev() {
				if tree.Compare(iter.cur.value, n.value) == 0 {
					n = iter.cur
				} else {
					break
				}
			}
			result[1] = n
			n = nil
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}
	}

	switch lastc {
	case 1:

		if result[1] != nil {

			result[0] = tree.iter.GetPrev(result[1], 1)
			result[2] = tree.iter.GetNext(result[1], 1)
		} else {
			result[0] = last
			result[2] = tree.iter.GetNext(last, 1)
		}

	case -1:

		if result[1] != nil {
			result[0] = tree.iter.GetPrev(result[1], 1)
			result[2] = tree.iter.GetNext(result[1], 1)
		} else {
			result[2] = last
			result[0] = tree.iter.GetPrev(last, 1)
		}

	case 0:

		if result[1] == nil {
			return
		}
		result[0] = tree.iter.GetPrev(result[1], 1)
		result[2] = tree.iter.GetNext(result[1], 1)
	}
	return
}

func (tree *Tree) GetNode(key interface{}) (*Node, bool) {

	for n := tree.root; n != nil; {
		switch c := tree.Compare(key, n.value); c {
		case -1:
			n = n.children[0]
		case 1:
			n = n.children[1]
		case 0:
			return n, true
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}
	}
	return nil, false
}

// Put return bool
func (tree *Tree) Put(value interface{}) (isInsert bool) {

	node := &Node{value: value, size: 1}
	if tree.root == nil {
		tree.root = node
		return true
	}

	for cur := tree.root; ; {

		if cur.size > 8 {
			factor := cur.size / 10 // or factor = 1
			ls, rs := getChildrenSize(cur)
			if rs >= ls*2+factor || ls >= rs*2+factor {
				tree.fixSize(cur, ls, rs)
			}
		}

		c := tree.Compare(value, cur.value)
		switch {
		case c < 0:
			if cur.children[0] == nil {
				cur.children[0] = node
				node.parent = cur

				for temp := cur; temp != nil; temp = temp.parent {
					temp.size++
				}

				if cur.parent != nil && cur.parent.size == 3 {
					if cur.parent.children[0] == nil {
						tree.lrrotate3(cur.parent)
					} else {
						tree.rrotate3(cur.parent)
					}
				}
				return true
			}
			cur = cur.children[0]
		case c > 0:
			if cur.children[1] == nil {
				cur.children[1] = node
				node.parent = cur

				for temp := cur; temp != nil; temp = temp.parent {
					temp.size++
				}

				if cur.parent != nil && cur.parent.size == 3 {
					if cur.parent.children[1] == nil {
						tree.rlrotate3(cur.parent)
					} else {
						tree.lrotate3(cur.parent)
					}
				}
				return true
			}
			cur = cur.children[1]
		default:
			cur.value = value
			return false
		}

	}
}

// TraversalMethod 遍历模式
type TraversalMethod int

const (
	// L = left R = right D = Value(dest)
	_ TraversalMethod = iota
	//DLR 先值 然后左递归 右递归 下面同理
	DLR
	//LDR 先从左边有序访问到右边 从小到大
	LDR
	// LRD 同理
	LRD

	// DRL 同理
	DRL

	// RDL 先从右边有序访问到左边 从大到小
	RDL

	// RLD 同理
	RLD
)

// Traversal 遍历的方法 默认是LDR 从小到大 Compare 为 l < r
func (tree *Tree) Traversal(every func(v interface{}) bool, traversalMethod ...interface{}) {
	if tree.root == nil {
		return
	}

	method := LDR
	if len(traversalMethod) != 0 {
		method = traversalMethod[0].(TraversalMethod)
	}

	switch method {
	case DLR:
		var traverasl func(cur *Node) bool
		traverasl = func(cur *Node) bool {
			if cur == nil {
				return true
			}
			if !every(cur.value) {
				return false
			}
			if !traverasl(cur.children[0]) {
				return false
			}
			if !traverasl(cur.children[1]) {
				return false
			}
			return true
		}
		traverasl(tree.root)
	case LDR:
		var traverasl func(cur *Node) bool
		traverasl = func(cur *Node) bool {
			if cur == nil {
				return true
			}
			if !traverasl(cur.children[0]) {
				return false
			}
			if !every(cur.value) {
				return false
			}
			if !traverasl(cur.children[1]) {
				return false
			}
			return true
		}
		traverasl(tree.root)
	case LRD:
		var traverasl func(cur *Node) bool
		traverasl = func(cur *Node) bool {
			if cur == nil {
				return true
			}
			if !traverasl(cur.children[0]) {
				return false
			}
			if !traverasl(cur.children[1]) {
				return false
			}
			if !every(cur.value) {
				return false
			}
			return true
		}
		traverasl(tree.root)
	case DRL:
		var traverasl func(cur *Node) bool
		traverasl = func(cur *Node) bool {
			if cur == nil {
				return true
			}
			if !every(cur.value) {
				return false
			}
			if !traverasl(cur.children[1]) {
				return false
			}
			if !traverasl(cur.children[0]) {
				return false
			}
			return true
		}
		traverasl(tree.root)
	case RDL:
		var traverasl func(cur *Node) bool
		traverasl = func(cur *Node) bool {
			if cur == nil {
				return true
			}
			if !traverasl(cur.children[1]) {
				return false
			}
			if !every(cur.value) {
				return false
			}
			if !traverasl(cur.children[0]) {
				return false
			}
			return true
		}
		traverasl(tree.root)
	case RLD:
		var traverasl func(cur *Node) bool
		traverasl = func(cur *Node) bool {
			if cur == nil {
				return true
			}
			if !traverasl(cur.children[1]) {
				return false
			}
			if !traverasl(cur.children[0]) {
				return false
			}
			if !every(cur.value) {
				return false
			}
			return true
		}
		traverasl(tree.root)
	}
}

func (tree *Tree) lrrotate3(cur *Node) {
	const l = 1
	const r = 0

	movparent := cur.children[l]
	mov := movparent.children[r]

	mov.value, cur.value = cur.value, mov.value //交换值达到, 相对位移

	cur.children[r] = mov
	mov.parent = cur

	cur.children[l] = movparent
	movparent.children[r] = nil

	cur.children[r] = mov
	mov.parent = cur

	// cur.size = 3
	// cur.children[r].size = 1
	cur.children[l].size = 1
}

func (tree *Tree) lrrotate(cur *Node) {

	const l = 1
	const r = 0

	movparent := cur.children[l]
	mov := movparent.children[r]

	mov.value, cur.value = cur.value, mov.value //交换值达到, 相对位移

	if mov.children[l] != nil {
		movparent.children[r] = mov.children[l]
		movparent.children[r].parent = movparent
		//movparent.children[r].child = l
	} else {
		movparent.children[r] = nil
	}

	if mov.children[r] != nil {
		mov.children[l] = mov.children[r]
		//mov.children[l].child = l
	} else {
		mov.children[l] = nil
	}

	if cur.children[r] != nil {
		mov.children[r] = cur.children[r]
		mov.children[r].parent = mov
	} else {
		mov.children[r] = nil
	}

	cur.children[r] = mov
	mov.parent = cur

	movparent.size = getChildrenSumSize(movparent) + 1
	mov.size = getChildrenSumSize(mov) + 1
	cur.size = getChildrenSumSize(cur) + 1
}

func (tree *Tree) rlrotate3(cur *Node) {
	const l = 0
	const r = 1

	movparent := cur.children[l]
	mov := movparent.children[r]

	mov.value, cur.value = cur.value, mov.value //交换值达到, 相对位移

	cur.children[r] = mov
	mov.parent = cur

	cur.children[l] = movparent
	movparent.children[r] = nil

	cur.children[r] = mov
	mov.parent = cur

	// cur.size = 3
	// cur.children[r].size = 1
	cur.children[l].size = 1
}

func (tree *Tree) rlrotate(cur *Node) {

	const l = 0
	const r = 1

	movparent := cur.children[l]
	mov := movparent.children[r]

	mov.value, cur.value = cur.value, mov.value //交换值达到, 相对位移

	if mov.children[l] != nil {
		movparent.children[r] = mov.children[l]
		movparent.children[r].parent = movparent
	} else {
		movparent.children[r] = nil
	}

	if mov.children[r] != nil {
		mov.children[l] = mov.children[r]
	} else {
		mov.children[l] = nil
	}

	if cur.children[r] != nil {
		mov.children[r] = cur.children[r]
		mov.children[r].parent = mov
	} else {
		mov.children[r] = nil
	}

	cur.children[r] = mov
	mov.parent = cur

	movparent.size = getChildrenSumSize(movparent) + 1
	mov.size = getChildrenSumSize(mov) + 1
	cur.size = getChildrenSumSize(cur) + 1
}

func (tree *Tree) rrotate3(cur *Node) {
	const l = 0
	const r = 1
	// 1 right 0 left
	mov := cur.children[l]

	mov.value, cur.value = cur.value, mov.value //交换值达到, 相对位移

	cur.children[r] = mov

	cur.children[l] = mov.children[l]
	cur.children[l].parent = cur

	mov.children[l] = nil

	mov.size = 1
}

func (tree *Tree) rrotate(cur *Node) {

	const l = 0
	const r = 1
	// 1 right 0 left
	mov := cur.children[l]

	mov.value, cur.value = cur.value, mov.value //交换值达到, 相对位移

	//  mov.children[l]不可能为nil
	mov.children[l].parent = cur

	cur.children[l] = mov.children[l]

	// 解决mov节点孩子转移的问题
	if mov.children[r] != nil {
		mov.children[l] = mov.children[r]
	} else {
		mov.children[l] = nil
	}

	if cur.children[r] != nil {
		mov.children[r] = cur.children[r]
		mov.children[r].parent = mov
	} else {
		mov.children[r] = nil
	}

	// 连接转移后的节点 由于mov只是与cur交换值,parent不变
	cur.children[r] = mov

	mov.size = getChildrenSumSize(mov) + 1
	cur.size = getChildrenSumSize(cur) + 1
}

func (tree *Tree) lrotate3(cur *Node) {
	const l = 1
	const r = 0
	// 1 right 0 left
	mov := cur.children[l]

	mov.value, cur.value = cur.value, mov.value //交换值达到, 相对位移

	cur.children[r] = mov

	cur.children[l] = mov.children[l]
	cur.children[l].parent = cur

	mov.children[l] = nil

	mov.size = 1
}

func (tree *Tree) lrotate(cur *Node) {

	const l = 1
	const r = 0
	// 1 right 0 left
	mov := cur.children[l]

	mov.value, cur.value = cur.value, mov.value //交换值达到, 相对位移

	//  mov.children[l]不可能为nil
	mov.children[l].parent = cur

	cur.children[l] = mov.children[l]

	// 解决mov节点孩子转移的问题
	if mov.children[r] != nil {
		mov.children[l] = mov.children[r]
	} else {
		mov.children[l] = nil
	}

	if cur.children[r] != nil {
		mov.children[r] = cur.children[r]
		mov.children[r].parent = mov
	} else {
		mov.children[r] = nil
	}

	// 连接转移后的节点 由于mov只是与cur交换值,parent不变
	cur.children[r] = mov

	mov.size = getChildrenSumSize(mov) + 1
	cur.size = getChildrenSumSize(cur) + 1
}

func getChildrenSumSize(cur *Node) int {
	return getSize(cur.children[0]) + getSize(cur.children[1])
}

func getChildrenSize(cur *Node) (int, int) {
	return getSize(cur.children[0]), getSize(cur.children[1])
}

func getSize(cur *Node) int {
	if cur == nil {
		return 0
	}
	return cur.size
}

func (tree *Tree) fixSizeWithRemove(cur *Node) {
	for cur != nil {
		cur.size--
		if cur.size > 8 {
			factor := cur.size / 10 // or factor = 1
			ls, rs := getChildrenSize(cur)
			if rs >= ls*2+factor || ls >= rs*2+factor {
				tree.fixSize(cur, ls, rs)
			}
		} else if cur.size == 3 {
			if cur.children[0] == nil {
				if cur.children[1].children[0] == nil {
					tree.lrotate3(cur)
				} else {
					tree.lrrotate3(cur)
				}
			} else if cur.children[1] == nil {
				if cur.children[0].children[1] == nil {
					tree.rrotate3(cur)
				} else {
					tree.rlrotate3(cur)
				}
			}
		}
		cur = cur.parent
	}
}

func (tree *Tree) fixSize(cur *Node, ls, rs int) {
	if ls > rs {
		llsize, lrsize := getChildrenSize(cur.children[0])
		if lrsize > llsize {
			tree.rlrotate(cur)
		} else {
			tree.rrotate(cur)
		}
	} else {
		rlsize, rrsize := getChildrenSize(cur.children[1])
		if rlsize > rrsize {
			tree.lrrotate(cur)
		} else {
			tree.lrotate(cur)
		}
	}
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

	*str += spew.Sprint(node.value) + ":" + spew.Sprint(node.value) + "\n"

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
		parentv = spew.Sprint(node.parent.value)
	}
	suffix += parentv + "|" + spew.Sprint(node.size) + ")"
	*str += spew.Sprint(node.value) + suffix + "\n"

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

func (tree *Tree) debugString() string {
	str := "VBTree-Dup\n"
	if tree.root == nil {
		return str + "nil"
	}
	outputfordebug(tree.root, "", true, &str)
	return str
}
