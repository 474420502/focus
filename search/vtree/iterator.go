package vtree

import "log"

// Iterator 迭代器
type Iterator struct {
	tree   *Tree
	dir    int
	up     *Node
	cur    *Node
	tstack *stack
	min    []byte
	max    []byte
}

// NewIterator create iterator by *Node
func NewIterator(tree *Tree, n *Node) *Iterator {
	iter := &Iterator{tstack: newStack()}
	iter.up = n
	iter.tree = tree
	return iter
}

// GetNode iter get current node
func (iter *Iterator) GetNode() *Node {
	return iter.cur
}

// func (iter *Iterator) SetNode(n *Node) {
// 	iter.up = n
// 	iter.dir = 0
// 	iter.tstack.Clear()
// }

// Key get iter current key
func (iter *Iterator) Key() []byte {
	return iter.cur.key
}

// Value get iter current value
func (iter *Iterator) Value() []byte {
	return iter.cur.value
}

// func (iter *Iterator) GetNext(cur *Node, idx int) *Node {

// 	// iter := NewIterator(cur)
// 	iter.SetNode(cur)
// 	iter.curPushNextStack(iter.up)
// 	iter.up = iter.getNextUp(iter.up)

// 	for i := 0; i < idx; i++ {

// 		if iter.tstack.Size() == 0 {
// 			if iter.up == nil {
// 				return nil
// 			}
// 			iter.tstack.Push(iter.up)
// 			iter.up = iter.getNextUp(iter.up)
// 		}

// 		if v, ok := iter.tstack.Pop(); ok {
// 			iter.cur = v
// 			if i == idx-1 {
// 				return iter.cur
// 			}
// 			iter.curPushNextStack(iter.cur)
// 		} else {
// 			return nil
// 		}
// 	}

// 	return cur
// }

// SetLimit 设置限制条件
func (iter *Iterator) SetLimit(min, max []byte) {
	iter.min = min // 可以修改为 先Seek Node节点. 然后用Node的指针地址判断是否相当. 减少PrevLimit的消耗
	iter.max = max
}

// PrevLimit 带Prev限制的Prev()
func (iter *Iterator) PrevLimit() (result bool) {
	if iter.Prev() {
		if iter.tree.compartor(iter.cur.key, iter.min) == -1 { // 可以优化减少消耗
			return false
		}
		return true
	}
	return false
}

// NextLimit 带Next限制的Next()
func (iter *Iterator) NextLimit() (result bool) {
	if iter.Next() {
		if iter.tree.compartor(iter.cur.key, iter.max) == 1 {
			return false
		}
		log.Println(string(iter.cur.key), string(iter.max))
		return true
	}
	return false
}

// Next 下一个 从小到大
func (iter *Iterator) Next() (result bool) {

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

// func (iter *Iterator) GetPrev(cur *Node, idx int) *Node {

// 	// iter := NewIterator(cur)
// 	iter.SetNode(cur)
// 	iter.curPushPrevStack(iter.up)
// 	iter.up = iter.getPrevUp(iter.up)

// 	for i := 0; i < idx; i++ {

// 		if iter.tstack.Size() == 0 {
// 			if iter.up == nil {
// 				return nil
// 			}
// 			iter.tstack.Push(iter.up)
// 			iter.up = iter.getPrevUp(iter.up)
// 		}

// 		if v, ok := iter.tstack.Pop(); ok {
// 			iter.cur = v
// 			if i == idx-1 {
// 				return iter.cur
// 			}
// 			iter.curPushPrevStack(iter.cur)
// 		} else {
// 			return nil
// 		}
// 	}

// 	return cur
// }

// Prev 上一个 从大到小
func (iter *Iterator) Prev() (result bool) {

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

func getRelationship(cur *Node) int {
	if cur.parent.children[1] == cur {
		return 1
	}
	return 0
}

func (iter *Iterator) getPrevUp(cur *Node) *Node {
	for cur.parent != nil {
		if getRelationship(cur) == 1 { // next 在 降序 小值. 如果child在右边, parent 比 child 小, parent才有效, 符合降序
			return cur.parent
		}
		cur = cur.parent
	}
	return nil
}

func (iter *Iterator) curPushPrevStack(cur *Node) {
	Prev := cur.children[0] // 当前的左然后向右找, 找到最大, 就是最接近cur 并且小于cur的值

	if Prev != nil {
		iter.tstack.Push(Prev)
		for Prev.children[1] != nil {
			Prev = Prev.children[1]
			iter.tstack.Push(Prev) // 入栈 用于回溯
		}
	}
}

func (iter *Iterator) getNextUp(cur *Node) *Node {
	for cur.parent != nil {
		if getRelationship(cur) == 0 { // Prev 在 降序 大值. 如果child在左边, parent 比 child 大, parent才有效 , 符合降序
			return cur.parent
		}
		cur = cur.parent
	}
	return nil
}

func (iter *Iterator) curPushNextStack(cur *Node) {
	next := cur.children[1]

	if next != nil {
		iter.tstack.Push(next)
		for next.children[0] != nil {
			next = next.children[0]
			iter.tstack.Push(next)
		}
	}
}
