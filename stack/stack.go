package lastack

import (
	"github.com/davecgh/go-spew/spew"
)

type Node struct {
	value interface{}
	down  *Node
}

type Stack struct {
	top  *Node
	size int
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

func (as *Stack) Size() int {
	return as.size
}

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
	if as.size <= 0 {
		return nil, false
	}

	as.size--

	result := as.top
	as.top = as.top.down
	result.down = nil
	return result.value, true
}

func (as *Stack) Peek() (interface{}, bool) {
	if as.size <= 0 {
		return nil, false
	}
	return as.top.value, true
}
