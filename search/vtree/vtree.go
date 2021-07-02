package vtree

import (
	linkedlist "github.com/474420502/focus/list/linked_list"
	"github.com/davecgh/go-spew/spew"
)

// Node tree的节点
type Node struct {
	children [2]*Node
	parent   *Node
	relation byte
	size     int
	key      []byte
	value    []byte
}

// IteratorBase return iterator and start by node
func (n *Node) IteratorBase(tree *Tree) *IteratorBase {
	return NewIteratorBase(tree, n)
}

// IteratorRange return iterator and start by node
func (n *Node) IteratorRange(tree *Tree) *IteratorRange {
	return NewIteratorRange(tree, n)
}

// IteratorPrefix return iterator and start by node
func (n *Node) IteratorPrefix(tree *Tree) *IteratorPrefix {
	return NewIteratorPrefix(tree, n)
}

// Key get node key
func (n *Node) Key() []byte {
	return n.key
}

// Value get node value
func (n *Node) Value() []byte {
	return n.value
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}
	return spew.Sprint(string(n.key)) + ":" + spew.Sprint(string(n.value))
}

func (n *Node) debugString() string {
	if n == nil {
		return "nil"
	}

	p := "nil"
	if n.parent != nil {
		p = spew.Sprint(string(n.parent.value))
	}
	return spew.Sprint(string(n.value)) + "(" + p + "|" + spew.Sprint(n.size) + ")"
}

// Tree increasing
type Tree struct {
	compartor func([]byte, []byte) int
	root      *Node
	// iter *Iterator
}

func assertImplementation() {

}

// NewWithCompartor Create a vtree with diffent compare func
func NewWithCompartor(compartor func([]byte, []byte) int) *Tree {
	return &Tree{compartor: compartor}
}

// New Create a vtree
func New() *Tree {
	return &Tree{compartor: CompatorMath}
}

func (tree *Tree) String() string {
	str := "VTree-Key\n"
	if tree.root == nil {
		return str + "nil"
	}
	output(tree.root, "", true, &str)
	return str
}

// Size get tree size
func (tree *Tree) Size() int {
	if tree.root == nil {
		return 0
	}
	return tree.root.size
}

// func (tree *Tree) seekRangeEx(start, end []byte) {
// 	// sn := tree.seekNodeNext(start)
// 	en := tree.seekNodePrev(end)

// 	var sn, lastn *Node
// 	// lastn = tree.root
// 	for sn = tree.root; sn != nil; {
// 		c := tree.compartor(start, sn.key)
// 		switch c {
// 		case -1:
// 			lastn = sn
// 			sn = sn.children[0]
// 		case 1:
// 			// lastn = n
// 			sn = sn.children[1]
// 		case 0:
// 			break
// 		default:
// 			panic("Get Compare only is allowed in -1, 0, 1")
// 		}
// 	}

// 	if sn == nil {
// 		sn = lastn
// 	}

// }

// SeekRange start end
func (tree *Tree) SeekRange(start, end []byte) Iterator {
	node := tree.seekNodeNext(start)

	var iter *IteratorRange
	if node != nil {
		iter = node.IteratorRange(tree)
		if tree.compartor(start, end) == 1 {
			iter.SetLimit(end, start)
		} else {
			iter.SetLimit(start, end)
		}
	}
	return iter
}

// SeekRangeString start end
func (tree *Tree) SeekRangeString(start, end string) Iterator {
	return tree.SeekRange([]byte(start), []byte(end))
}

// SeekPrefix prefix range
func (tree *Tree) SeekPrefix(prefix []byte) Iterator {
	if node := tree.seekNodeNext(prefix); node != nil {
		iter := node.IteratorPrefix(tree)
		iter.SetLimit(prefix)
		return iter
	}
	return nil
}

// SeekPrefixString prefix range
func (tree *Tree) SeekPrefixString(prefix string) Iterator {
	return tree.SeekPrefix([]byte(prefix))
}

// Seek search key . like rocksdb/leveldb api
func (tree *Tree) seekNodePrev(key []byte) *Node {
	// lastc := 0
	var n, lastn *Node
	// lastn = tree.root
	for n = tree.root; n != nil; {

		c := tree.compartor(key, n.key)
		switch c {
		case -1:
			// lastn = n
			n = n.children[0]
		case 1:
			lastn = n
			n = n.children[1]
		case 0:
			return n
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}

	}

	return lastn
}

