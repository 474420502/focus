package tree

// IBSTreeKey Compare函数可以自定义所以key不一定value, 可以是value结构体中一个属性
type IBSTreeKey interface {
	String() string
	Size() int
	Remove(key interface{}) (interface{}, bool)
	Clear()
	// Values 返回先序遍历的值
	Values() []interface{}
	Get(key interface{}) (interface{}, bool)
	Put(key, value interface{})
	Traversal(every func(k, v interface{}) bool, traversalMethod ...interface{})
}

// IBSTree
type IBSTree interface {
	String() string
	Size() int
	Remove(key interface{}) (interface{}, bool)
	Clear()
	// Values 返回先序遍历的值
	Values() []interface{}
	Get(key interface{}) (interface{}, bool)
	Put(value interface{})
	Traversal(every func(v interface{}) bool, traversalMethod ...interface{})
}
