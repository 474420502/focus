package vtree

type sNode struct {
}

type stack struct {
	elements []*Node
	cur      int
	size     int
}

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
