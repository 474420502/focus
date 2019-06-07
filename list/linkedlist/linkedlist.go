package linkedlist

import "fmt"

type Node struct {
	prev  *Node
	next  *Node
	value interface{}
}

func (node *Node) Value() interface{} {
	return node.value
}

type LinkedList struct {
	head *Node
	tail *Node
	size uint
}

// var nodePool *sync.Pool = &sync.Pool{
// 	New: func() interface{} {
// 		return &Node{}
// 	},
// }

func New() *LinkedList {
	l := &LinkedList{}
	l.head = &Node{}
	l.head.prev = nil

	l.tail = &Node{}
	l.tail.next = nil

	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LinkedList) Clear() {
	l.head.next = nil
	l.tail.prev = nil
	l.size = 0
}

func (l *LinkedList) Empty() bool {
	return l.size == 0
}

func (l *LinkedList) Size() uint {
	return l.size
}

func (l *LinkedList) PushFront(values ...interface{}) {

	var node *Node
	l.size += uint(len(values))
	for _, v := range values {
		node = &Node{}
		node.value = v

		hnext := l.head.next
		hnext.prev = node

		node.next = hnext
		node.prev = l.head
		l.head.next = node
	}
}

func (l *LinkedList) PushBack(values ...interface{}) {

	var node *Node
	l.size += uint(len(values))
	for _, v := range values {
		node = &Node{}
		node.value = v

		tprev := l.tail.prev
		tprev.next = node

		node.prev = tprev
		node.next = l.tail
		l.tail.prev = node
	}
}

func (l *LinkedList) PopFront() (result interface{}, found bool) {
	if l.size != 0 {
		l.size--

		temp := l.head.next
		hnext := temp.next
		hnext.prev = l.head
		l.head.next = hnext

		result = temp.value
		found = true
		return
	}
	return nil, false
}

func (l *LinkedList) PopBack() (result interface{}, found bool) {
	if l.size != 0 {
		l.size--

		temp := l.tail.prev
		tprev := temp.prev
		tprev.next = l.tail
		l.tail.prev = tprev

		result = temp.value
		found = true
		return
	}
	return nil, false
}

func (l *LinkedList) Front() (result interface{}, found bool) {
	if l.size != 0 {
		return l.head.next.value, true
	}
	return nil, false
}

func (l *LinkedList) Back() (result interface{}, found bool) {
	if l.size != 0 {
		return l.tail.prev.value, true
	}
	return nil, false
}

func (l *LinkedList) Find(every func(idx uint, value interface{}) bool) (interface{}, bool) {

	idx := uint(0)
	// 头部
	for cur := l.head.next; cur != l.tail; cur = cur.next {
		if every(idx, cur.value) {
			return cur.value, true
		}
		idx++
	}
	return nil, false
}

func (l *LinkedList) FindMany(every func(idx uint, value interface{}) int) (result []interface{}, isfound bool) {
	// the default of isfould  is  false
	idx := uint(0)
	// 头部
	for cur := l.head.next; cur != l.tail; cur = cur.next {
		j := every(idx, cur.value)
		switch {
		case j > 0:
			result = append(result, cur.value)
			isfound = true
		case j < 0:
			return result, isfound
		}
		idx++
	}
	return result, isfound
}

func (l *LinkedList) Index(idx uint) (interface{}, bool) {
	if idx >= l.size {
		return nil, false
	}

	if idx > l.size/2 {
		idx = l.size - 1 - idx
		// 尾部
		for cur := l.tail.prev; cur != l.head; cur = cur.prev {
			if idx == 0 {
				return cur.value, true
			}
			idx--
		}

	} else {
		// 头部
		for cur := l.head.next; cur != l.tail; cur = cur.next {
			if idx == 0 {
				return cur.value, true
			}
			idx--
		}
	}

	return nil, false
}

