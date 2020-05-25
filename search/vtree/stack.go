package vtree

type sNode struct {
}

type stack struct {
	elements []*Node
	cur      int
	size     int
}

// func assertImplementation() {
// 	// var _ stack.IStack = (*stack)(nil)
// }

func (as *stack) grow() {
	if len(as.elements) == as.size {
		temp := make([]*Node, as.size<<1)
		copy(temp, as.elements)
		as.elements = temp
	}
}

func newStack() *stack {
	s := &stack{}
	s.elements = make([]*Node, 8)
	s.cur = -1
	s.size = 0
	return s
}

func (as *stack) Clear() {
	as.cur = 0
	as.size = 0
}

func (as *stack) Empty() bool {
	return as.size == 0
}

func (as *stack) Size() int {
	return as.size
}

// func (as *stack) Values() []interface{} {
// 	result := make([]interface{}, as.size, as.size)

// 	cur := as.top
// 	n := 0
// 	for ; cur != nil; cur = cur.down {
// 		for i := range cur.elements {
// 			if cur.cur >= i {
// 				result[n] = cur.elements[cur.cur-i]
// 				n++
// 			}
// 		}
// 	}

// 	return result
// }

// func (as *stack) Index(idx int) (interface{}, bool) {
// 	if idx < 0 {
// 		return nil, false
// 	}

// 	cur := as.top
// 	for cur != nil && idx-cur.cur > 0 {
// 		idx = idx - cur.cur - 1
// 		cur = cur.down
// 	}

// 	if cur == nil {
// 		return nil, false
// 	}

// 	return cur.elements[cur.cur-idx], true
// }

func (as *stack) Push(v *Node) {
	as.grow()
	as.cur++
	as.size++
	as.elements[as.cur] = v
}

func (as *stack) Pop() (*Node, bool) {
	if as.size == 0 {
		return nil, false
	}
	epop := as.elements[as.cur]
	as.cur--
	as.size--
	return epop, true
}

func (as *stack) Peek() (*Node, bool) {
	if as.size == 0 {
		return nil, false
	}
	return as.elements[as.cur], true
}
