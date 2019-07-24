package list

// IList 通用接口
type IList interface {
	Push(value interface{})
	Contains(values ...interface{}) bool
	Index(idx uint) (interface{}, bool)
	Remove(idx uint) (result interface{}, isfound bool)
	Values() []interface{}

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
