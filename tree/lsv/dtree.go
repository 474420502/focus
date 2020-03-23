package lsv

import (
	"log"

	"github.com/davecgh/go-spew/spew"
)

// DNode  节点
type DNode struct {
	family [3]*DNode
	size   int
	key    []rune
	value  []rune
}

// DTree 用于数据的树
type DTree struct {
	root    *DNode
	feature *DNode
	Compare func(s1, s2 []rune) int
	iter    *Iterator
}

func compareRunes(s1, s2 []rune) int {
	switch {
	case len(s1) > len(s2):
		for i := 0; i < len(s2); i++ {
			if s1[i] != s2[i] {
				if s1[i] > s2[i] {
					return 1
				}
				return -1
			}
		}
		return 1
	case len(s1) < len(s2):
		for i := 0; i < len(s1); i++ {
			if s1[i] != s2[i] {
				if s1[i] > s2[i] {
					return 1
				}
				return -1
			}
		}
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

func (n *DNode) String() string {
	if n == nil {
		return "nil"
	}

	p := "nil"
	if n.family[0] != nil {
		p = spew.Sprint(n.family[0].value)
	}
	return spew.Sprint(n.value) + "(" + p + "|" + spew.Sprint(n.size) + ")"
}

// func assertImplementation() {
// 	var _ tree.IBSTreeDupKey = (*Tree)(nil)
// }

// newDataTree 创建一个树
func newDataTree(Compare func(s1, s2 []rune) int) *DTree {
	return &DTree{Compare: Compare, iter: NewIteratorWithCap(nil, 16)}
}

func (tree *DTree) String() string {
	str := "VBTree-Dup\n"
	if tree.root == nil {
		return str + "nil"
	}
	output(tree.root, "", true, &str)
	return str
}

func (tree *DTree) Iterator() *Iterator {
	return initIterator(tree)
}

func (tree *DTree) Size() int {
	if tree.root == nil {
		return 0
	}
	return tree.root.size
}

// IndexNode 索引节点
func (tree *DTree) IndexNode(idx int) *DNode {
	cur := tree.root
	if idx >= 0 {
		for cur != nil {
			ls := getSize(cur.family[1])
			if idx == ls {
				return cur
			} else if idx < ls {
				cur = cur.family[1]
			} else {
				idx = idx - ls - 1
				cur = cur.family[2]
			}
		}
	} else {
		idx = -idx - 1
		for cur != nil {
			rs := getSize(cur.family[2])
			if idx == rs {
				return cur
			} else if idx < rs {
				cur = cur.family[2]
			} else {
				idx = idx - rs - 1
				cur = cur.family[1]
			}
		}
	}
	return nil
}

func (tree *DTree) Index(idx int) (interface{}, bool) {
	n := tree.IndexNode(idx)
	if n != nil {
		return n.value, true
	}
	return nil, false
}

func (tree *DTree) IndexRange(idx1, idx2 int) (result []interface{}, ok bool) { // 0 -1

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

func (tree *DTree) RemoveIndex(idx int) (interface{}, bool) {
	n := tree.IndexNode(idx)
	if n != nil {
		tree.RemoveNode(n)
		return n.value, true
	}
	return nil, false
}

func (tree *DTree) RemoveNode(n *DNode) {
	if tree.root.size == 1 {
		tree.root = nil
		tree.feature = nil
		// return n
		return
	}

	if n == tree.feature {
		iter := NewIterator(n)
		iter.Prev()
		tree.feature = iter.cur
	}

	ls, rs := getChildrenSize(n)
	if ls == 0 && rs == 0 {
		p := n.family[0]
		p.family[getRelationship(n)] = nil
		tree.fixSizeWithRemove(p)
		// return n
		return
	}

	var cur *DNode
	if ls > rs {
		cur = n.family[1]
		for cur.family[2] != nil {
			cur = cur.family[2]
		}

		cleft := cur.family[1]
		cur.family[0].family[getRelationship(cur)] = cleft
		if cleft != nil {
			cleft.family[0] = cur.family[0]
		}

	} else {
		cur = n.family[2]
		for cur.family[1] != nil {
			cur = cur.family[1]
		}

		cright := cur.family[2]
		cur.family[0].family[getRelationship(cur)] = cright

		if cright != nil {
			cright.family[0] = cur.family[0]
		}
	}

	cparent := cur.family[0]
	// 修改为interface 交换
	n.key, n.value, cur.key, cur.value = cur.key, cur.value, n.key, n.value

	// 考虑到刚好替换的节点是 被替换节点的孩子节点的时候, 从自身修复高度
	if cparent == n {
		tree.fixSizeWithRemove(n)
	} else {
		tree.fixSizeWithRemove(cparent)
	}

	// return cur
	return
}

func (tree *DTree) Remove(key []rune) (interface{}, bool) {

	if n, ok := tree.GetNode(key); ok {
		tree.RemoveNode(n)
		return n.value, true
	}
	// return nil
	return nil, false
}

func (tree *DTree) Clear() {
	tree.root = nil
	tree.iter = NewIteratorWithCap(nil, 16)
}

// Values 返回先序遍历的值
func (tree *DTree) Values() [][]rune {
	mszie := 0
	if tree.root != nil {
		mszie = tree.root.size
	}
	result := make([][]rune, 0, mszie)
	tree.Traversal(func(k, v []rune) bool {
		result = append(result, v)
		return true
	}, LDR)
	return result
}

func (tree *DTree) GetRange(k1, k2 []rune) (result []interface{}) {
	c := tree.Compare(k2, k1)
	switch c {
	case 1:

		var min, max *DNode
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

		var min, max *DNode
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

func (tree *DTree) GetString(key string) (interface{}, bool) {
	n, ok := tree.GetNode([]rune(key))
	if ok {
		return n.value, true
	}
	return n, false
}

func (tree *DTree) Get(key []rune) ([]rune, bool) {
	n, ok := tree.GetNode(key)
	if ok {
		return n.value, true
	}
	return nil, false
}

func (tree *DTree) GetAround(key []rune) (result [3]interface{}) {
	an := tree.getArountNode(key)
	for i, n := range an {
		if n != nil {
			result[i] = n.value
		}
	}
	return
}

func (tree *DTree) getArountNode(key []rune) (result [3]*DNode) {
	var last *DNode
	var lastc int

	for n := tree.root; n != nil; {
		last = n
		c := tree.Compare(key, n.key)
		switch c {
		case -1:
			n = n.family[1]
			lastc = c
		case 1:
			n = n.family[2]
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

func (tree *DTree) GetNode(key []rune) (*DNode, bool) {

	for n := tree.root; n != nil; {
		switch c := tree.Compare(key, n.key); c {
		case -1:
			n = n.family[1]
		case 1:
			n = n.family[2]
		case 0:

			tree.iter.SetNode(n)
			iter := tree.iter
			iter.Prev()
			for iter.Prev() {
				if tree.Compare(iter.cur.key, n.key) == 0 {
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

// PutString Key Value with Type string
func (tree *DTree) PutString(key, value string) (isInsert bool) {
	return tree.Put([]rune(key), []rune(value))
}

func (tree *DTree) putfeature(node *DNode) {
	for cur := tree.root; cur != nil; cur = cur.family[2] {
		cur.size++
		if cur.family[2] == nil {
			cur.family[2] = node
			node.family[0] = cur
			tree.feature = node
			return
		}
	}
	log.Println("error")
}

// Put return bool
func (tree *DTree) Put(key, value []rune) (isInsert bool) {

	if tree.root == nil {
		node := &DNode{key: key, value: value, size: 1}
		tree.root = node
		tree.feature = node
		return true
	}

	for cur := tree.root; ; {

		if cur.size > 8 {
			factor := cur.size >> 3 // or factor = 1
			ls, rs := cur.family[1].size, cur.family[2].size
			if rs >= (ls<<1)+factor || ls >= (rs<<1)+factor {
				tree.fixSize(cur, ls, rs)
			}
		}

		c := tree.Compare(key, cur.key)
		switch {
		case c < 0:
			if cur.family[1] == nil {
				node := &DNode{key: key, value: value, size: 1}
				cur.family[1] = node
				node.family[0] = cur

				for temp := cur; temp != nil; temp = temp.family[0] {
					temp.size++
				}

				if cur.family[0] != nil && cur.family[0].size == 3 {
					if cur.family[0].family[1] == nil {
						tree.lrrotate3(cur.family[0])
					} else {
						tree.rrotate3(cur.family[0])
					}
				}

				return true
			}
			cur = cur.family[1]
		case c > 0:
			if cur.family[2] == nil {
				node := &DNode{key: key, value: value, size: 1}
				cur.family[2] = node
				node.family[0] = cur

				for temp := cur; temp != nil; temp = temp.family[0] {
					temp.size++
				}

				if cur.family[0] != nil && cur.family[0].size == 3 {
					if cur.family[0].family[2] == nil {
						tree.rlrotate3(cur.family[0])
					} else {
						tree.lrotate3(cur.family[0])
					}
				}

				if tree.Compare(node.key, tree.feature.key) > 0 {
					tree.feature = node
				}
				return true
			}
			cur = cur.family[2]
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
func (tree *DTree) Traversal(every func(k, v []rune) bool, traversalMethod ...interface{}) {
	if tree.root == nil {
		return
	}

	method := LDR
	if len(traversalMethod) != 0 {
		method = traversalMethod[0].(TraversalMethod)
	}

	switch method {
	case DLR:
		var traverasl func(cur *DNode) bool
		traverasl = func(cur *DNode) bool {
			if cur == nil {
				return true
			}
			if !every(cur.key, cur.value) {
				return false
			}
			if !traverasl(cur.family[1]) {
				return false
			}
			if !traverasl(cur.family[2]) {
				return false
			}
			return true
		}
		traverasl(tree.root)
	case LDR:
		var traverasl func(cur *DNode) bool
		traverasl = func(cur *DNode) bool {
			if cur == nil {
				return true
			}
			if !traverasl(cur.family[1]) {
				return false
			}
			if !every(cur.key, cur.value) {
				return false
			}
			if !traverasl(cur.family[2]) {
				return false
			}
			return true
		}
		traverasl(tree.root)
	case LRD:
		var traverasl func(cur *DNode) bool
		traverasl = func(cur *DNode) bool {
			if cur == nil {
				return true
			}
			if !traverasl(cur.family[1]) {
				return false
			}
			if !traverasl(cur.family[2]) {
				return false
			}
			if !every(cur.key, cur.value) {
				return false
			}
			return true
		}
		traverasl(tree.root)
	case DRL:
		var traverasl func(cur *DNode) bool
		traverasl = func(cur *DNode) bool {
			if cur == nil {
				return true
			}
			if !every(cur.key, cur.value) {
				return false
			}
			if !traverasl(cur.family[1]) {
				return false
			}
			if !traverasl(cur.family[2]) {
				return false
			}
			return true
		}
		traverasl(tree.root)
	case RDL:
		var traverasl func(cur *DNode) bool
		traverasl = func(cur *DNode) bool {
			if cur == nil {
				return true
			}
			if !traverasl(cur.family[2]) {
				return false
			}
			if !every(cur.key, cur.value) {
				return false
			}
			if !traverasl(cur.family[1]) {
				return false
			}
			return true
		}
		traverasl(tree.root)
	case RLD:
		var traverasl func(cur *DNode) bool
		traverasl = func(cur *DNode) bool {
			if cur == nil {
				return true
			}
			if !traverasl(cur.family[2]) {
				return false
			}
			if !traverasl(cur.family[1]) {
				return false
			}
			if !every(cur.key, cur.value) {
				return false
			}
			return true
		}
		traverasl(tree.root)
	}
}

func (tree *DTree) lrrotate3(cur *DNode) {
	const l = 2
	const r = 1

	movparent := cur.family[l]
	mov := movparent.family[r]

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

	cur.family[r] = mov
	mov.family[0] = cur

	cur.family[l] = movparent
	movparent.family[r] = nil

	cur.family[r] = mov
	mov.family[0] = cur

	// cur.size = 3
	// cur.family[r].size = 1
	cur.family[l].size = 1
}

func (tree *DTree) lrrotate(cur *DNode) {

	const l = 2
	const r = 1

	movparent := cur.family[l]
	mov := movparent.family[r]

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

	if mov.family[l] != nil {
		movparent.family[r] = mov.family[l]
		movparent.family[r].family[0] = movparent
		//movparent.family[r].child = l
	} else {
		movparent.family[r] = nil
	}

	if mov.family[r] != nil {
		mov.family[l] = mov.family[r]
		//mov.family[l].child = l
	} else {
		mov.family[l] = nil
	}

	if cur.family[r] != nil {
		mov.family[r] = cur.family[r]
		mov.family[r].family[0] = mov
	} else {
		mov.family[r] = nil
	}

	cur.family[r] = mov
	mov.family[0] = cur

	movparent.size = getChildrenSumSize(movparent) + 1
	mov.size = getChildrenSumSize(mov) + 1
	cur.size = getChildrenSumSize(cur) + 1
}

func (tree *DTree) rlrotate3(cur *DNode) {
	const l = 1
	const r = 2

	movparent := cur.family[l]
	mov := movparent.family[r]

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

	cur.family[r] = mov
	mov.family[0] = cur

	cur.family[l] = movparent
	movparent.family[r] = nil

	cur.family[r] = mov
	mov.family[0] = cur

	// cur.size = 3
	// cur.family[r].size = 1
	cur.family[l].size = 1
}

func (tree *DTree) rlrotate(cur *DNode) {

	const l = 1
	const r = 2

	movparent := cur.family[l]
	mov := movparent.family[r]

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

	if mov.family[l] != nil {
		movparent.family[r] = mov.family[l]
		movparent.family[r].family[0] = movparent
	} else {
		movparent.family[r] = nil
	}

	if mov.family[r] != nil {
		mov.family[l] = mov.family[r]
	} else {
		mov.family[l] = nil
	}

	if cur.family[r] != nil {
		mov.family[r] = cur.family[r]
		mov.family[r].family[0] = mov
	} else {
		mov.family[r] = nil
	}

	cur.family[r] = mov
	mov.family[0] = cur

	movparent.size = getChildrenSumSize(movparent) + 1
	mov.size = getChildrenSumSize(mov) + 1
	cur.size = getChildrenSumSize(cur) + 1
}

func (tree *DTree) rrotate3(cur *DNode) {
	const l = 1
	const r = 2
	// 1 right 0 left
	mov := cur.family[l]

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

	cur.family[r] = mov

	cur.family[l] = mov.family[l]
	cur.family[l].family[0] = cur

	mov.family[l] = nil

	mov.size = 1
}

func (tree *DTree) rrotate(cur *DNode) {

	const l = 1
	const r = 2
	// 1 right 0 left
	mov := cur.family[l]

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

	//  mov.family[l]不可能为nil
	mov.family[l].family[0] = cur

	cur.family[l] = mov.family[l]

	// 解决mov节点孩子转移的问题
	if mov.family[r] != nil {
		mov.family[l] = mov.family[r]
	} else {
		mov.family[l] = nil
	}

	if cur.family[r] != nil {
		mov.family[r] = cur.family[r]
		mov.family[r].family[0] = mov
	} else {
		mov.family[r] = nil
	}

	// 连接转移后的节点 由于mov只是与cur交换值,parent不变
	cur.family[r] = mov

	mov.size = getChildrenSumSize(mov) + 1
	cur.size = getChildrenSumSize(cur) + 1
}

func (tree *DTree) lrotate3(cur *DNode) {
	const l = 2
	const r = 1
	// 1 right 0 left
	mov := cur.family[l]

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

	cur.family[r] = mov

	cur.family[l] = mov.family[l]
	cur.family[l].family[0] = cur

	mov.family[l] = nil

	mov.size = 1
}

func (tree *DTree) lrotate(cur *DNode) {

	const l = 2
	const r = 1
	// 1 right 0 left
	mov := cur.family[l]

	mov.key, mov.value, cur.key, cur.value = cur.key, cur.value, mov.key, mov.value //交换值达到, 相对位移

	//  mov.family[l]不可能为nil
	mov.family[l].family[0] = cur

	cur.family[l] = mov.family[l]

	// 解决mov节点孩子转移的问题
	if mov.family[r] != nil {
		mov.family[l] = mov.family[r]
	} else {
		mov.family[l] = nil
	}

	if cur.family[r] != nil {
		mov.family[r] = cur.family[r]
		mov.family[r].family[0] = mov
	} else {
		mov.family[r] = nil
	}

	// 连接转移后的节点 由于mov只是与cur交换值,parent不变
	cur.family[r] = mov

	mov.size = getChildrenSumSize(mov) + 1
	cur.size = getChildrenSumSize(cur) + 1
}

func getChildrenSumSize(cur *DNode) int {
	return getSize(cur.family[1]) + getSize(cur.family[2])
}

func getChildrenSize(cur *DNode) (int, int) {
	return getSize(cur.family[1]), getSize(cur.family[2])
}

func getSize(cur *DNode) int {
	if cur == nil {
		return 0
	}
	return cur.size
}

func (tree *DTree) fixSizeWithRemove(cur *DNode) {
	for cur != nil {
		cur.size--
		if cur.size > 8 {
			factor := cur.size >> 3 // or factor = 1
			ls, rs := getChildrenSize(cur)
			if rs >= (ls<<1)+factor || ls >= (rs<<1)+factor {
				tree.fixSize(cur, ls, rs)
			}
		} else if cur.size == 3 {
			if cur.family[1] == nil {
				if cur.family[2].family[1] == nil {
					tree.lrotate3(cur)
				} else {
					tree.lrrotate3(cur)
				}
			} else if cur.family[2] == nil {
				if cur.family[1].family[2] == nil {
					tree.rrotate3(cur)
				} else {
					tree.rlrotate3(cur)
				}
			}
		}
		cur = cur.family[0]
	}
}

func (tree *DTree) fixSize(cur *DNode, ls, rs int) {
	if ls > rs {
		llsize, lrsize := getChildrenSize(cur.family[1])
		if lrsize > llsize {
			tree.rlrotate(cur)
		} else {
			tree.rrotate(cur)
		}
	} else {
		rlsize, rrsize := getChildrenSize(cur.family[2])
		if rlsize > rrsize {
			tree.lrrotate(cur)
		} else {
			tree.lrotate(cur)
		}
	}
}

func output(node *DNode, prefix string, isTail bool, str *string) {

	if node.family[2] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.family[2], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}

	*str += spew.Sprint(node.key) + ":" + spew.Sprint(node.value) + "\n"

	if node.family[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.family[1], newPrefix, true, str)
	}

}

func outputfordebug(node *DNode, prefix string, isTail bool, str *string) {

	if node.family[2] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		outputfordebug(node.family[2], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}

	suffix := "("
	parentv := ""
	if node.family[0] == nil {
		parentv = "nil"
	} else {
		parentv = spew.Sprint(node.family[0].key)
	}
	suffix += parentv + "|" + spew.Sprint(node.size) + ")"
	*str += spew.Sprint(node.key) + suffix + "\n"

	if node.family[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		outputfordebug(node.family[1], newPrefix, true, str)
	}
}

func (tree *DTree) debugString() string {
	str := "VBTree-Dup\n"
	if tree.root == nil {
		return str + "nil"
	}
	outputfordebug(tree.root, "", true, &str)
	return str
}
