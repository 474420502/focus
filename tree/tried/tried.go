package tried

import "github.com/davecgh/go-spew/spew"

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

// Put the word in tried
func (tried *Tried) Put(words string) {
	cur := tried.root
	var n *Node

	bytes := []byte(words)

	for i := 0; i < len(bytes); i++ {
		w := tried.wiStore.Byte2Index(bytes[i])

		if cur.data == nil {
			cur.data = make([]*Node, tried.wiStore.DataSize)
		}

		if n = cur.data[w]; n == nil {
			n = new(Node)
			cur.data[w] = n
		}
		cur = n
	}
	cur.value = tried
}

// PutWithValue the word with value in tried.eg. you can count word in value
func (tried *Tried) PutWithValue(words string, value interface{}) {
	cur := tried.root
	var n *Node

	bytes := []byte(words)

	for i := 0; i < len(bytes); i++ {
		w := tried.wiStore.Byte2Index(bytes[i])

		if cur.data == nil {
			cur.data = make([]*Node, tried.wiStore.DataSize)
		}

		if n = cur.data[w]; n == nil {
			n = new(Node)
			cur.data[w] = n
		}
		cur = n
	}

	cur.value = value
}

func (tried *Tried) Get(words string) interface{} {
	cur := tried.root
	if cur.data == nil {
		return nil
	}

	var n *Node
	bytes := []byte(words)

	for i := 0; i < len(bytes); i++ {
		w := tried.wiStore.Byte2Index(bytes[i]) //TODO: 升级Index 函数
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

func (tried *Tried) HasPrefix(words string) bool {
	cur := tried.root
	var n *Node
	bytes := []byte(words)

	for i := 0; i < len(bytes); i++ {
		w := tried.wiStore.Byte2Index(bytes[i]) //TODO: 升级Index 函数
		if n = cur.data[w]; n == nil {
			return false
		}
		cur = n
	}
	return true
}

func (tried *Tried) PrefixWords(words string) []string {
	cur := tried.root
	var n *Node
	bytes := []byte(words)

	var header []byte
	for i := 0; i < len(bytes); i++ {
		curbyte := bytes[i]
		header = append(header, curbyte)
		w := tried.wiStore.Byte2Index(curbyte)
		if n = cur.data[w]; n == nil {
			return nil
		}
		cur = n
	}

	var result []string

	var traversal func([]byte, *Node)
	traversal = func(prefix []byte, cur *Node) {

		for i, n := range cur.data {
			if n != nil {
				nextPrefix := append(prefix, tried.wiStore.Index2Byte(uint(i)))
				traversal(nextPrefix, n)
				if n.value != nil {
					result = append(result, string(append(header, nextPrefix...)))
				}
			}
		}

	}
	// 拼接头
	if n != nil {
		if n.value != nil {
			result = append(result, string(header))
		}
		traversal([]byte{}, n)
	}

	return result
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

func (tried *Tried) WordsArray() []string {
	var result []string

	var traversal func([]byte, *Node)
	traversal = func(prefix []byte, cur *Node) {

		for i, n := range cur.data {
			if n != nil {
				nextPrefix := append(prefix, tried.wiStore.Index2Byte(uint(i)))
				traversal(nextPrefix, n)
				if n.value != nil {
					result = append(result, string(nextPrefix))
				}
			}
		}

	}

	if tried.root != nil {
		traversal([]byte{}, tried.root)
	}

	return result
}

func (tried *Tried) String() string {
	return spew.Sprint(tried.WordsArray())
}
