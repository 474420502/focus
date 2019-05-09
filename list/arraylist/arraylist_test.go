package arraylist

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestPush(t *testing.T) {
	l := New()
	for i := 0; i < 5; i++ {
		l.Push(i)
	}
	var result string
	result = spew.Sprint(l.Values())
	if result != "[4 3 2 1 0]" {
		t.Error(result)
	}

	l.Push(0)
	result = spew.Sprint(l.Values())
	if result != "[0 4 3 2 1 0]" {
		t.Error(result)
	}
}

func TestPop(t *testing.T) {
	l := New()
	for i := 0; i < 5; i++ {
		l.Push(i)
	}

	if v, ok := l.Pop(); ok {
		if v != 4 {
			t.Error(v)
		}
	} else {
		t.Error("Pop should ok, but is not ok")
	}

	var result string
	result = spew.Sprint(l.Values())
	if result != "[3 2 1 0]" {
		t.Error(result)
	}

	for i := 3; l.Size() != 0; i-- {
		if v, ok := l.Pop(); ok {
			if v != i {
				t.Error(i, v, "is not equals")
			}
		} else {
			t.Error("Pop should ok, but is not ok", i)
		}
	}

	l.Push(0)
	result = spew.Sprint(l.Values())
	if result != "[0]" {
		t.Error(result)
	}

	if l.Size() != 1 {
		t.Error("l.Size() == 1, but is error, size = ", l.Size())
	}
}

func TestRemove(t *testing.T) {
	l := New()
	for i := 0; i < 5; i++ {
		l.Push(i)
	}

	for i := 0; i < 5; i++ {
		if n, ok := l.Remove(0); ok {
			if n != 4-i {
				t.Error(n)
			}
		} else {
			t.Error("Pop should ok, but is not ok", i)
		}
	}

	if l.Size() != 0 {
		t.Error("l.Size() == 0, but is error, size = ", l.Size())
	}
}