// Seek search key . like rocksdb/leveldb api
func (tree *Tree) seekNodeNext(key []byte) *Node {
	// lastc := 0
	// var n, switchParent, lastn *Node
	var n, lastn *Node
	// lastn = tree.root
	for n = tree.root; n != nil; {

		c := tree.compartor(key, n.key)
		switch c {
		case -1:
			lastn = n
			n = n.children[0]
		case 1:
			// lastn = n
			n = n.children[1]
		case 0:
			return n
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}

	}
	return lastn
}

// Seek search key . like rocksdb/leveldb api
func (tree *Tree) seekNodePrevEx(key []byte) *Node {
	var n, lastleft, lastright *Node
	// lastn = tree.root
	for n = tree.root; n != nil; {

		c := tree.compartor(key, n.key)
		switch c {
		case -1:
			lastleft = n
			n = n.children[0]
		case 1:
			lastright = n
			n = n.children[1]
		case 0:
			return n
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}
	}

	if lastright == nil {
		return lastleft
	}
	return lastright
}

// Seek search key . like rocksdb/leveldb api
func (tree *Tree) seekNodeNextEx(key []byte) *Node {
	var n, lastleft, lastright *Node
	// lastn = tree.root
	for n = tree.root; n != nil; {

		c := tree.compartor(key, n.key)
		switch c {
		case -1:
			lastleft = n
			n = n.children[0]
		case 1:
			lastright = n
			n = n.children[1]
		case 0:
			return n
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}
	}

	if lastleft == nil {
		return lastright
	}
	return lastleft
}

// Seek search key . like rocksdb/leveldb api
func (tree *Tree) Seek(key []byte) Iterator {
	if node := tree.seekNodeNextEx(key); node != nil {
		return node.IteratorBase(tree)
	}
	return nil
}

// SeekString search key(string) . like rocksdb/leveldb api
func (tree *Tree) SeekString(key string) Iterator {
	return tree.Seek([]byte(key))
}

// Index get the Iterator by index(0, 1, 2, 3 ... or -1, -2, -3 ...)
func (tree *Tree) Index(idx int) Iterator {
	node := tree.IndexNode(idx)
	if node != nil {
		return node.IteratorBase(tree)
	}
	return nil
}

// IndexNode get the node by index(0, 1, 2, 3 ... or -1, -2, -3 ...)
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

// IndexKey 第idx个节点的Key值
func (tree *Tree) IndexKey(idx int) ([]byte, bool) {
	n := tree.IndexNode(idx)
	if n != nil {
		return n.key, true
	}
	return nil, false
}

// IndexValue 第idx个节点的Value值
func (tree *Tree) IndexValue(idx int) ([]byte, bool) {
	n := tree.IndexNode(idx)
	if n != nil {
		return n.value, true
	}
	return nil, false
}

// RemoveNode remove the node
func (tree *Tree) removeNodeWithNoFixSize(n *Node) {

}

// RemoveNode remove the node
func (tree *Tree) RemoveNode(n *Node) {
	if tree.root.size == 1 {
		tree.root = nil
		// return n
		return
	}

	ls, rs := getChildrenSize(n)
	if ls == 0 && rs == 0 {
		p := n.parent
		p.children[n.relation] = nil
		tree.fixSizeWithRemove(p)
		// return n
		return
	}

	var cur *Node
	if ls > rs { // 该节点的有序下个或上个节点交换拼接
		cur = n.children[0]
		for cur.children[1] != nil {
			cur = cur.children[1]
		}

		cleft := cur.children[0]
		cur.parent.children[cur.relation] = cleft

		if cleft != nil {
			cleft.parent = cur.parent
			cleft.relation = cur.relation
		}

	} else {
		cur = n.children[1]
		for cur.children[0] != nil {
			cur = cur.children[0]
		}

		cright := cur.children[1]
		cur.parent.children[cur.relation] = cright

		if cright != nil {
			cright.parent = cur.parent
			cright.relation = cur.relation
		}

	}

	// 修改为interface 交换

	cur.children = n.children
	if cur.children[0] != nil {
		cur.children[0].parent = cur
	}
	if cur.children[1] != nil {
		cur.children[1].parent = cur
	}

	var cparent *Node
	if cur.parent != n {
		cparent = cur.parent
	} else {
		cparent = cur
	}

	cur.parent = n.parent
	cur.relation = n.relation
	cur.size = n.size
	if cur.parent != nil {
		cur.parent.children[cur.relation] = cur
	} else {
		tree.root = cur
	}

	//n.key, n.value, cur.key, cur.value = cur.key, cur.value, n.key, n.value

	// 考虑到刚好替换的节点是 被替换节点的孩子节点的时候, 从自身修复高度
	tree.fixSizeWithRemove(cparent)

	// return cur
	return
}

