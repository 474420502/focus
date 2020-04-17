package stack

type IStack interface {
	Clear()

	Empty() bool

	Size() uint

	// String 从左到右 左边第一个表示Top 如链表 a(top)->b->c
	String() string

	Values() []interface{}

	Push(interface{})

	Pop() (interface{}, bool)

	Peek() (interface{}, bool)
}
