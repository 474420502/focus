package list

// IList 通用接口
type IList interface {
	Push(value interface{})
	Contains(values ...interface{}) bool
	Index(idx int) (interface{}, bool)
	Remove(idx int) (result interface{}, isfound bool)
	Values() []interface{}
	Traversal(every func(interface{}) bool)
	String() string

	Clear()
	Empty() bool
	Size() uint
}

// ILinkedList 通用接口
type ILinkedList interface {
	PushFront(values ...interface{})
	PushBack(values ...interface{})
	PopFront() (result interface{}, found bool)
	PopBack() (result interface{}, found bool)
}