type pnode struct {
	node      *Node
	leftright int
}

type searchpath struct {
	paths []*Node
}

func (sp *searchpath) Append(n *Node, leftright int) {
	sp.paths = append(sp.paths, n)
}

// RemoveRange remove the node [start, end], contain end. not [start, end)
func (tree *Tree) RemoveRange(start, end []byte) int {

	if tree.root.size == 0 {
		return 0
	}

	switch tree.compartor(start, end) {
	case 0:
		if _, ok := tree.Remove(start); ok {
			return 1
		}
		return 0
	case 1:
		start, end = end, start
	}

	var minpath = &searchpath{}
	var maxpath = &searchpath{}

BREAK_LEFT:
	for n := tree.root; ; {
		if n == nil {
			minpath.Append(nil, 1)
			break
		}

		switch c := tree.compartor(start, n.key); c {
		case -1:
			minpath.Append(n, 0)
			n = n.children[0]
		case 1:
			minpath.Append(n, 1)
			n = n.children[1]
		case 0:
			minpath.Append(n, 0)
			minpath.Append(n.children[0], 1)
			break BREAK_LEFT
		default:
			panic("Compare only is allowed in -1, 0, 1")
		}
	}

BREAK_RIGHT:
	for n := tree.root; ; {

		if n == nil {
			maxpath.Append(nil, 0)
			break
		}

		switch c := tree.compartor(end, n.key); c {
		case -1:
			maxpath.Append(n, 0)
			n = n.children[0]
		case 1:
			maxpath.Append(n, 1)
			n = n.children[1]
		case 0:
			maxpath.Append(n, 1)
			maxpath.Append(n.children[1], 0)
			break BREAK_RIGHT
		default:
			panic("Compare only is allowed in -1, 0, 1")
		}
	}

	var rootpath *Node
	reducesize := 0
	for i, min := range minpath.paths {

		if i < len(maxpath.paths) {
			max := maxpath.paths[i]

			if min != max {
				reducesize += tree.removebranch(minpath, i, 0)
				reducesize += tree.removebranch(maxpath, i, 1)
				up := rootpath.parent
				for up != nil {
					up.size -= reducesize
					up = up.parent
				}

				tree.RemoveNode(rootpath)
				return reducesize + 1
			}
			rootpath = min
		} else {
			break
		}
	}

	minlast := minpath.paths[len(minpath.paths)-2] // 倒数第二个为最后一个可能删除的节点
	maxlast := maxpath.paths[len(maxpath.paths)-2]
	if minlast.relation != maxlast.relation {
		tree.RemoveNode(minlast) // 删除最后一个相同
		return 1
	}

	// rootpath.size -= reducesize
	return reducesize
}

func (tree *Tree) removebranch(minpath *searchpath, i int, selchild byte) int {
	reducesize := 0
	curreduce := 0

	var LEFT = selchild
	var RIGHT byte = 1
	if selchild == 1 {
		LEFT = 1
		RIGHT = 0
	}

	var top = minpath.paths[i-1]
	var last = minpath.paths[len(minpath.paths)-1]
	var up = top
	for ii := i; ii < len(minpath.paths)-1; ii++ {
		p := minpath.paths[ii]

		if p.relation == LEFT {
			pright := p.children[RIGHT]
			if pright != nil {
				curreduce += pright.size + 1
			} else {
				curreduce++
			}

			continue
		}

		if p.parent != up {
			p.parent = up
			up.children[up.relation] = p
			for rup := up; rup != top.parent; rup = rup.parent {
				rup.size -= curreduce
			}
		}
		up = p
		reducesize += curreduce
		curreduce = 0
	}

	if last != nil {
		if last.parent != up {
			last.parent = up
			up.children[up.relation] = last
			for rup := up; rup != top.parent; rup = rup.parent {
				rup.size -= curreduce
			}
		}
	} else {
		up.children[up.relation] = nil
		for rup := up; rup != top.parent; rup = rup.parent {
			rup.size -= curreduce
		}
	}

	reducesize += curreduce
	return reducesize
}

