package liststack

import (
	"github.com/474420502/focus/stack"
	"github.com/davecgh/go-spew/spew"
)

type Node struct {
	value interface{}
	down  *Node
}

type Stack struct {
	top  *Node
	size uint
}

func assertImplementation() {
	var _ stack.IStack = (*Stack)(nil)
}

func New() *Stack {
	s := &Stack{}
	s.size = 0
	return s
}

func (as *Stack) Clear() {
	as.size = 0
	as.top = nil
}

func (as *Stack) Empty() bool {
	return as.size == 0
}

func (as *Stack) Size() uint {
	return as.size
}

// String 从左到右 左边第一个表示Top 如链表 a(top)->b->c
func (as *Stack) String() string {
	content := ""
	cur := as.top
	for ; cur != nil; cur = cur.down {
		content += spew.Sprint(cur.value) + " "
	}

	if len(content) > 0 {
		content = content[0 : len(content)-1]
	} else {
		content = ""
	}

	return content
}

func (as *Stack) Values() []interface{} {

	if as.size == 0 {
		return nil
	}

	result := make([]interface{}, as.size, as.size)

	cur := as.top
	n := 0
	for ; cur != nil; cur = cur.down {
		result[n] = cur.value
		n++
	}

	return result
}

func (as *Stack) Push(v interface{}) {
	nv := &Node{value: v}
	nv.down = as.top
	as.top = nv
	as.size++
}

func (as *Stack) Pop() (interface{}, bool) {
	if as.size == 0 {
		return nil, false
	}

	as.size--

	result := as.top
	as.top = as.top.down
	result.down = nil
	return result.value, true
}

func (as *Stack) Peek() (interface{}, bool) {
	if as.size == 0 {
		return nil, false
	}
	return as.top.value, true
}
