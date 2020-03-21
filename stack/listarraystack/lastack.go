package lastack

import (
	"focus/stack"

	"github.com/davecgh/go-spew/spew"
)

type Node struct {
	elements []interface{}
	cur      int
	down     *Node
}

type Stack struct {
	top   *Node
	cache *Node
	size  uint
}

func assertImplementation() {
	var _ stack.IStack = (*Stack)(nil)
}

func (as *Stack) grow() bool {
	if as.top.cur+1 == cap(as.top.elements) {

		var grownode *Node
		if as.cache != nil {
			grownode = as.cache
			grownode.cur = -1
			as.cache = nil
		} else {
			var growsize uint
			if as.size <= 256 {
				growsize = as.size << 1
			} else {
				growsize = 256 + as.size>>2
			}
			grownode = &Node{elements: make([]interface{}, growsize, growsize), cur: -1}
		}

		grownode.down = as.top
		as.top = grownode
		return true
	}

	return false
}

func (as *Stack) cacheRemove() bool {
	if as.top.cur == 0 && as.top.down != nil {
		as.cache = as.top
		as.top = as.top.down
		as.cache.down = nil
		return true
	}

	return false
}

func New() *Stack {
	s := &Stack{}
	s.size = 0
	s.top = &Node{elements: make([]interface{}, 8, 8), cur: -1}
	return s
}

func NewWithCap(cap int) *Stack {
	s := &Stack{}
	s.size = 0
	s.top = &Node{elements: make([]interface{}, cap, cap), cur: -1}
	return s
}

func (as *Stack) Clear() {
	as.size = 0

	as.top.down = nil
	as.top.cur = -1
}

func (as *Stack) Empty() bool {
	return as.size == 0
}

func (as *Stack) Size() uint {
	return as.size
}

// String 左为Top
func (as *Stack) String() string {
	content := ""
	cur := as.top
	for ; cur != nil; cur = cur.down {
		for i, _ := range cur.elements {
			if cur.cur >= i {
				content += spew.Sprint(cur.elements[cur.cur-i]) + " "
			}
		}
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
		for i, _ := range cur.elements {
			if cur.cur >= i {
				result[n] = cur.elements[cur.cur-i]
				n++
			}
		}
	}

	return result
}

func (as *Stack) Index(idx int) (interface{}, bool) {
	if idx < 0 {
		return nil, false
	}

	cur := as.top
	for cur != nil && idx-cur.cur > 0 {
		idx = idx - cur.cur - 1
		cur = cur.down
	}

	if cur == nil {
		return nil, false
	}

	return cur.elements[cur.cur-idx], true
}

func (as *Stack) Push(v interface{}) {
	as.grow()
	as.top.cur++
	as.top.elements[as.top.cur] = v
	as.size++
}

func (as *Stack) Pop() (interface{}, bool) {
	if as.size <= 0 {
		return nil, false
	}

	as.size--
	if as.cacheRemove() {
		return as.cache.elements[as.cache.cur], true
	}

	epop := as.top.elements[as.top.cur]
	as.top.cur--
	return epop, true
}

func (as *Stack) Peek() (interface{}, bool) {
	if as.size <= 0 {
		return nil, false
	}
	return as.top.elements[as.top.cur], true
}