// Remove key
func (tree *Tree) Remove(key []byte) ([]byte, bool) {

	if n, ok := tree.GetNode(key); ok {
		tree.RemoveNode(n)
		return n.value, true
	}
	// return nil
	return nil, false
}

// Clear clear tree
func (tree *Tree) Clear() {
	tree.root = nil
	// tree.iter = NewIteratorWithCap(nil, 16)
}

// Values 返回先序遍历的值
func (tree *Tree) Values() [][]byte {
	mszie := 0
	if tree.root != nil {
		mszie = tree.root.size
	}
	result := make([][]byte, 0, mszie)
	tree.Traversal(func(k, v []byte) bool {
		result = append(result, v)
		return true
	}, LDR)
	return result
}

// GetRange get key
func (tree *Tree) GetRange(start, end []byte) (result [][]byte) {

	if node := tree.seekNodeNext(start); node != nil {
		iter := node.IteratorRange(tree)
		if tree.compartor(start, end) == 1 {
			iter.SetLimit(end, start)
			for iter.Prev() {
				result = append(result, iter.Value())
			}
		} else {
			iter.SetLimit(start, end)
			for iter.Next() {
				result = append(result, iter.Value())
			}
		}
	}
	return
}

// Get get key
func (tree *Tree) Get(key []byte) ([]byte, bool) {
	n, ok := tree.GetNode(key)
	if ok {
		return n.value, true
	}
	return nil, false
}

// GetString get key by string
func (tree *Tree) GetString(key string) ([]byte, bool) {
	n, ok := tree.GetNode([]byte(key))
	if ok {
		return n.value, true
	}
	return nil, false
}

