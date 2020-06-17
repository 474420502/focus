package vtree

// IteratorRange 迭代器
type IteratorRange struct {
	tree   *Tree
	dir    int
	up     *Node
	cur    *Node
	tstack *stack

	min []byte
	max []byte
}

// NewIteratorRange create iterator by *Node
func NewIteratorRange(tree *Tree, n *Node) *IteratorRange {
	iter := &IteratorRange{tstack: newStack()}
	iter.up = n
	iter.tree = tree
	return iter
}

// GetNode iter get current node
func (iter *IteratorRange) GetNode() *Node {
	return iter.cur
}

// func (iter *Iterator) SetNode(n *Node) {
// 	iter.up = n
// 	iter.dir = 0
// 	iter.tstack.Clear()
// }

// Key get iter current key
func (iter *IteratorRange) Key() []byte {
	return iter.cur.key
}

// Value get iter current value
func (iter *IteratorRange) Value() []byte {
	return iter.cur.value
}

// SetLimit 设置限制条件
func (iter *IteratorRange) SetLimit(min, max []byte) {
	iter.min = min // 可以修改为 先Seek Node节点. 然后用Node的指针地址判断是否相当. 减少PrevLimit的消耗
	iter.max = max
}

// Prev 带Prev限制的Prev()
func (iter *IteratorRange) Prev() (result bool) {
	if iter.prev() {
		if iter.tree.compartor(iter.cur.key, iter.min) == -1 { // 可以优化减少消耗
			return false
		}
		return true
	}
	return false
}

// Next 带Next限制的Next()
func (iter *IteratorRange) Next() (result bool) {
	if iter.next() {
		if iter.tree.compartor(iter.cur.key, iter.max) == 1 {
			return false
		}
		// log.Println(string(iter.cur.key), string(iter.max))
		return true
	}
	return false
}

// next 下一个 从小到大
func (iter *IteratorRange) next() (result bool) {

	if iter.dir > -1 {
		if iter.dir == 1 && iter.cur != nil {
			iter.tstack.Clear()
			iter.curPushNextStack(iter.cur)
			iter.up = iter.getNextUp(iter.cur)
		}
		iter.dir = -1
	}

	if iter.tstack.Size() == 0 {
		if iter.up == nil {
			return false
		}
		iter.tstack.Push(iter.up)
		iter.up = iter.getNextUp(iter.up)
	}

	if v, ok := iter.tstack.Pop(); ok {
		iter.cur = v
		iter.curPushNextStack(iter.cur)
		return true
	}

	return false
}

// prev 上一个 从大到小
func (iter *IteratorRange) prev() (result bool) {

	if iter.dir < 1 { // 非 1(next 方向定义 -1 为 prev)
		if iter.dir == -1 && iter.cur != nil { // 如果上次为prev方向, 则清空辅助计算的栈
			iter.tstack.Clear()
			iter.curPushPrevStack(iter.cur)    // 把当前cur计算的逆向回朔
			iter.up = iter.getPrevUp(iter.cur) // cur 寻找下个要计算up
		}
		iter.dir = 1
	}

	// 如果栈空了, 把up的递归计算入栈, 重新计算 下次的up值
	if iter.tstack.Size() == 0 {
		if iter.up == nil {
			return false
		}
		iter.tstack.Push(iter.up)
		iter.up = iter.getPrevUp(iter.up)
	}

	if v, ok := iter.tstack.Pop(); ok {
		iter.cur = v
		iter.curPushPrevStack(iter.cur)
		return true
	}

	// 如果再次计算的栈为空, 则只能返回false
	return false
}

func (iter *IteratorRange) getPrevUp(cur *Node) *Node {
	for cur.parent != nil {
		if getRelationship(cur) == 1 { // next 在 降序 小值. 如果child在右边, parent 比 child 小, parent才有效, 符合降序
			return cur.parent
		}
		cur = cur.parent
	}
	return nil
}

func (iter *IteratorRange) curPushPrevStack(cur *Node) {
	Prev := cur.children[0] // 当前的左然后向右找, 找到最大, 就是最接近cur 并且小于cur的值

	if Prev != nil {
		iter.tstack.Push(Prev)
		for Prev.children[1] != nil {
			Prev = Prev.children[1]
			iter.tstack.Push(Prev) // 入栈 用于回溯
		}
	}
}

func (iter *IteratorRange) getNextUp(cur *Node) *Node {
	for cur.parent != nil {
		if getRelationship(cur) == 0 { // Prev 在 降序 大值. 如果child在左边, parent 比 child 大, parent才有效 , 符合降序
			return cur.parent
		}
		cur = cur.parent
	}
	return nil
}

func (iter *IteratorRange) curPushNextStack(cur *Node) {
	next := cur.children[1]

	if next != nil {
		iter.tstack.Push(next)
		for next.children[0] != nil {
			next = next.children[0]
			iter.tstack.Push(next)
		}
	}
}
