package vtree

import (
	linkedlist "github.com/474420502/focus/list/linked_list"
	"github.com/davecgh/go-spew/spew"
)

// Node tree的节点
type Node struct {
	children [2]*Node
	parent   *Node
	size     int
	key      []byte
	value    []byte
}

// Iterator return iterator and start by node
func (n *Node) Iterator() *Iterator {
	return NewIterator(n)
}

// Key get node key
func (n *Node) Key() []byte {
	return n.value
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
		p = spew.Sprint(n.parent.value)
	}
	return spew.Sprint(n.value) + "(" + p + "|" + spew.Sprint(n.size) + ")"
}

// Tree increasing
type Tree struct {
	root *Node
	// iter *Iterator
}

func assertImplementation() {

}

func New() *Tree {
	return &Tree{}
}

func (tree *Tree) String() string {
	str := "VTree-Key\n"
	if tree.root == nil {
		return str + "nil"
	}
	output(tree.root, "", true, &str)
	return str
}

// func (tree *Tree) Iterator() *Iterator {
// 	return initIterator(tree)
// }

// Size get tree size
func (tree *Tree) Size() int {
	if tree.root == nil {
		return 0
	}
	return tree.root.size
}

func comparebyte(s1, s2 []byte) int {

	switch {
	case len(s1) > len(s2):
		// for i := 0; i < len(s2); i++ {
		// 	if s1[i] != s2[i] {
		// 		if s1[i] > s2[i] {
		// 			return 1
		// 		}
		// 		return -1
		// 	}
		// }
		return 1
	case len(s1) < len(s2):
		// for i := 0; i < len(s1); i++ {
		// 	if s1[i] != s2[i] {
		// 		if s1[i] > s2[i] {
		// 			return 1
		// 		}
		// 		return -1
		// 	}
		// }
		return -1
	default:
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if s1[i] > s2[i] {
					return 1
				}
				return -1
			}
		}
		return 0
	}
}

func compare(k1 []byte, k2 []byte) int {
	return comparebyte(k1, k2)
}

// SeekRange start end
func (tree *Tree) SeekRange(start, end []byte) *Iterator {
	iter := tree.Seek(start)
	if compare(start, end) == 1 {
		iter.SetLimit(end, start)
	} else {
		iter.SetLimit(start, end)
	}
	return iter
}

// Seek TODO: 错误
func (tree *Tree) Seek(key []byte) *Iterator {
	lastc := 0
	var n, switchParent, lastn *Node
	for n = tree.root; n != nil; {

		c := compare(key, n.key)
		if lastc*c == -1 {
			switchParent = n.parent
		}
		lastc = c

		switch c {
		case -1:
			lastn = n
			n = n.children[0]
		case 1:
			lastn = n
			n = n.children[1]
		case 0:
			return n.Iterator()
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}

	}

	switch lastc {
	case -1:
		return lastn.Iterator()
	case 1:
		if switchParent != nil {
			return switchParent.Iterator()
		} else {
			return lastn.Iterator()
		}

	default:
		return nil
	}

	// log.Println(lastc, string(lastn.key))
	// return switchParent.Iterator()
}

// Index get the Iterator by index(0, 1, 2, 3 ... or -1, -2, -3 ...)
func (tree *Tree) Index(idx int) *Iterator {
	node := tree.indexNode(idx)
	if node != nil {
		return node.Iterator()
	}
	return nil
}

// indexNode get the node by index(0, 1, 2, 3 ... or -1, -2, -3 ...)
func (tree *Tree) indexNode(idx int) *Node {
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
	n := tree.indexNode(idx)
	if n != nil {
		return n.key, true
	}
	return nil, false
}

// IndexValue 第idx个节点的Value值
func (tree *Tree) IndexValue(idx int) ([]byte, bool) {
	n := tree.indexNode(idx)
	if n != nil {
		return n.value, true
	}
	return nil, false
}

// func (tree *Tree) IndexRange(idx1, idx2 int) (result []interface{}, ok bool) { // 0 -1

// 	if idx1^idx2 < 0 {
// 		if idx1 < 0 {
// 			idx1 = tree.root.size + idx1
// 		} else {
// 			idx2 = tree.root.size + idx2
// 		}
// 	}

// 	if idx1 > idx2 {
// 		ok = true
// 		if idx1 >= tree.root.size {
// 			idx1 = tree.root.size - 1
// 			ok = false
// 		}

// 		n := tree.IndexNode(idx1)
// 		tree.iter.SetNode(n)
// 		iter := tree.iter
// 		result = make([]interface{}, 0, idx1-idx2)
// 		for i := idx2; i <= idx1; i++ {
// 			if iter.Prev() {
// 				result = append(result, iter.Value())
// 			} else {
// 				ok = false
// 				return
// 			}
// 		}

// 		return

// 	} else {
// 		ok = true
// 		if idx2 >= tree.root.size {
// 			idx2 = tree.root.size - 1
// 			ok = false
// 		}

// 		if n := tree.IndexNode(idx1); n != nil {
// 			tree.iter.SetNode(n)
// 			iter := tree.iter
// 			result = make([]interface{}, 0, idx2-idx1)
// 			for i := idx1; i <= idx2; i++ {
// 				if iter.Next() {
// 					result = append(result, iter.Value())
// 				} else {
// 					ok = false
// 					return
// 				}
// 			}

// 			return
// 		}

// 	}

// 	return nil, false
// }

// func (tree *Tree) RemoveIndex(idx int) (interface{}, bool) {
// 	n := tree.IndexNode(idx)
// 	if n != nil {
// 		tree.RemoveNode(n)
// 		return n.value, true
// 	}
// 	return nil, false
// }

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
		p.children[getRelationship(n)] = nil
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
	n.key, n.value, cur.key, cur.value = cur.key, cur.value, n.key, n.value

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
	paths []*pnode
}