// GetNode get node
func (tree *Tree) GetNode(key []byte) (*Node, bool) {

	for n := tree.root; n != nil; {
		switch c := tree.compartor(key, n.key); c {
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

// PutNotCover if key is exists, return false
func (tree *Tree) PutNotCover(key, value []byte) bool {

	if tree.root == nil {
		tree.root = &Node{key: key, value: value, size: 1}
		return true
	}

	for cur := tree.root; ; {

		if cur.size > 8 {
			factor := cur.size >> 3 // or factor = 1
			ls, rs := getChildrenSize(cur)
			if rs >= ls<<1+factor || ls >= rs<<1+factor {
				cur = tree.fixSize(cur, ls, rs)
			}
		}

		c := tree.compartor(key, cur.key)
		switch {
		case c < 0:
			if cur.children[0] == nil {
				node := &Node{key: key, value: value, size: 1, relation: 0}
				cur.children[0] = node
				node.parent = cur

				for temp := cur; temp != nil; temp = temp.parent {
					temp.size++
				}

				if cur.parent != nil && cur.parent.size == 3 {
					if cur.parent.children[0] == nil {
						cur = tree.lrrotate3(cur.parent)
					} else {
						cur = tree.rrotate3(cur.parent)
					}
				}
				return true
			}
			cur = cur.children[0]
		case c > 0:
			if cur.children[1] == nil {
				node := &Node{key: key, value: value, size: 1, relation: 1}

				cur.children[1] = node
				node.parent = cur

				for temp := cur; temp != nil; temp = temp.parent {
					temp.size++
				}

				if cur.parent != nil && cur.parent.size == 3 {
					if cur.parent.children[1] == nil {
						cur = tree.rlrotate3(cur.parent)
					} else {
						cur = tree.lrotate3(cur.parent)
					}
				}
				return true
			}
			cur = cur.children[1]
		default:
			return false
		}
	}
}

// PutString put string
func (tree *Tree) PutString(key, value string) bool {
	return tree.Put([]byte(key), []byte(value))
}

// Put  put bytes
func (tree *Tree) Put(key, value []byte) bool {

	if tree.root == nil {
		tree.root = &Node{key: key, value: value, size: 1, relation: 255}
		return true
	}

	for cur := tree.root; ; {

		if cur.size > 8 {
			factor := cur.size >> 3 // or factor = 1
			ls, rs := getChildrenSize(cur)
			if rs >= ls<<1+factor || ls >= rs<<1+factor {
				cur = tree.fixSize(cur, ls, rs)
			}
		}

		c := tree.compartor(key, cur.key)
		switch {
		case c < 0:
			if cur.children[0] == nil {
				node := &Node{key: key, value: value, size: 1, relation: 0}
				cur.children[0] = node
				node.parent = cur

				for temp := cur; temp != nil; temp = temp.parent {
					temp.size++
				}

				if cur.parent != nil && cur.parent.size == 3 {
					if cur.parent.children[0] == nil {
						cur = tree.lrrotate3(cur.parent)
					} else {
						cur = tree.rrotate3(cur.parent)
					}
				}
				return true
			}
			cur = cur.children[0]
		case c > 0:
			if cur.children[1] == nil {
				node := &Node{key: key, value: value, size: 1, relation: 1}
				cur.children[1] = node
				node.parent = cur

				for temp := cur; temp != nil; temp = temp.parent {
					temp.size++
				}

				if cur.parent != nil && cur.parent.size == 3 {
					if cur.parent.children[1] == nil {
						cur = tree.rlrotate3(cur.parent)
					} else {
						cur = tree.lrotate3(cur.parent)
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

// TraversalMethod 遍历的方式
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

	// BFSLR 宽度遍历  left -> right
	BFSLR
	// BFSRL  right -> left
	BFSRL
)

// Traversal 遍历的方法 默认是LDR 从小到大 Compare 为 l < r
func (tree *Tree) Traversal(every func(k, v []byte) bool, traversalMethod ...TraversalMethod) {
	if tree.root == nil {
		return
	}

	method := LDR
	if len(traversalMethod) != 0 {
		method = traversalMethod[0]
	}

	switch method {
	case DLR:
		var traverasl func(cur *Node) bool
		traverasl = func(cur *Node) bool {
			if cur == nil {
				return true
			}

			if !every(cur.key, cur.value) {
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
			if !every(cur.key, cur.value) {
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
			if !every(cur.key, cur.value) {
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
			if !every(cur.key, cur.value) {
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
			if !every(cur.key, cur.value) {
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
			if !every(cur.key, cur.value) {
				return false
			}
			return true
		}
		traverasl(tree.root)
	case BFSLR:
		queue := linkedlist.New()
		queue.Push(tree.root)

		for icur, ok := queue.PopFront(); ok; icur, ok = queue.PopFront() {
			cur := icur.(*Node)
			if !every(cur.key, cur.value) {
				break
			}
			if cur.children[0] != nil {
				queue.PushBack(cur.children[0])
			}
			if cur.children[1] != nil {
				queue.PushBack(cur.children[1])
			}
		}

	case BFSRL:
		queue := linkedlist.New()
		queue.Push(tree.root)

		for icur, ok := queue.PopFront(); ok; icur, ok = queue.PopFront() {
			cur := icur.(*Node)
			if !every(cur.key, cur.value) {
				break
			}
			if cur.children[1] != nil {
				queue.PushBack(cur.children[1])
			}
			if cur.children[0] != nil {
				queue.PushBack(cur.children[0])
			}
		}

	}
}

func (tree *Tree) lrrotate3(cur *Node) *Node {
	const l = 1
	const r = 0

	movparent := cur.children[l]
	mov := movparent.children[r]

	if cur.parent == nil {
		tree.root = mov
		mov.parent = nil
		mov.relation = 255
	} else {
		cur.parent.children[cur.relation] = mov
		mov.parent = cur.parent
		mov.relation = cur.relation
	}

	cur.children[l] = nil

	cur.parent = mov
	mov.children[r] = cur
	cur.relation = r

	mov.children[l] = movparent
	movparent.parent = mov
	movparent.relation = l

	movparent.children[r] = nil

	mov.size = 3
	cur.size = 1
	movparent.size = 1
	return mov
}

func (tree *Tree) lrrotate(cur *Node) *Node {

	const l = 1
	const r = 0

	movparent := cur.children[l]
	mov := movparent.children[r]

	tree.rrotate(movparent)
	tree.lrotate(cur)

	return mov
}

func (tree *Tree) lrotate3(cur *Node) *Node {
	const l = 1
	const r = 0
	// 1 right 0 left
	mov := cur.children[l]

	if cur.parent == nil {
		tree.root = mov
		mov.parent = nil
		mov.relation = 255
	} else {
		cur.parent.children[cur.relation] = mov
		mov.parent = cur.parent
		mov.relation = cur.relation
	}

	cur.children[l] = nil

	mov.children[r] = cur
	cur.parent = mov
	cur.relation = r

	mov.size = 3
	cur.size = 1

	return mov
}

func (tree *Tree) lrotate(cur *Node) *Node {

	const l = 1
	const r = 0
	// 1 right 0 left
	mov := cur.children[l]
	movright := mov.children[r]

	if cur.parent == nil {
		tree.root = mov
		mov.parent = nil
		mov.relation = 255
	} else {
		cur.parent.children[cur.relation] = mov
		mov.parent = cur.parent
		mov.relation = cur.relation
	}

	if movright != nil {
		cur.children[l] = movright
		movright.parent = cur
		movright.relation = l
	} else {
		cur.children[l] = nil
	}

	mov.children[r] = cur
	cur.parent = mov
	cur.relation = r

	cur.size = getChildrenSumSize(cur) + 1
	mov.size = getChildrenSumSize(mov) + 1

	return mov
}

func (tree *Tree) rlrotate3(cur *Node) *Node {
	const l = 0
	const r = 1

	movparent := cur.children[l]
	mov := movparent.children[r]

	if cur.parent == nil {
		tree.root = mov
		mov.parent = nil
		mov.relation = 255
	} else {
		cur.parent.children[cur.relation] = mov
		mov.parent = cur.parent
		mov.relation = cur.relation
	}

	cur.children[l] = nil

	cur.parent = mov
	mov.children[r] = cur
	cur.relation = r

	mov.children[l] = movparent
	movparent.parent = mov
	movparent.relation = l

	movparent.children[r] = nil

	mov.size = 3
	cur.size = 1
	movparent.size = 1
	return mov
}

func (tree *Tree) rlrotate(cur *Node) *Node {

	const l = 0
	const r = 1

	movparent := cur.children[l]
	mov := movparent.children[r]

	tree.lrotate(movparent)
	tree.rrotate(cur)

	return mov
}

func (tree *Tree) rrotate3(cur *Node) *Node {
	const l = 0
	const r = 1

	// 1 right 0 left
	mov := cur.children[l]

	if cur.parent == nil {
		tree.root = mov
		mov.parent = nil
		mov.relation = 255
	} else {
		cur.parent.children[cur.relation] = mov
		mov.parent = cur.parent
		mov.relation = cur.relation
	}

	cur.children[l] = nil

	mov.children[r] = cur
	cur.parent = mov
	cur.relation = r

	mov.size = 3
	cur.size = 1

	return mov
}

func (tree *Tree) rrotate(cur *Node) *Node {

	const l = 0
	const r = 1
	// 1 right 0 left
	mov := cur.children[l]
	movright := mov.children[r]

	if cur.parent == nil {
		tree.root = mov
		mov.parent = nil
		mov.relation = 255
	} else {
		cur.parent.children[cur.relation] = mov
		mov.parent = cur.parent
		mov.relation = cur.relation
	}

	if movright != nil {
		cur.children[l] = movright
		movright.parent = cur
		movright.relation = l
	} else {
		cur.children[l] = nil
	}

	mov.children[r] = cur
	cur.parent = mov
	cur.relation = r

	cur.size = getChildrenSumSize(cur) + 1
	mov.size = getChildrenSumSize(mov) + 1

	return mov
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
			factor := cur.size >> 3 // or factor = 1
			ls, rs := getChildrenSize(cur)
			if rs >= ls<<1+factor || ls >= rs<<1+factor {
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

func (tree *Tree) fixSize(cur *Node, ls, rs int) *Node {
	if ls > rs {
		llsize, lrsize := getChildrenSize(cur.children[0])
		if lrsize > llsize {
			return tree.rlrotate(cur)
		}
		return tree.rrotate(cur)

	}

	rlsize, rrsize := getChildrenSize(cur.children[1])
	if rlsize > rrsize {
		return tree.lrrotate(cur)
	}
	return tree.lrotate(cur)
}

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

	*str += "(" + spew.Sprint(string(node.key)) + "->" + spew.Sprint(string(node.value)) + ")" + "\n"

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
		parentv = spew.Sprint(string(node.parent.key))
	}
	suffix += parentv + "|" + spew.Sprint(node.size) + ")"
	*str += spew.Sprint(string(node.key)) + suffix + "\n"

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

func (tree *Tree) debugString() string {
	str := "VTree:\n"
	if tree.root == nil {
		return str + "nil"
	}
	outputfordebug(tree.root, "", true, &str)
	return str
}
