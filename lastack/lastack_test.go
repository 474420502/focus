package lastack

import (
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/emirpasic/gods/stacks/arraystack"

	"github.com/Pallinder/go-randomdata"
)

func TestPush(t *testing.T) {
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

}

func BenchmarkGet(b *testing.B) {
	s := New()
	b.N = 20000000

	for i := 0; i < b.N; i++ {
		v := randomdata.Number(0, 65535)
		s.Push(v)
	}

	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s.Get(i)
	}
}

func BenchmarkPush(b *testing.B) {
	s := New()
	b.N = 20000000
	for i := 0; i < b.N; i++ {
		v := randomdata.Number(0, 65535)
		s.Push(v)
	}
}

func BenchmarkGodsPush(b *testing.B) {
	s := arraystack.New()
	b.N = 2000000
	for i := 0; i < b.N; i++ {
		v := randomdata.Number(0, 65535)
		s.Push(v)
	}
}

func BenchmarkPop(b *testing.B) {
	s := New()
	b.N = 2000000

	for i := 0; i < b.N; i++ {
		v := randomdata.Number(0, 65535)
		s.Push(v)
	}

	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func BenchmarkGodsPop(b *testing.B) {
	s := arraystack.New()
	b.N = 2000000

	for i := 0; i < b.N; i++ {
		v := randomdata.Number(0, 65535)
		s.Push(v)
	}

	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func BenchmarkValues(b *testing.B) {
	s := New()
	for i := 0; i < b.N; i++ {
		v := randomdata.Number(0, 65535)
		s.Push(v)
	}

	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s.Values()
	}
}

func BenchmarkGodsValues(b *testing.B) {
	s := arraystack.New()
	for i := 0; i < b.N; i++ {
		v := randomdata.Number(0, 65535)
		s.Push(v)
	}

	b.ResetTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		s.Values()
	}
}
