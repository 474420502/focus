package lsv

import (
	lastack "focus/stack/listarraystack"
)

type Iterator struct {
	dir    int
	up     *DNode
	cur    *DNode
	tstack *lastack.Stack
	// curnext *Node
}

func initIterator(avltree *DTree) *Iterator {
	iter := &Iterator{tstack: lastack.New()}
	iter.up = avltree.root
	return iter
}

func NewIterator(n *DNode) *Iterator {
	iter := &Iterator{tstack: lastack.New()}
	iter.up = n
	return iter
}

func NewIteratorWithCap(n *DNode, cap int) *Iterator {
	iter := &Iterator{tstack: lastack.NewWithCap(cap)}
	iter.up = n
	return iter
}

func (iter *Iterator) GetNode() *DNode {
	return iter.cur
}

func (iter *Iterator) ToHead() {
	if iter.cur == nil {
		iter.cur = iter.up
	}

	for iter.cur.family[0] != nil {
		iter.cur = iter.cur.family[0]
	}

	for iter.cur.family[1] != nil {
		iter.cur = iter.cur.family[1]
	}
	iter.SetNode(iter.cur)
	iter.cur = nil
}

func (iter *Iterator) ToTail() {

	if iter.cur == nil {
		iter.cur = iter.up
	}

	for iter.cur.family[0] != nil {
		iter.cur = iter.cur.family[0]
	}

	for iter.cur.family[2] != nil {
		iter.cur = iter.cur.family[2]
	}
	iter.SetNode(iter.cur)
	iter.cur = nil
}

func (iter *Iterator) SetNode(n *DNode) {
	iter.up = n
	iter.dir = 0
	iter.tstack.Clear()
}

func (iter *Iterator) Key() interface{} {
	return iter.cur.key
}

func (iter *Iterator) Value() interface{} {
	return iter.cur.value
}

func (iter *Iterator) GetNext(cur *DNode, idx int) *DNode {

	// iter := NewIterator(cur)
	iter.SetNode(cur)
	iter.curPushNextStack(iter.up)
	iter.up = iter.getNextUp(iter.up)

	for i := 0; i < idx; i++ {

		if iter.tstack.Size() == 0 {
			if iter.up == nil {
				return nil
			}
			iter.tstack.Push(iter.up)
			iter.up = iter.getNextUp(iter.up)
		}

		if v, ok := iter.tstack.Pop(); ok {
			iter.cur = v.(*DNode)
			if i == idx-1 {
				return iter.cur
			}
			iter.curPushNextStack(iter.cur)
		} else {
			return nil
		}
	}

	return cur
}

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
		iter.cur = v.(*DNode)
		iter.curPushNextStack(iter.cur)
		return true
	}

	return false
}
func (iter *Iterator) GetPrev(cur *DNode, idx int) *DNode {

	// iter := NewIterator(cur)
	iter.SetNode(cur)
	iter.curPushPrevStack(iter.up)
	iter.up = iter.getPrevUp(iter.up)

	for i := 0; i < idx; i++ {

		if iter.tstack.Size() == 0 {
			if iter.up == nil {
				return nil
			}
			iter.tstack.Push(iter.up)
			iter.up = iter.getPrevUp(iter.up)
		}

		if v, ok := iter.tstack.Pop(); ok {
			iter.cur = v.(*DNode)
			if i == idx-1 {
				return iter.cur
			}
			iter.curPushPrevStack(iter.cur)
		} else {
			return nil
		}
	}

	return cur
}

func (iter *Iterator) Prev() (result bool) {

	if iter.dir < 1 { // 非 1(next 方向定义 -1 为 prev)
		if iter.dir == -1 && iter.cur != nil { // 如果上次为prev方向, 则清空辅助计算的栈
			iter.tstack.Clear()
			iter.curPushPrevStack(iter.cur)    // 把当前cur计算的回朔
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
		iter.cur = v.(*DNode)
		iter.curPushPrevStack(iter.cur)
		return true
	}

	// 如果再次计算的栈为空, 则只能返回false
	return false
}

func getRelationship(cur *DNode) int {
	if cur.family[0].family[2] == cur {
		return 2
	}
	return 1
}

func (iter *Iterator) getPrevUp(cur *DNode) *DNode {
	for cur.family[0] != nil {
		if getRelationship(cur) == 1 { // next 在 降序 小值. 如果child在右边, parent 比 child 小, parent才有效, 符合降序
			return cur.family[0]
		}
		cur = cur.family[0]
	}
	return nil
}

func (iter *Iterator) curPushPrevStack(cur *DNode) {
	Prev := cur.family[1] // 当前的左然后向右找, 找到最大, 就是最接近cur 并且小于cur的值

	if Prev != nil {
		iter.tstack.Push(Prev)
		for Prev.family[2] != nil {
			Prev = Prev.family[2]
			iter.tstack.Push(Prev) // 入栈 用于回溯
		}
	}
}

func (iter *Iterator) getNextUp(cur *DNode) *DNode {
	for cur.family[0] != nil {
		if getRelationship(cur) == 0 { // Prev 在 降序 大值. 如果child在左边, parent 比 child 大, parent才有效 , 符合降序
			return cur.family[0]
		}
		cur = cur.family[0]
	}
	return nil
}

func (iter *Iterator) curPushNextStack(cur *DNode) {
	next := cur.family[2]

	if next != nil {
		iter.tstack.Push(next)
		for next.family[1] != nil {
			next = next.family[1]
			iter.tstack.Push(next)
		}
	}
}
