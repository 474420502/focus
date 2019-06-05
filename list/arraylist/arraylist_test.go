package arraylist

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

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
		last, _ := l.Index(selval)
		if v, isfound := l.Remove(selval); isfound {
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
