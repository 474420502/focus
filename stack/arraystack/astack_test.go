package astack

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestPushSimple(t *testing.T) {
	var result string
	s := New()

	result = spew.Sprint(s.Values())
	if result != "[]" {
		t.Error(result)
	}

	l := []int{10, 7, 3, 4, 5, 15}
	for _, v := range l {
		s.Push(v)
	}

	result = spew.Sprint(s.Values())
	if result != "[15 5 4 3 7 10]" {
		t.Error(result)
	}

	if v, ok := s.Peek(); ok {
		if v != 15 {
			t.Error(v)
		}
	}

	result = spew.Sprint(s.Values())
	if result != "[15 5 4 3 7 10]" {
		t.Error(result)
	}

	if v, ok := s.Pop(); ok {
		if v != 15 {
			t.Error(v)
		}
	}

	result = spew.Sprint(s.Values())
	if result != "[5 4 3 7 10]" {
		t.Error(result)
	}

	for s.Size() != 1 {
		s.Pop()
	}

	result = spew.Sprint(s.Values())
	if result != "[10]" {
		t.Error(result)
	}

	if v, ok := s.Pop(); ok {
		if v != 10 {
			t.Error(v)
		}
	}

	result = spew.Sprint(s.Values())
	if result != "[]" {
		t.Error(result)
	}

	for i := 0; i < 100; i++ {
		s.Push(i)
	}

	for i := 0; i < 50; i++ {
		s.Pop()
	}

	if v, _ := s.Peek(); v != 49 {
		t.Error(s.Peek())
	}
}

func TestPushPop(t *testing.T) {
	s := New()

	l := []int{10, 7, 3, 4, 5, 15}

	for _, v := range l {
		s.Push(v)
	}

	curSize := uint(len(l))
	for !s.Empty() {
		if s.Size() != curSize {
			t.Error("size error")
		}

		v1, _ := s.Pop()
		v2 := l[len(l)-1]
		if v1 != v2 {
			t.Error(v1, v2)
		}
		l = l[0 : len(l)-1]

		curSize--
	}
}

func TestBase(t *testing.T) {
	s := New()

	l := []int{10, 7, 3, 4, 5, 15}
	for _, v := range l {
		s.Push(v)
	}

	if s.String() != "15 5 4 3 7 10" {
		t.Error(s.String())
	}

	if s.Size() != 6 {
		t.Error("Size error, is", s.Size())
	}

	s.Clear()
	if !s.Empty() {
		t.Error("Size should be Empty, is Clean.")
	}

}
