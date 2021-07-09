package avlkeydup

import (
	"github.com/davecgh/go-spew/spew"

	"github.com/474420502/focus/compare"
	"github.com/474420502/focus/tree"
)

const HeightDiff = 1

type Node struct {
	Children   [2]*Node
	parent     *Node
	height     int
	Key, Value interface{}
}

func (n *Node) String() string {
	if n == nil {
		return "nil"
	}

	p := "nil"
	if n.parent != nil {
		p = spew.Sprint(n.parent.Value)
	}
	return spew.Sprint(n.Value) + "(" + p + "|" + spew.Sprint(n.height) + ")"
}

type Tree struct {
	Root    *Node
	size    int
	Compare compare.Compare
	iter    *Iterator

	RotateLog string
	Count     int
}

func assertImplementation() {
	var _ tree.IBSTreeKey = (*Tree)(nil)
}

func New(Compare compare.Compare) *Tree {
	return &Tree{Compare: Compare, iter: NewIteratorWithCap(nil, 16)}
}

func (tree *Tree) String() string {
	if tree.size == 0 {
		return ""
	}
	str := "AVLTree\n"
	output(tree.Root, "", true, &str)

	return str
}

func (tree *Tree) Iterator() *Iterator {
	return initIterator(tree)
}

func (tree *Tree) Size() int {
	return tree.size
}

func (tree *Tree) Remove(key interface{}) (interface{}, bool) {

	if n, ok := tree.GetNode(key); ok {

		tree.size--
		if tree.size == 0 {
			tree.Root = nil
			return n.Value, true
		}

		left := getHeight(n.Children[0])
		right := getHeight(n.Children[1])

		if left == -1 && right == -1 {
			p := n.parent
			p.Children[getRelationship(n)] = nil
			tree.fixRemoveHeight(p)
			return n.Value, true
		}

		var cur *Node
		if left > right {
			cur = n.Children[0]
			for cur.Children[1] != nil {
				cur = cur.Children[1]
			}

			cleft := cur.Children[0]
			cur.parent.Children[getRelationship(cur)] = cleft
			if cleft != nil {
				cleft.parent = cur.parent
			}

		} else {
			cur = n.Children[1]
			for cur.Children[0] != nil {
				cur = cur.Children[0]
			}

			cright := cur.Children[1]
			cur.parent.Children[getRelationship(cur)] = cright

			if cright != nil {
				cright.parent = cur.parent
			}
		}

		cparent := cur.parent
		// 修改为interface 交换
		n.Value, cur.Value = cur.Value, n.Value
		n.Key, cur.Key = cur.Key, n.Key

		// 考虑到刚好替换的节点是 被替换节点的孩子节点的时候, 从自身修复高度
		if cparent == n {
			tree.fixRemoveHeight(n)
		} else {
			tree.fixRemoveHeight(cparent)
		}

		return cur.Value, true
	}

	return nil, false
}

func (tree *Tree) Clear() {
	tree.size = 0
	tree.Root = nil
	tree.iter = NewIteratorWithCap(nil, 16)
}

// Values 返回先序遍历的值
func (tree *Tree) Values() []interface{} {
	mszie := 0
	if tree.Root != nil {
		mszie = tree.size
	}
	result := make([]interface{}, 0, mszie)
	tree.Traversal(func(k, v interface{}) bool {
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
			return []interface{}{n.Value}
		}
		return []interface{}{}
	}

	return
}

func (tree *Tree) Get(key interface{}) (interface{}, bool) {
	n, ok := tree.GetNode(key)
	if ok {
		return n.Value, true
	}
	return n, false
}

func (tree *Tree) GetAround(key interface{}) (result [3]interface{}) {
	an := tree.getArountNode(key)
	for i, n := range an {
		if n != nil {
			result[i] = n.Value
		}
	}
	return
}