func (sp *searchpath) Append(n *Node, leftright int) {
	sp.paths = append(sp.paths, &pnode{n, leftright})
}

// RemoveRange remove the node [start, end], contain end. not [start, end)
func (tree *Tree) RemoveRange(start, end []byte) {

	if tree.root.size == 0 {
		return
	}

	switch compare(start, end) {
	case 0:
		tree.Remove(start)
		return
	case 1:
		start, end = end, start
	}

	var minpath = &searchpath{}
	var maxpath = &searchpath{}

BREAK_LEFT:
	for n := tree.root; ; {
		// log.Println(string(n.key))

		if n == nil {
			minpath.Append(nil, 1)
			break
		}

		switch c := compare(start, n.key); c {
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

		switch c := compare(end, n.key); c {
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

	var rootpath *pnode
	reducesize := 0
	for i, min := range minpath.paths {

		if i < len(maxpath.paths) {
			max := maxpath.paths[i]

			if min.node != max.node {
				reducesize += tree.removebranch(minpath, i, 0)
				reducesize += tree.removebranch(maxpath, i, 1)
				up := rootpath.node.parent
				for up != nil {
					up.size -= reducesize
					up = up.parent
				}
				tree.RemoveNode(rootpath.node)
				return
			}
			rootpath = min
		} else {
			break
		}
	}

	minlast := minpath.paths[len(minpath.paths)-2] // 倒数第二个为最后一个可能删除的节点
	maxlast := maxpath.paths[len(maxpath.paths)-2]
	if minlast.leftright != maxlast.leftright {
		tree.RemoveNode(minlast.node) // 删除最后一个相同
	}

	// rootpath.size -= reducesize

	// log.Println(reducesize, string(rootpath.node.key))

}

func (tree *Tree) removebranch(minpath *searchpath, i int, selchild int) int {
	reducesize := 0
	curreduce := 0

	var LEFT = selchild
	var RIGHT = 1
	if selchild == 1 {
		LEFT = 1
		RIGHT = 0
	}
	// log.Println(RIGHT)
	var top = minpath.paths[i-1]
	var last = minpath.paths[len(minpath.paths)-1]
	var up = top
	for ii := i; ii < len(minpath.paths)-1; ii++ {
		p := minpath.paths[ii]
		// log.Println(string(up.node.key), string(p.node.key))
		if p.leftright == LEFT {
			pright := p.node.children[RIGHT]
			if pright != nil {
				curreduce += pright.size + 1
				// reducesize += pright.size + 1
			} else {
				curreduce++
			}

			continue
		}

		if p.node.parent != up.node {
			p.node.parent = up.node
			up.node.children[up.leftright] = p.node
			for rup := up.node; rup != top.node.parent; rup = rup.parent {
				rup.size -= curreduce
			}
		}
		up = p
		reducesize += curreduce
		curreduce = 0
	}

	if last.node != nil {
		if last.node.parent != up.node {
			last.node.parent = up.node
			up.node.children[up.leftright] = last.node
			for rup := up.node; rup != top.node.parent; rup = rup.parent {
				rup.size -= curreduce
			}
		}
	} else {
		up.node.children[up.leftright] = nil
		for rup := up.node; rup != top.node.parent; rup = rup.parent {
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
	iter := tree.Seek(start)
	if compare(start, end) == 1 {
		iter.SetLimit(end, start)
		for iter.PrevLimit() {
			result = append(result, iter.Value())
		}
	} else {
		iter.SetLimit(start, end)
		for iter.NextLimit() {
			result = append(result, iter.Value())
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

// GetNode get node
func (tree *Tree) GetNode(key []byte) (*Node, bool) {

	for n := tree.root; n != nil; {
		switch c := compare(key, n.key); c {
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
			factor := cur.size / 10 // or factor = 1
			ls, rs := cur.children[0].size, cur.children[1].size
			if rs >= ls*2+factor || ls >= rs*2+factor {
				tree.fixSize(cur, ls, rs)
			}
		}

		c := compare(key, cur.key)
		switch {
		case c < 0:
			if cur.children[0] == nil {
				node := &Node{key: key, value: value, size: 1}
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
				node := &Node{key: key, value: value, size: 1}

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
		tree.root = &Node{key: key, value: value, size: 1}
		return true
	}

	for cur := tree.root; ; {

		if cur.size > 8 {
			factor := cur.size / 10 // or factor = 1
			ls, rs := cur.children[0].size, cur.children[1].size
			if rs >= ls*2+factor || ls >= rs*2+factor {
				tree.fixSize(cur, ls, rs)
			}
		}

		c := compare(key, cur.key)
		switch {
		case c < 0:
			if cur.children[0] == nil {
				node := &Node{key: key, value: value, size: 1}
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
				node := &Node{key: key, value: value, size: 1}
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

func (tree *Tree) lrrotate3(cur *Node) {
	const l = 1
	const r = 0

	movparent := cur.children[l]
	mov := movparent.children[r]

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

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

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

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

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

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

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

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

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

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

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

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

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

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

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

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

	*str += "(" + spew.Sprint(string(node.key)) + "->" + spew.Sprint(string(node.value)) + ")" + "\n"

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
		parentv = spew.Sprint(string(node.parent.key))
	}
	suffix += parentv + "|" + spew.Sprint(node.size) + ")"
	*str += spew.Sprint(string(node.key)) + suffix + "\n"

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
	str := "AVLTree\n"
	if tree.root == nil {
		return str + "nil"
	}
	outputfordebug(tree.root, "", true, &str)
	return str
}
