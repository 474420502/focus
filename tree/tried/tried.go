package tried

// func (ts TriedString) WordIndex(idx uint) uint {
// 	w := ts[idx]
// 	if w >= 'a' && w <= 'z' {
// 		return uint(w) - 'a'
// 	} else if w >= 'A' && w <= 'Z' {
// 		return uint(w) - 'A' + 26
// 	} else {
// 		return uint(w) - '0' + 52
// 	}
// }

type Tried struct {
	root    *Node
	wiStore *wordIndexStore
}

type Node struct {
	data  []*Node
	value interface{}
}

// New 默认 WordIndexLower 意味着只支持小写
func New() *Tried {
	tried := &Tried{}
	tried.root = new(Node)

	tried.wiStore = WordIndexDict[WordIndexLower]
	return tried
}

// NewWithWordType 选择单词的类型 WordIndexLower 意味着只支持小写
func NewWithWordType(t WordIndexType) *Tried {
	tried := &Tried{}
	tried.root = new(Node)

	tried.wiStore = WordIndexDict[t]

	return tried
}

func (tried *Tried) Put(words string, values ...interface{}) {
	cur := tried.root
	var n *Node

	for i := 0; i < len(words); i++ {
		w := tried.wiStore.Byte2Index(words[i])

		if cur.data == nil {
			cur.data = make([]*Node, tried.wiStore.DataSize)
		}

		if n = cur.data[w]; n == nil {
			n = new(Node)
			cur.data[w] = n
		}
		cur = n
	}

	vlen := len(values)
	switch vlen {
	case 0:
		cur.value = tried
	case 1:
		cur.value = values[0]
	case 2:
		// TODO: 执行函数 values[1] 为函数类型 func (cur *Node, value interface{}) ...可以插入, 也可以不插入
	default:
		panic("unknow select to do")
	}

}

func (tried *Tried) Get(words string) interface{} {
	cur := tried.root
	var n *Node

	for i := 0; i < len(words); i++ {
		w := tried.wiStore.Byte2Index(words[i]) //TODO: 升级Index 函数
		if n = cur.data[w]; n == nil {
			return nil
		}
		cur = n
	}
	return n.value
}

func (tried *Tried) Has(words string) bool {
	return tried.Get(words) != nil
}

func (tried *Tried) Traversal(every func(cidx uint, value interface{}) bool) {

	var traversal func(*Node)
	traversal = func(cur *Node) {
		if cur != nil {
			for i, n := range cur.data {
				if n != nil {
					if n.value != nil {
						if !every(uint(i), n.value) {
							return
						}
					}
					traversal(n)
				}
			}
		}
	}

	root := tried.root
	traversal(root)
}

// func (tried *Tried) String() []string {
// 	var result []string
// 	tried.Traversal(func(cidx uint, value interface{}) bool {
// 		result = append(result, spew.)
// 	})
// 	return result
// }
