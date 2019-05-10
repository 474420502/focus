package linkedlist

type Node struct {
	next  *Node
	value interface{}
}

func (node *Node) Value() interface{} {
	return node.value
}

type LinkedList struct {
	head *Node
	size uint
}

func New() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) Size() uint {
	return l.size
}

func (l *LinkedList) Push(v interface{}) {
	l.size++
	if l.head == nil {
		l.head = &Node{value: v}
		return
	}
	l.head = &Node{value: v, next: l.head}
}

func (l *LinkedList) PushNode(n *Node) {
	l.size++
	if l.head == nil {
		l.head = n
		return
	}

	n.next = l.head
	l.head = n
}

func (l *LinkedList) Pop() (result interface{}, found bool) {
	if n, ok := l.PopNode(); ok {
		return n.value, ok
	}
	return nil, false
}

func (l *LinkedList) PopNode() (result *Node, found bool) {
	if l.head == nil {
		return nil, false
	}

	result = l.head
	found = true
	l.head = result.next
	result.next = nil
	l.size--
	return
}

func (l *LinkedList) Remove(idx uint) (result *Node, found bool) {
	if l.size == 0 {
		return nil, false
	}

	if idx == 0 {
		result = l.head
		found = true
		l.head = result.next
		result.next = nil
		l.size--
		return
	}

	for cur := l.head; cur.next != nil; cur = cur.next {
		if idx == 1 {
			l.size--
			result = cur.next
			found = true
			cur.next = result.next
			result.next = nil
			return
		}
		idx--
	}

	return nil, false
}

func (l *LinkedList) Values() (result []interface{}) {
	l.Traversal(func(cur *Node) bool {
		result = append(result, cur.value)
		return true
	})
	return
}

func (l *LinkedList) Traversal(every func(*Node) bool) {
	for cur := l.head; cur != nil; cur = cur.next {
		if !every(cur) {
			break
		}
	}
}
