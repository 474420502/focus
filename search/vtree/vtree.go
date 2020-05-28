package vtree

import (
	"log"

	linkedlist "github.com/474420502/focus/list/linked_list"
	"github.com/davecgh/go-spew/spew"
)

type Node struct {
	children [2]*Node
	parent   *Node
	size     int
	key      []byte
	value    []byte
}

func (n *Node) Iterator() *Iterator {
	return NewIterator(n)
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

// func (tree *Tree) IndexNode(idx int) *Node {
// 	cur := tree.root
// 	if idx >= 0 {
// 		for cur != nil {
// 			ls := getSize(cur.children[0])
// 			if idx == ls {
// 				return cur
// 			} else if idx < ls {
// 				cur = cur.children[0]
// 			} else {
// 				idx = idx - ls - 1
// 				cur = cur.children[1]
// 			}
// 		}
// 	} else {
// 		idx = -idx - 1
// 		for cur != nil {
// 			rs := getSize(cur.children[1])
// 			if idx == rs {
// 				return cur
// 			} else if idx < rs {
// 				cur = cur.children[1]
// 			} else {
// 				idx = idx - rs - 1
// 				cur = cur.children[0]
// 			}
// 		}
// 	}
// 	return nil
// }

// func (tree *Tree) Index(idx int) (interface{}, bool) {
// 	n := tree.IndexNode(idx)
// 	if n != nil {
// 		return n.value, true
// 	}
// 	return nil, false
// }

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

// RemoveRange remove the node
func (tree *Tree) RemoveRange(start, end []byte) {

	if compare(start, end) == 1 {
		start, end = end, start
	}

	siter := tree.Seek(start)
	siter.SetLimit(start, end)

	if siter.NextLimit() {
		min := siter.GetNode()

		eiter := tree.Seek(end)
		eiter.SetLimit(start, end)
		if !eiter.PrevLimit() {
			panic("max is not exist, check tree")
		}
		max := eiter.GetNode()
		log.Println(string(start), string(max.value))

		cur := min
		preducesize := 0
		checknode := min


		checknode := min

		for cur != nil {
			parent := cur.parent 

			cright := cur.children[1]
			if cright != nil {
				preducesize += cright.size 
			}
			preducesize++

			cleft := cur.children[0]
			if cleft != nil {
				cleft.parent = cur.parent
			} 

			if cur.parent != nil {
				relation := getRelationship(cur)
				cur.parent.children[relation] = cleft
				
				if relation == 0 {
					switch compare(max.key, cur.parent.key) {
					case 1:
						
					}
				}

			}else {
				tree.root = cleft
			}
		}

		// for cur.parent != nil {
		// 	relation := getRelationship(cur)
		// 	if relation == 0 {
		// 		switch compare(max.key, cur.parent.key) {
		// 		case 1:

		// 			cright := cur.children[1]
		// 			if cright != nil {
		// 				preducesize += (cur.size - cright.size)
		// 				// cur.parent.size -= preducesize
		// 				cur.parent.children[1] = cright
		// 				cright.parent = cur.parent
		// 			}

		// 			for checknode != cur {
		// 				checknode.size -= preducesize
		// 				checknode = checknode.parent
		// 			}
		// 			cur.size -= preducesize
		// 			checknode = cur

		// 		case -1: // 确认39最大

		// 			child := cur.children[1]
		// 			for child != nil {
		// 				// cur.children[1] = nil
		// 				cright := child.children[1]
		// 				if cright != nil {
		// 					switch compare(max.key, cur.parent.key) { 
		// 					case 1:
		// 						preducesize += (child.size - cright.size)
		// 						// cur.parent.size -= preducesize
		// 						child.parent.children[1] = cright
		// 						cright.parent = child.parent
		// 						child = cright
		// 					case -1:
		// 						child = child.children[0]
		// 					}
		// 				}
		// 				cur = cur.parent
		// 			}

		// 		default: // ==0的时候
		// 		}
		// 	}
		// 	cur = cur.parent
		// }

		// if cur.parent != nil {
		// FIND_RIGHT:
		// 	for cur.parent != nil {
		// 		switch compare(max.key, cur.parent.key) {
		// 		case 1:

		// 			// cur.children[1] = nil
		// 			cleft := cur.children[0]
		// 			if cleft != nil {
		// 				preducesize += (cur.size - cleft.size)
		// 				// cur.parent.size -= preducesize
		// 				cur.parent.children[0] = cleft
		// 				cleft.parent = cur.parent
		// 			}
		// 			cur = cur.parent

		// 			// TOOD: 计算 size
		// 		case -1:

		// 			child := cur.children[1]
		// 			// preducesize++

		// 			for child != nil {
		// 				switch compare(max.key, child.key) {
		// 				case 1:

		// 					// 删除左边
		// 					cright := child.children[1]
		// 					preducesize += child.size - cright.size
		// 					cright.parent = cur
		// 					cur.children[1] = cright
		// 					child = cright

		// 				case -1:
		// 					child = child.children[0]
		// 				default:
		// 					tree.RemoveNode(child)
		// 					break FIND_RIGHT
		// 				}
		// 			}

		// 		default:
		// 			parent := cur.parent
		// 			if parent == nil {
		// 				tree.root = nil
		// 				return
		// 			}

		// 			cleft := cur.children[0]
		// 			if cleft != nil {
		// 				preducesize += (cur.size - cleft.size) + 1
		// 				cleft.parent = parent.parent
		// 			}
		// 			parent.children[0] = cleft
		// 			break FIND_RIGHT
		// 		}
		// 	}

		// }

		// for temp := cur; temp != nil; temp = temp.parent {
		// 	temp.size -= preducesize
		// }

	}
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

	*str += spew.Sprint(node.key) + ":" + spew.Sprint(node.value) + "\n"

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