func (tree *Tree) getArountNode(key interface{}) (result [3]*Node) {
	var last *Node
	var lastc int

	for n := tree.Root; n != nil; {
		last = n
		c := tree.Compare(key, n.Key)
		switch c {
		case -1:
			n = n.Children[0]
			lastc = c
		case 1:
			n = n.Children[1]
			lastc = c
		case 0:
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

	for n := tree.Root; n != nil; {
		switch c := tree.Compare(key, n.Key); c {
		case -1:
			n = n.Children[0]
		case 1:
			n = n.Children[1]
		case 0:
			return n, true
		default:
			panic("Get Compare only is allowed in -1, 0, 1")
		}
	}
	return nil, false
}

func (tree *Tree) Put(key, value interface{}) {
	tree.RotateLog = ""
	if tree.size == 0 {
		tree.size++
		tree.Root = &Node{Key: key, Value: value}
		return
	}

	for cur, c := tree.Root, 0; ; {
		c = tree.Compare(key, cur.Key)
		if c == -1 {
			if cur.Children[0] == nil {
				tree.size++
				cur.Children[0] = &Node{Key: key, Value: value}
				cur.Children[0].parent = cur
				if cur.height == 0 {
					tree.fixPutHeight(cur)
				}
				return
			}
			cur = cur.Children[0]
		} else if c == 1 {
			if cur.Children[1] == nil {
				tree.size++
				cur.Children[1] = &Node{Key: key, Value: value}
				cur.Children[1].parent = cur
				if cur.height == 0 {
					tree.fixPutHeight(cur)
				}
				return
			}
			cur = cur.Children[1]
		} else {
			cur.Key = key
			cur.Value = value
			return
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
func (tree *Tree) Traversal(every func(k, v interface{}) bool, traversalMethod ...interface{}) {
	if tree.Root == nil {
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
			if !every(cur.Key, cur.Value) {
				return false
			}
			if !traverasl(cur.Children[0]) {
				return false
			}
			if !traverasl(cur.Children[1]) {
				return false
			}
			return true
		}
		traverasl(tree.Root)
	case LDR:
		var traverasl func(cur *Node) bool
		traverasl = func(cur *Node) bool {
			if cur == nil {
				return true
			}
			if !traverasl(cur.Children[0]) {
				return false
			}
			if !every(cur.Key, cur.Value) {
				return false
			}
			if !traverasl(cur.Children[1]) {
				return false
			}
			return true
		}
		traverasl(tree.Root)
	case LRD:
		var traverasl func(cur *Node) bool
		traverasl = func(cur *Node) bool {
			if cur == nil {
				return true
			}
			if !traverasl(cur.Children[0]) {
				return false
			}
			if !traverasl(cur.Children[1]) {
				return false
			}
			if !every(cur.Key, cur.Value) {
				return false
			}
			return true
		}
		traverasl(tree.Root)
	case DRL:
		var traverasl func(cur *Node) bool
		traverasl = func(cur *Node) bool {
			if cur == nil {
				return true
			}
			if !every(cur.Key, cur.Value) {
				return false
			}
			if !traverasl(cur.Children[1]) {
				return false
			}
			if !traverasl(cur.Children[0]) {
				return false
			}
			return true
		}
		traverasl(tree.Root)
	case RDL:
		var traverasl func(cur *Node) bool
		traverasl = func(cur *Node) bool {
			if cur == nil {
				return true
			}
			if !traverasl(cur.Children[1]) {
				return false
			}
			if !every(cur.Key, cur.Value) {
				return false
			}
			if !traverasl(cur.Children[0]) {
				return false
			}
			return true
		}
		traverasl(tree.Root)
	case RLD:
		var traverasl func(cur *Node) bool
		traverasl = func(cur *Node) bool {
			if cur == nil {
				return true
			}
			if !traverasl(cur.Children[1]) {
				return false
			}
			if !traverasl(cur.Children[0]) {
				return false
			}
			if !every(cur.Key, cur.Value) {
				return false
			}
			return true
		}
		traverasl(tree.Root)
	}
}

func (tree *Tree) lrrotate(cur *Node) {

	tree.Count += 2
	tree.RotateLog += string(cur.Key.([]byte)) + ": lrotate->rrotate "

	const l = 1
	const r = 0

	movparent := cur.Children[l]
	mov := movparent.Children[r]

	mov.Value, cur.Value = cur.Value, mov.Value //交换值达到, 相对位移
	mov.Key, cur.Key = cur.Key, mov.Key

	if mov.Children[l] != nil {
		movparent.Children[r] = mov.Children[l]
		movparent.Children[r].parent = movparent
		//movparent.children[r].child = l
	} else {
		movparent.Children[r] = nil
	}

	if mov.Children[r] != nil {
		mov.Children[l] = mov.Children[r]
		//mov.children[l].child = l
	} else {
		mov.Children[l] = nil
	}

	if cur.Children[r] != nil {
		mov.Children[r] = cur.Children[r]
		mov.Children[r].parent = mov
	} else {
		mov.Children[r] = nil
	}

	cur.Children[r] = mov
	mov.parent = cur

	mov.height = getMaxChildrenHeight(mov) + 1
	movparent.height = getMaxChildrenHeight(movparent) + 1
	cur.height = getMaxChildrenHeight(cur) + 1
}

func (tree *Tree) rlrotate(cur *Node) {

	tree.Count += 2
	tree.RotateLog += string(cur.Key.([]byte)) + ": rrotate->lrotate "

	const l = 0
	const r = 1

	movparent := cur.Children[l]
	mov := movparent.Children[r]

	mov.Value, cur.Value = cur.Value, mov.Value //交换值达到, 相对位移
	mov.Key, cur.Key = cur.Key, mov.Key

	if mov.Children[l] != nil {
		movparent.Children[r] = mov.Children[l]
		movparent.Children[r].parent = movparent
	} else {
		movparent.Children[r] = nil
	}

	if mov.Children[r] != nil {
		mov.Children[l] = mov.Children[r]
	} else {
		mov.Children[l] = nil
	}

	if cur.Children[r] != nil {
		mov.Children[r] = cur.Children[r]
		mov.Children[r].parent = mov
	} else {
		mov.Children[r] = nil
	}

	cur.Children[r] = mov
	mov.parent = cur

	mov.height = getMaxChildrenHeight(mov) + 1
	movparent.height = getMaxChildrenHeight(movparent) + 1
	cur.height = getMaxChildrenHeight(cur) + 1
}

func (tree *Tree) rrotate(cur *Node) {

	tree.Count++
	tree.RotateLog += string(cur.Key.([]byte)) + ": rrotate "

	const l = 0
	const r = 1
	// 1 right 0 left
	mov := cur.Children[l]

	mov.Value, cur.Value = cur.Value, mov.Value //交换值达到, 相对位移
	mov.Key, cur.Key = cur.Key, mov.Key

	//  mov.children[l]不可能为nil
	mov.Children[l].parent = cur
	cur.Children[l] = mov.Children[l]

	// 解决mov节点孩子转移的问题
	if mov.Children[r] != nil {
		mov.Children[l] = mov.Children[r]
	} else {
		mov.Children[l] = nil
	}

	if cur.Children[r] != nil {
		mov.Children[r] = cur.Children[r]
		mov.Children[r].parent = mov
	} else {
		mov.Children[r] = nil
	}

	// 连接转移后的节点 由于mov只是与cur交换值,parent不变
	cur.Children[r] = mov

	mov.height = getMaxChildrenHeight(mov) + 1
	cur.height = getMaxChildrenHeight(cur) + 1
}

func (tree *Tree) lrotate(cur *Node) {

	tree.Count++
	tree.RotateLog += string(cur.Key.([]byte)) + ": lrotate "

	const l = 1
	const r = 0

	mov := cur.Children[l]

	mov.Value, cur.Value = cur.Value, mov.Value //交换值达到, 相对位移
	mov.Key, cur.Key = cur.Key, mov.Key

	// 不可能为nil
	mov.Children[l].parent = cur
	cur.Children[l] = mov.Children[l]

	if mov.Children[r] != nil {
		mov.Children[l] = mov.Children[r]
	} else {
		mov.Children[l] = nil
	}

	if cur.Children[r] != nil {
		mov.Children[r] = cur.Children[r]
		mov.Children[r].parent = mov
	} else {
		mov.Children[r] = nil
	}

	cur.Children[r] = mov

	mov.height = getMaxChildrenHeight(mov) + 1
	cur.height = getMaxChildrenHeight(cur) + 1
}

func getMaxAndChildrenHeight(cur *Node) (h1, h2, maxh int) {
	h1 = getHeight(cur.Children[0])
	h2 = getHeight(cur.Children[1])
	if h1 > h2 {
		maxh = h1
	} else {
		maxh = h2
	}

	return
}

func getMaxChildrenHeight(cur *Node) int {
	h1 := getHeight(cur.Children[0])
	h2 := getHeight(cur.Children[1])
	if h1 > h2 {
		return h1
	}
	return h2
}

func getHeight(cur *Node) int {
	if cur == nil {
		return -1
	}
	return cur.height
}

func (tree *Tree) fixRemoveHeight(cur *Node) {
	for {

		lefth, rigthh, lrmax := getMaxAndChildrenHeight(cur)

		// 判断当前节点是否有变化, 如果没变化的时候, 不需要往上修复
		curheight := lrmax + 1
		cur.height = curheight

		// 计算高度的差值 绝对值大于2的时候需要旋转
		diff := lefth - rigthh
		if diff < -HeightDiff {
			r := cur.Children[1] // 根据左旋转的右边节点的子节点 左右高度选择旋转的方式
			if getHeight(r.Children[0]) > getHeight(r.Children[1]) {
				tree.lrrotate(cur)
			} else {
				tree.lrotate(cur)
			}
		} else if diff > HeightDiff {
			l := cur.Children[0]
			if getHeight(l.Children[1]) > getHeight(l.Children[0]) {
				tree.rlrotate(cur)
			} else {
				tree.rrotate(cur)
			}
		} else {
			if cur.height == curheight {
				return
			}
		}

		if cur.parent == nil {
			return
		}

		cur = cur.parent
	}

}

func (tree *Tree) fixPutHeight(cur *Node) {

	for {

		lefth := getHeight(cur.Children[0])
		rigthh := getHeight(cur.Children[1])

		// 计算高度的差值 绝对值大于2的时候需要旋转
		diff := lefth - rigthh
		if diff < -HeightDiff {
			r := cur.Children[1] // 根据左旋转的右边节点的子节点 左右高度选择旋转的方式
			if getHeight(r.Children[0]) > getHeight(r.Children[1]) {
				tree.lrrotate(cur)
			} else {
				tree.lrotate(cur)
			}
		} else if diff > HeightDiff {
			l := cur.Children[0]
			if getHeight(l.Children[1]) > getHeight(l.Children[0]) {
				tree.rlrotate(cur)
			} else {
				tree.rrotate(cur)
			}

		} else {
			// 选择一个child的最大高度 + 1为 高度
			if lefth > rigthh {
				cur.height = lefth + 1
			} else {
				cur.height = rigthh + 1
			}
		}

		if cur.parent == nil || cur.height < cur.parent.height {
			return
		}
		cur = cur.parent
	}
}

func output(node *Node, prefix string, isTail bool, str *string) {

	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}

	switch k := node.Key.(type) {
	case []byte:
		*str += spew.Sprint(string(k)) + "\n"
	default:
		*str += spew.Sprint(k) + "\n"
	}

	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Children[0], newPrefix, true, str)
	}

}

func outputfordebug(node *Node, prefix string, isTail bool, str *string) {

	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		outputfordebug(node.Children[1], newPrefix, false, str)
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
		parentv = spew.Sprint(node.parent.Value)
	}
	suffix += parentv + "|" + spew.Sprint(node.height) + ")"
	*str += spew.Sprint(node.Value) + suffix + "\n"

	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		outputfordebug(node.Children[0], newPrefix, true, str)
	}
}

func (tree *Tree) debugString() string {
	if tree.size == 0 {
		return ""
	}
	str := "AVLTree\n"
	outputfordebug(tree.Root, "", true, &str)
	return str
}