func (l *LinkedList) Insert(idx uint, values ...interface{}) {
	if idx > l.size {
		return
	}

	if idx > l.size/2 {
		idx = l.size - idx
		// 尾部
		for cur := l.tail.prev; cur != nil; cur = cur.prev {

			if idx == 0 {

				var start *Node
				var end *Node

				start = &Node{value: values[0]}
				end = start

				for _, value := range values[1:] {
					node := &Node{value: value}
					end.next = node
					node.prev = end
					end = node
				}

				cnext := cur.next

				cur.next = start
				start.prev = cur

				end.next = cnext
				cnext.prev = end

				break
			}

			idx--
		}

	} else {
		// 头部
		for cur := l.head.next; cur != nil; cur = cur.next {
			if idx == 0 {

				var start *Node
				var end *Node

				start = &Node{value: values[0]}
				end = start

				for _, value := range values[1:] {
					node := &Node{value: value}
					end.next = node
					node.prev = end
					end = node
				}

				cprev := cur.prev

				cprev.next = start
				start.prev = cprev

				end.next = cur
				cur.prev = end

				break
			}

			idx--
		}
	}

	l.size += uint(len(values))
}

func (l *LinkedList) InsertIf(every func(idx uint, value interface{}) int, values ...interface{}) {

	idx := uint(0)
	// 头部
	for cur := l.head.next; cur != nil; cur = cur.next {
		isInsert := every(idx, cur.value)
		if isInsert != 0 { // 1 为前 -1 为后 insert here(-1) ->cur-> insert here(1)
			var start *Node
			var end *Node

			start = &Node{value: values[0]}
			end = start

			for _, value := range values[1:] {
				node := &Node{value: value}
				end.next = node
				node.prev = end
				end = node
			}

			if isInsert < 0 {
				cprev := cur.prev
				cprev.next = start
				start.prev = cprev
				end.next = cur
				cur.prev = end
			} else {
				cnext := cur.next
				cnext.prev = end
				start.prev = cur
				cur.next = start
				end.next = cnext
			}

			l.size += uint(len(values))
			break
		}
	}
}

func remove(cur *Node) {
	curPrev := cur.prev
	curNext := cur.next
	curPrev.next = curNext
	curNext.prev = curPrev
	cur.prev = nil
	cur.next = nil
}

func (l *LinkedList) Remove(idx uint) (interface{}, bool) {
	if idx >= l.size {
		panic(fmt.Sprintf("out of list range, size is %d, idx is %d", l.size, idx))
	}

	l.size--
	if idx > l.size/2 {
		idx = l.size - idx // l.size - 1 - idx,  先减size
		// 尾部
		for cur := l.tail.prev; cur != l.head; cur = cur.prev {
			if idx == 0 {
				remove(cur)
				return cur.value, true
			}
			idx--
		}

	} else {
		// 头部
		for cur := l.head.next; cur != l.tail; cur = cur.next {
			if idx == 0 {
				remove(cur)
				return cur.value, true

			}
			idx--
		}
	}

	panic(fmt.Sprintf("unknown error"))
}

func (l *LinkedList) RemoveIf(every func(idx uint, value interface{}) int) (result []interface{}, isRemoved bool) {
	// 头部
	idx := uint(0)
	for cur := l.head.next; cur != l.tail; idx++ {
		j := every(idx, cur.value)
		switch {
		case j > 0:
			result = append(result, cur.value)
			isRemoved = true
			temp := cur.next
			remove(cur)
			cur = temp
			l.size--
			continue
		case j < 0:
			return
		}

		cur = cur.next
	}
	return
}

func (l *LinkedList) Values() (result []interface{}) {
	l.Traversal(func(value interface{}) bool {
		result = append(result, value)
		return true
	})
	return
}

func (l *LinkedList) Traversal(every func(interface{}) bool) {
	for cur := l.head.next; cur != l.tail; cur = cur.next {
		if !every(cur.value) {
			break
		}
	}
}