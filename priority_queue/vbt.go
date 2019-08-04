package pqueue

import (
	"github.com/davecgh/go-spew/spew"

	"github.com/474420502/focus/compare"
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

type vbTree struct {
	root    *Node
	Compare compare.Compare

	top *Node

	iter *Iterator
}

func newVBT(Compare compare.Compare) *vbTree {
	return &vbTree{Compare: Compare, iter: NewIteratorWithCap(nil, 16)}
}

func (tree *vbTree) String() string {
	str := "AVLTree\n"
	if tree.root == nil {
		return str + "nil"
	}
	output(tree.root, "", true, &str)
	return str
}

func (tree *vbTree) Iterator() *Iterator {
	return initIterator(tree)
}

func (tree *vbTree) Size() int {
	if tree.root == nil {
		return 0
	}
	return tree.root.size
}

func (tree *vbTree) indexNode(idx int) *Node {
	cur := tree.root
	if idx >= 0 {
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
	} else {
		idx = -idx - 1
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
	}
	return nil
}

func (tree *vbTree) Index(idx int) (interface{}, bool) {
	n := tree.indexNode(idx)
	if n != nil {
		return n.value, true
	}
	return nil, false
}

func (tree *vbTree) IndexRange(idx1, idx2 int) (result []interface{}, ok bool) { // 0 -1

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

		n := tree.indexNode(idx1)
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

		if n := tree.indexNode(idx1); n != nil {
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

func (tree *vbTree) RemoveIndex(idx int) (interface{}, bool) {
	n := tree.indexNode(idx)
	if n != nil {
		tree.removeNode(n)
		return n.value, true
	}
	return nil, false
}

func (tree *vbTree) removeNode(n *Node) {
	if tree.root.size == 1 {
		tree.root = nil
		tree.top = nil
		// return n
		return
	}

	if tree.top == n {
		tree.top = tree.iter.GetNext(n, 1)
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
	// n.value, cur.value = cur.value, n.value
	tree.replace(n, cur)

	// 考虑到刚好替换的节点是 被替换节点的孩子节点的时候, 从自身修复高度
	if cparent == n {
		tree.fixSizeWithRemove(cur)
	} else {
		tree.fixSizeWithRemove(cparent)
	}

	// return cur
	return
}

func (tree *vbTree) Remove(key interface{}) (interface{}, bool) {

	if n, ok := tree.GetNode(key); ok {
		tree.removeNode(n)
		return n.value, true
	}
	// return nil
	return nil, false
}

// Values 返回先序遍历的值
func (tree *vbTree) Values() []interface{} {
	mszie := 0
	if tree.root != nil {
		mszie = tree.root.size
	}
	result := make([]interface{}, 0, mszie)
	tree.Traversal(func(v interface{}) bool {
		result = append(result, v)
		return true
	}, RDL)
	return result
}

func (tree *vbTree) GetRange(k1, k2 interface{}) (result []interface{}) {
	c := tree.Compare(k2, k1)
	switch c {
	case 1:

		var min, max *Node
		resultmin := tree.getArounNode(k1)
		resultmax := tree.getArounNode(k2)
		for i := 1; i < 3 && min == nil; i++ {
			min = resultmin[i]
		}

		for i := 1; i > -1 && max == nil; i-- {
			max = resultmax[i]
		}

		if max == nil {
			return []interface{}{}
		}

		result = make([]interface{}, 0, 8)

		// iter := NewIterator(min)
		tree.iter.SetNode(min)
		iter := tree.iter
		for iter.Prev() {
			result = append(result, iter.Value())
			if iter.cur == max {
				break
			}
		}
	case -1:

		var min, max *Node
		resultmin := tree.getArounNode(k2)
		resultmax := tree.getArounNode(k1)
		for i := 1; i < 3 && min == nil; i++ {
			min = resultmin[i]
		}
		for i := 1; i > -1 && max == nil; i-- {
			max = resultmax[i]
		}

		if min == nil {
			return []interface{}{}
		}

		result = make([]interface{}, 0, 8)

		// iter := NewIterator(max)
		tree.iter.SetNode(max)
		iter := tree.iter
		for iter.Next() {
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

func (tree *vbTree) Get(key interface{}) (interface{}, bool) {
	n, ok := tree.GetNode(key)
	if ok {
		return n.value, true
	}
	return n, false
}

// GetAround 改成Big To Small
func (tree *vbTree) GetAround(key interface{}) (result [3]interface{}) {
	an := tree.getArounNode(key)
	for i, n := range an {
		if n != nil {
			result[2-i] = n.value
		}
	}
	return
}

func (tree *vbTree) getArounNode(key interface{}) (result [3]*Node) {
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
			iter.Next()
			for iter.Next() {
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

			result[0] = tree.iter.GetNext(result[1], 1)
			result[2] = tree.iter.GetPrev(result[1], 1)
		} else {
			result[0] = last
			result[2] = tree.iter.GetPrev(last, 1)
		}

	case -1:

		if result[1] != nil {
			result[0] = tree.iter.GetNext(result[1], 1)
			result[2] = tree.iter.GetPrev(result[1], 1)
		} else {
			result[2] = last
			result[0] = tree.iter.GetNext(last, 1)
		}

	case 0:

		if result[1] == nil {
			return
		}
		result[0] = tree.iter.GetNext(result[1], 1)
		result[2] = tree.iter.GetPrev(result[1], 1)
	}
	return
}

func (tree *vbTree) GetNode(value interface{}) (*Node, bool) {

	for n := tree.root; n != nil; {
		switch c := tree.Compare(value, n.value); c {
		case -1:
			n = n.children[0]
		case 1:
			n = n.children[1]
		case 0:

			tree.iter.SetNode(n)
			iter := tree.iter
			iter.Next()
			for iter.Next() {
				if tree.Compare(iter.cur.value, n.value) == 0 {
					n = iter.cur
				} else {
					break
				}
			}
			return n, true
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}
	}
	return nil, false
}

func (tree *vbTree) Put(key interface{}) {

	Node := &Node{value: key, size: 1}
	if tree.root == nil {
		tree.root = Node
		tree.top = Node
		return
	}

	if tree.Compare(key, tree.top.value) > 0 {
		tree.top = Node
	}

	for cur := tree.root; ; {

		if cur.size > 8 {
			factor := cur.size / 10 // or factor = 1
			ls, rs := cur.children[0].size, cur.children[1].size
			if rs >= ls*2+factor || ls >= rs*2+factor {
				cur = tree.fixSize(cur, ls, rs)
			}
		}

		cur.size++
		c := tree.Compare(key, cur.value)
		if c < 0 {
			if cur.children[0] == nil {
				cur.children[0] = Node
				Node.parent = cur

				if cur.parent != nil && cur.parent.size == 3 {
					if cur.parent.children[0] == nil {
						tree.lrrotate3(cur.parent)
					} else {
						tree.rrotate3(cur.parent)
					}
				}

				return
			}
			cur = cur.children[0]
		} else {
			if cur.children[1] == nil {
				cur.children[1] = Node
				Node.parent = cur

				if cur.parent != nil && cur.parent.size == 3 {
					if cur.parent.children[1] == nil {
						tree.rlrotate3(cur.parent)
					} else {
						tree.lrotate3(cur.parent)
					}
				}
				return
			}
			cur = cur.children[1]
		}
	}
}

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
func (tree *vbTree) Traversal(every func(v interface{}) bool, traversalMethod ...interface{}) {
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
			if !traverasl(cur.children[0]) {
				return false
			}
			if !traverasl(cur.children[1]) {
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

func setChildNotNil(cur *Node, cidx int, child *Node) {
	cur.children[cidx] = child
	cur.children[cidx].parent = cur
}

func setChild(cur *Node, cidx int, child *Node) {
	cur.children[cidx] = child
	if child != nil {
		cur.children[cidx].parent = cur
	}
}

func (tree *vbTree) replace(old, new *Node) {

	setChild(new, 0, old.children[0])
	setChild(new, 1, old.children[1])

	if old.parent == nil {
		tree.root = new
	} else {
		if old.parent.children[1] == old {
			old.parent.children[1] = new
		} else {
			old.parent.children[0] = new
		}
	}
	new.size = old.size
	new.parent = old.parent
}

func (tree *vbTree) takeParent(token, person *Node) {
	if token.parent == nil {
		tree.root = person
	} else {
		if token.parent.children[1] == token {
			token.parent.children[1] = person
		} else {
			token.parent.children[0] = person
		}
	}
	person.parent = token.parent
}

func (tree *vbTree) lrrotate3(cur *Node) *Node {
	const l = 1
	const r = 0

	ln := cur.children[l]
	cur.children[l] = nil

	lrn := ln.children[r]
	ln.children[r] = nil

	tree.takeParent(cur, lrn)
	setChildNotNil(lrn, l, ln)
	setChildNotNil(lrn, r, cur)

	lrn.size = 3
	lrn.children[l].size = 1
	lrn.children[r].size = 1
	return lrn
}

func (tree *vbTree) lrrotate(cur *Node) *Node {

	const l = 1
	const r = 0

	ln := cur.children[l]
	lrn := ln.children[r]

	lrln := lrn.children[l]
	lrrn := lrn.children[r]

	tree.takeParent(cur, lrn)

	setChild(ln, r, lrln)
	setChild(cur, l, lrrn)

	setChildNotNil(lrn, l, ln)
	setChildNotNil(lrn, r, cur)

	ln.size = getChildrenSumSize(ln) + 1
	cur.size = getChildrenSumSize(cur) + 1
	lrn.size = getChildrenSumSize(lrn) + 1

	return lrn
}

func (tree *vbTree) rlrotate3(cur *Node) *Node {
	const l = 0
	const r = 1

	ln := cur.children[l]
	cur.children[l] = nil

	lrn := ln.children[r]
	ln.children[r] = nil

	tree.takeParent(cur, lrn)
	setChildNotNil(lrn, l, ln)
	setChildNotNil(lrn, r, cur)

	lrn.size = 3
	lrn.children[l].size = 1
	lrn.children[r].size = 1
	return lrn
}

func (tree *vbTree) rlrotate(cur *Node) *Node {

	const l = 0
	const r = 1

	ln := cur.children[l]
	lrn := ln.children[r]

	lrln := lrn.children[l]
	lrrn := lrn.children[r]

	tree.takeParent(cur, lrn)

	setChild(ln, r, lrln)
	setChild(cur, l, lrrn)

	setChildNotNil(lrn, l, ln)
	setChildNotNil(lrn, r, cur)

	ln.size = getChildrenSumSize(ln) + 1
	cur.size = getChildrenSumSize(cur) + 1
	lrn.size = getChildrenSumSize(lrn) + 1

	return lrn
}

func (tree *vbTree) rrotate3(cur *Node) *Node {
	const l = 0
	const r = 1
	// 1 right 0 left
	mov := cur.children[l]
	cur.children[l] = nil

	tree.takeParent(cur, mov)
	setChildNotNil(mov, r, cur)

	mov.size = 3
	cur.size = 1
	return mov
}

func (tree *vbTree) rrotate(cur *Node) *Node {
	const l = 0
	const r = 1
	// 1 right 0 left
	ln := cur.children[l]
	lrn := ln.children[r]

	tree.takeParent(cur, ln)
	setChild(cur, l, lrn)
	setChildNotNil(ln, r, cur)

	cur.size = getChildrenSumSize(cur) + 1
	ln.size = getChildrenSumSize(ln) + 1

	return ln
}

func (tree *vbTree) lrotate3(cur *Node) *Node {
	const l = 1
	const r = 0

	// 1 right 0 left
	mov := cur.children[l]
	cur.children[l] = nil

	tree.takeParent(cur, mov)
	setChildNotNil(mov, r, cur)

	mov.size = 3
	cur.size = 1
	return mov
}

func (tree *vbTree) lrotate(cur *Node) *Node {

	const l = 1
	const r = 0

	// 1 right 0 left
	ln := cur.children[l]
	lrn := ln.children[r]

	tree.takeParent(cur, ln)
	setChild(cur, l, lrn)
	setChildNotNil(ln, r, cur)

	cur.size = getChildrenSumSize(cur) + 1
	ln.size = getChildrenSumSize(ln) + 1

	return ln
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

func (tree *vbTree) fixSizeWithRemove(cur *Node) {
	for cur != nil {
		cur.size--
		if cur.size > 8 {
			factor := cur.size / 10 // or factor = 1
			ls, rs := getChildrenSize(cur)
			if rs >= ls*2+factor || ls >= rs*2+factor {
				cur = tree.fixSize(cur, ls, rs)
			}
		} else if cur.size == 3 {
			if cur.children[0] == nil {
				if cur.children[1].children[0] == nil {
					cur = tree.lrotate3(cur)
				} else {
					cur = tree.lrrotate3(cur)
				}
			} else if cur.children[1] == nil {
				if cur.children[0].children[1] == nil {
					cur = tree.rrotate3(cur)
				} else {
					cur = tree.rlrotate3(cur)
				}
			}
		}
		cur = cur.parent
	}
}

func (tree *vbTree) fixSize(cur *Node, ls, rs int) *Node {
	if ls > rs {
		llsize, lrsize := getChildrenSize(cur.children[0])
		if lrsize > llsize {
			return tree.rlrotate(cur)
		}
		return tree.rrotate(cur)

	} else {
		rlsize, rrsize := getChildrenSize(cur.children[1])
		if rlsize > rrsize {
			return tree.lrrotate(cur)
		}
		return tree.lrotate(cur)
	}
}

func output(Node *Node, prefix string, isTail bool, str *string) {

	if Node.children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(Node.children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}

	*str += spew.Sprint(Node.value) + "\n"

	if Node.children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(Node.children[0], newPrefix, true, str)
	}

}

func outputfordebug(Node *Node, prefix string, isTail bool, str *string) {

	if Node.children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		outputfordebug(Node.children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}

	suffix := "("
	parentv := ""
	if Node.parent == nil {
		parentv = "nil"
	} else {
		parentv = spew.Sprint(Node.parent.value)
	}
	suffix += parentv + "|" + spew.Sprint(Node.size) + ")"
	*str += spew.Sprint(Node.value) + suffix + "\n"

	if Node.children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		outputfordebug(Node.children[0], newPrefix, true, str)
	}
}

func (tree *vbTree) debugString() string {
	str := "AVLTree\n"
	if tree.root == nil {
		return str + "nil"
	}
	outputfordebug(tree.root, "", true, &str)
	return str
}
