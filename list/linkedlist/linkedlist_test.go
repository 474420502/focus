package linkedlist

import (
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/davecgh/go-spew/spew"
)

func TestPushFront(t *testing.T) {
	l := New()
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}
	var result string
	result = spew.Sprint(l.Values())
	if result != "[4 3 2 1 0]" {
		t.Error(result)
	}

	l.PushFront(0)
	result = spew.Sprint(l.Values())
	if result != "[0 4 3 2 1 0]" {
		t.Error(result)
	}
}

func TestPushBack(t *testing.T) {
	l := New()
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	var result string
	result = spew.Sprint(l.Values())
	if result != "[0 1 2 3 4]" {
		t.Error(result)
	}

	l.PushBack(0)
	result = spew.Sprint(l.Values())
	if result != "[0 1 2 3 4 0]" {
		t.Error(result)
	}
}

func TestPopFront(t *testing.T) {
	l := New()
	// "[0 1 2 3 4]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}
	// var result string

	for i := 4; i >= 0; i-- {
		if v, ok := l.PopFront(); ok {
			if v != i {
				t.Error("[4 3 2 1 0] PopFront value should be ", i, ", but is ", v)
			}
		} else {
			t.Error("PopFront is not ok")
		}

		if l.Size() != uint(i) {
			t.Error("l.Size() is error, is", l.Size())
		}
	}
}

func TestPopBack(t *testing.T) {
	l := New()
	// "[0 1 2 3 4]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}
	// var result string

	for i := 0; i < 5; i++ {
		if v, ok := l.PopBack(); ok {
			if v != i {
				t.Error("[4 3 2 1 0] PopFront value should be ", i, ", but is ", v)
			}
		} else {
			t.Error("PopFront is not ok")
		}

		if l.Size() != uint(5-i-1) {
			t.Error("l.Size() is error, is", l.Size())
		}
	}

}

func TestInsert(t *testing.T) {
	l1 := New()
	l2 := New()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l1.Insert(0, i)
		l2.PushFront(i)
	}

	var result1, result2 string
	result1 = spew.Sprint(l1.Values())
	result2 = spew.Sprint(l2.Values())
	if result1 != result2 {
		t.Error(result1, result2)
	}

	for i := 0; i < 5; i++ {
		l1.Insert(l1.Size(), i)
		l2.PushBack(i)
		// t.Error(l1.Values(), l2.Values())
	}

	result1 = spew.Sprint(l1.Values())
	result2 = spew.Sprint(l2.Values())
	if result1 != result2 {
		t.Error(result1, result2)
	}

	if result1 != "[4 3 2 1 0 0 1 2 3 4]" {
		t.Error("result should be [4 3 2 1 0 0 1 2 3 4]\n but result is", result1)
	}

	l1.Insert(1, 99)
	result1 = spew.Sprint(l1.Values())
	if result1 != "[4 99 3 2 1 0 0 1 2 3 4]" {
		t.Error("[4 3 2 1 0 0 1 2 3 4] insert with index 1, should be [4 99 3 2 1 0 0 1 2 3 4]\n but result is", result1)
	}
}

func BenchmarkPushBack(b *testing.B) {

	ec := 5
	cs := 2000000
	b.N = cs * ec

	for c := 0; c < ec; c++ {
		l := New()
		for i := 0; i < cs; i++ {
			l.PushBack(i)
		}
	}
}

func BenchmarkPushFront(b *testing.B) {

	ec := 5
	cs := 2000000
	b.N = cs * ec

	for c := 0; c < ec; c++ {
		l := New()
		for i := 0; i < cs; i++ {
			l.PushFront(i)
		}
	}

}

func BenchmarkInsert(b *testing.B) {

	ec := 10
	cs := 1000
	b.N = cs * ec

	for c := 0; c < ec; c++ {
		l := New()
		for i := 0; i < cs; i++ {
			ridx := randomdata.Number(0, int(l.Size())+1)
			l.Insert(uint(ridx), i)
		}
	}
}

// func TestPop(t *testing.T) {
// 	l := New()
// 	for i := 0; i < 5; i++ {
// 		l.Push(i)
// 	}

// 	if v, ok := l.Pop(); ok {
// 		if v != 4 {
// 			t.Error(v)
// 		}
// 	} else {
// 		t.Error("Pop should ok, but is not ok")
// 	}

// 	var result string
// 	result = spew.Sprint(l.Values())
// 	if result != "[3 2 1 0]" {
// 		t.Error(result)
// 	}

// 	for i := 3; l.Size() != 0; i-- {
// 		if v, ok := l.Pop(); ok {
// 			if v != i {
// 				t.Error(i, v, "is not equals")
// 			}
// 		} else {
// 			t.Error("Pop should ok, but is not ok", i)
// 		}
// 	}

// 	l.Push(0)
// 	result = spew.Sprint(l.Values())
// 	if result != "[0]" {
// 		t.Error(result)
// 	}

// 	if l.Size() != 1 {
// 		t.Error("l.Size() == 1, but is error, size = ", l.Size())
// 	}
// }

// func TestRemove(t *testing.T) {
// 	l := New()
// 	for i := 0; i < 5; i++ {
// 		l.Push(i)
// 	}

// 	for i := 0; i < 5; i++ {
// 		l.Remove(0)
// 		if l.head != nil {
// 			if l.head.Value() != 4-i-1 {
// 				t.Error("l.head is error")
// 			}
// 		}
// 		t.Error(l.Size())
// 	}

// 	if l.Size() != 0 {
// 		t.Error("l.Size() == 0, but is error, size = ", l.Size())
// 	}
// }
