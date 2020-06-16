package astack

import (
	"github.com/474420502/focus/stack"
	"github.com/davecgh/go-spew/spew"
)

// type IStack interface {
// 	Clear()

// 	Empty() bool

// 	Size() uint

// 	// String 从左到右 左边第一个表示Top 如链表 a(top)->b->c
// 	String() string

// 	Values() []interface{}

// 	Push(v interface{})

// 	Pop() (interface{}, bool)

// 	Peek() (interface{}, bool)
// }

// Stack 栈
type Stack struct {
	element []interface{}
}

func assertImplementation() {
	var _ stack.IStack = (*Stack)(nil)
}

// New  创建一个Stack
func New() *Stack {
	st := &Stack{}
	return st
}

// Push 压栈
func (st *Stack) Push(v interface{}) {
	st.element = append(st.element, v)
}

// Peek 相当与栈顶
func (st *Stack) Peek() (interface{}, bool) {
	if len(st.element) == 0 {
		return nil, false
	}
	return st.element[len(st.element)-1], true
}

// Pop 出栈
func (st *Stack) Pop() (interface{}, bool) {

	if len(st.element) == 0 {
		return nil, false
	}

	last := len(st.element) - 1
	ele := st.element[last]
	st.element = st.element[0:last]
	return ele, true
}

// Clear 清空栈数据
func (st *Stack) Clear() {
	st.element = st.element[0:0]
}

// Empty 如果空栈返回true
func (st *Stack) Empty() bool {
	l := len(st.element)
	return l == 0
}

// Size 数据量
func (st *Stack) Size() uint {
	return uint(len(st.element))
}

// String 从左到右 左边第一个表示Top 如链表 a(top)->b->c 为了兼容其他栈
func (st *Stack) String() string {
	content := ""
	for i := len(st.element) - 1; i > -1; i-- {
		content += spew.Sprint(st.element[i]) + " "
	}
	return content[0 : len(content)-1]
}

// Values 同上
func (st *Stack) Values() []interface{} {
	result := make([]interface{}, len(st.element))
	ilen := len(st.element) - 1
	for i := 0; i < len(st.element); i++ {
		result[i] = st.element[ilen-i]
	}
	return result
}
