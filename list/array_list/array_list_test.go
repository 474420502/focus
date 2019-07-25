package arraylist

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestIterator(t *testing.T) {
	l := New()

	for i := 0; i < 5; i++ {
		l.Push(i)
	}

	iter := l.Iterator()

	var result []int
	for iter.Next() {
		result = append(result, iter.Value().(int))
	}

	if spew.Sprint(result) != "[0 1 2 3 4]" {
		t.Error(result)
	}

	iter = l.Iterator()
	result = nil
	for iter.Prev() {
		result = append(result, iter.Value().(int))
	}

	if spew.Sprint(result) != "[4 3 2 1 0]" {
		t.Error(result)
	}

	citer := l.CircularIterator()
	result = nil
	for i := 0; i < 11; i++ {
		if citer.Next() {
			result = append(result, citer.Value().(int))
		}
	}

	if len(result) != 11 {
		t.Error("len(result) != 11, is ", len(result))
	}

	if spew.Sprint(result) != "[0 1 2 3 4 0 1 2 3 4 0]" {
		t.Error(result)
	}

	citer = l.CircularIterator()
	result = nil
	for i := 0; i < 11; i++ {
		if citer.Prev() {
			result = append(result, citer.Value().(int))
		}
	}

	if len(result) != 11 {
		t.Error("len(result) != 11, is ", len(result))
	}

	if spew.Sprint(result) != "[4 3 2 1 0 4 3 2 1 0 4]" {
		t.Error(result)
	}
}

func TestPush(t *testing.T) {
	l := New()

	for i := 0; i < 2; i++ {
		l.PushFront(1)
	}
	var result string
	result = spew.Sprint(l.Values())
	if result != "[1 1]" {
		t.Error(result)
	}

	for i := 0; i < 2; i++ {
		l.PushBack(2)
	}
	result = spew.Sprint(l.Values())
	if result != "[1 1 2 2]" {
		t.Error(result)
	}

	l.Push(3)
	result = spew.Sprint(l.Values())
	if result != "[1 1 2 2 3]" {
		t.Error(result)
	}

}

func TestGrowth(t *testing.T) {
	l := New()
	for i := 0; i < 5; i++ {
		l.PushFront(1)
	}

	var result string
	result = spew.Sprint(l.Values())
	if result != "[1 1 1 1 1]" {
		t.Error(result)
	}

	l = New()
	for i := 0; i < 7; i++ {
		l.PushBack(1)
	}

	result = spew.Sprint(l.Values())
	if result != "[1 1 1 1 1 1 1]" {
		t.Error(result)
	}

	// for i := 0; i < 2; i++ {
	// 	l.PushBack(2)
	// }
	// result = spew.Sprint(l.Values())
	// if result != "[1 1 2 2]" {
	// 	t.Error(result)
	// }
}

func TestPop(t *testing.T) {
	l := New()
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	for i := 4; i >= 0; i-- {
		if v, ok := l.PopFront(); ok {
			if v != i {
				t.Error("should be ", v)
			}
		} else {
			t.Error("should be ok, value is", v)
		}
	}

	if v, ok := l.PopFront(); ok {
		t.Error("should not be ok, v = ", v)
	}

	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	for i := 0; i < 5; i++ {
		if v, ok := l.PopBack(); ok {
			if v != i {
				t.Error("should be ", v)
			}
		} else {
			t.Error("should be ok, value is", v)
		}
	}

	if v, ok := l.PopBack(); ok {
		t.Error("should not be ok, v = ", v)
	}
}

func TestRemove(t *testing.T) {
	l := New()
	for i := 0; i < 5; i++ {
		l.PushFront(uint(i))
	}

	var result string

	for _, selval := range []uint{4, 3} {
		last, _ := l.Index((int)(selval))
		if v, isfound := l.Remove((int)(selval)); isfound {
			if v != last {
				t.Error(v, " != ", last)
			}
		} else {
			t.Error("should be found")
		}
	}

	result = spew.Sprint(l.Values())
	if result != "[4 3 2]" {
		t.Error("should be [4 3 2], value =", result)
	}

	v, _ := l.Remove(1)
	if v != uint(3) {
		t.Error(v)
	}

	v, _ = l.Remove(1)
	if v != uint(2) {
		t.Error(v)
	}

	v, _ = l.Remove(1)
	if v != nil && l.Size() != 1 {
		t.Error(v)
	}

	v, _ = l.Remove(0)
	if v != uint(4) && l.Size() != 0 {
		t.Error(v, "size = ", l.Size())
	}

}

func TestTraversal(t *testing.T) {
	l := New()
	for i := 0; i < 5; i++ {
		l.PushFront(uint(i))
	}

	var result []interface{}

	l.Traversal(func(v interface{}) bool {
		result = append(result, v)
		return true
	})

	if spew.Sprint(result) != "[4 3 2 1 0]" {
		t.Error(result)
	}

	l.PushBack(7, 8)
	result = nil
	l.Traversal(func(v interface{}) bool {
		result = append(result, v)
		return true
	})

	if spew.Sprint(result) != "[4 3 2 1 0 7 8]" {
		t.Error(result)
	}
}

func TestRemain(t *testing.T) {
	l := New()
	for i := 0; i < 10; i++ {
		l.Push(i)
		if !l.Contains(i) {
			t.Error("Contains", i)
		}
	}

	if l.String() != "[0 1 2 3 4 5 6 7 8 9]" {
		t.Error(l.String())
	}

	for i := 10; i < 100; i++ {
		l.Push(i)
	}

	for !l.Empty() {
		l.PopBack()
	}

	for i := 10; i < 100; i++ {
		l.Push(i)
	}

	l.Clear()

	if l.Size() != 0 {
		t.Error("Size != 0")
	}
}

// func loadTestData() []int {
// 	data, err := ioutil.ReadFile("../../l.log")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	var l []int
// 	decoder := gob.NewDecoder(bytes.NewReader(data))
// 	decoder.Decode(&l)
// 	return l
// }

// func BenchmarkPush(b *testing.B) {
// 	l := loadTestData()
// 	b.N = len(l)

// 	arr := New()

// 	for i := 0; i < b.N; i++ {
// 		arr.PushBack(l[i])
// 	}
// }
