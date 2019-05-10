package plist

import (
	"strings"

	"github.com/474420502/focus/compare"
	"github.com/davecgh/go-spew/spew"
)

type Node struct {
	prev, next *Node
	value      interface{}
}

type PriorityList struct {
	head, tail *Node
	size       int
	Compare    compare.Compare
}

func New(Compare compare.Compare) *PriorityList {
	pl := &PriorityList{head: &Node{}, tail: &Node{}, size: 0, Compare: Compare}
	pl.head.next = pl.tail
	pl.tail.prev = pl.head
	return pl
}

func (pl *PriorityList) String() string {
	content := ""

	cur := pl.head.next

	for ; cur != pl.tail; cur = cur.next {
		content += spew.Sprint(cur.value) + " "
	}
	content = strings.TrimRight(content, " ")
	return content
}

func (pl *PriorityList) RString() string {
	content := ""

	cur := pl.tail.prev

	for ; cur != pl.head; cur = cur.prev {
		content += spew.Sprint(cur.value) + " "
	}
	content = strings.TrimRight(content, " ")
	return content
}

func (pl *PriorityList) Iterator() *Iterator {
	return &Iterator{pl: pl, cur: pl.head}
}

func (pl *PriorityList) CircularIterator() *CircularIterator {
	return &CircularIterator{pl: pl, cur: pl.head}
}

func (pl *PriorityList) Size() int {
	return pl.size
}

func (pl *PriorityList) Push(value interface{}) {
	pl.size++
	pnode := &Node{value: value}
	if pl.size == 1 {
		pl.head.next = pnode
		pl.tail.prev = pnode
		pnode.prev = pl.head
		pnode.next = pl.tail
		return
	}

	cur := pl.head
	for ; cur.next != pl.tail; cur = cur.next {
		if pl.Compare(value, cur.next.value) > 0 {
			cnext := cur.next

			cur.next = pnode
			cnext.prev = pnode
			pnode.prev = cur
			pnode.next = cnext

			return
		}
	}

	cur.next = pnode
	pnode.prev = cur
	pnode.next = pl.tail
	pl.tail.prev = pnode
}

func (pl *PriorityList) Top() (result interface{}, ok bool) {
	if pl.size > 0 {
		return pl.head.next.value, true
	}
	return nil, false
}

func (pl *PriorityList) Pop() (result interface{}, ok bool) {
	if pl.size > 0 {
		pl.size--
		temp := pl.head.next
		temp.next.prev = pl.head
		pl.head.next = temp.next
		return temp.value, true
	}
	return nil, false
}

func (pl *PriorityList) Get(idx int) (interface{}, bool) {
	if n, ok := pl.GetNode(idx); ok {
		return n.value, true
	}
	return nil, false
}

func (pl *PriorityList) GetNode(idx int) (*Node, bool) {
	if idx >= 0 {
		cur := pl.head.next
		for i := 0; cur != pl.tail; i++ {
			if i == idx {
				return cur, true
			}
			cur = cur.next
		}
	} else {
		cur := pl.tail.prev
		for i := -1; cur != pl.head; i-- {
			if i == idx {
				return cur, true
			}
			cur = cur.prev
		}
	}
	return nil, false
}

func (pl *PriorityList) RemoveWithIndex(idx int) {
	if n, ok := pl.GetNode(idx); ok {
		pl.Remove(n)
	}
}

func (pl *PriorityList) Remove(node *Node) {

	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev

	node.prev = nil
	node.next = nil

	pl.size--
}
func (pl *PriorityList) Values() []interface{} {
	values := make([]interface{}, pl.size, pl.size)
	for i, cur := 0, pl.head.next; cur != pl.tail; i, cur = i+1, cur.next {
		values[i] = cur.value
	}
	return values
}
