package lastack

import "testing"

func TestBase(t *testing.T) {
	s := New()
	if !s.Empty() {
		t.Error("stack not empty")
	}

	if s.Size() != 0 {
		t.Error("size != 0")
	}

	if s.Values() != nil {
		t.Error(s.Values())
	}
}

func TestPush(t *testing.T) {
	s := New()

	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	if s.Empty() {
		t.Error("stack is empty")
	}

	if s.Size() == 0 {
		t.Error("size == 0")
	}

	if s.Values() == nil {
		t.Error("Values() != nil")
	}

	if s.String() != "4 3 2 1 0" {
		t.Error(s.String())
	}

	if v, ok := s.Peek(); ok {
		if v != 4 {
			t.Error("why top != 4")
		}
	} else {
		t.Error("not ok")
	}

	if v, ok := s.Pop(); ok {
		if v != 4 {
			t.Error("why top != 4")
		}
	} else {
		t.Error("not ok")
	}

	if s.Size() != 4 {
		t.Error("pop a element, size: 5 - 1 = 4")
	}

	//
	if v, ok := s.Pop(); ok {
		if v != 3 {
			t.Error("why top != 3")
		}
	} else {
		t.Error("not ok")
	}

	if s.Size() != 3 {
		t.Error("pop a element, size: 4 - 1 = 3")
	}

	for _, ok := s.Pop(); ok != false; _, ok = s.Pop() {

	}

	if !s.Empty() && s.Size() != 0 {
		t.Error("pop all, stack should be empty")
	}

	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	if s.Size() != 5 {
		t.Error("size != 5")
	}

	s.Clear()
	if !s.Empty() && s.Size() != 0 {
		t.Error("pop all, stack should be empty")
	}

	if v, ok := s.Peek(); v != nil || ok != false {
		t.Error("should be v == nil and ok == false")
	}
}

// func BenchmarkPush(b *testing.B) {
// 	s := New()
// 	b.N = 200000
// 	for i := 0; i < b.N; i++ {
// 		v := randomdata.Number(0, 65535)
// 		s.Push(v)
// 	}
// }

// func BenchmarkGodsPush(b *testing.B) {
// 	s := arraystack.New()
// 	b.N = 200000
// 	for i := 0; i < b.N; i++ {
// 		v := randomdata.Number(0, 65535)
// 		s.Push(v)
// 	}
// }

// func BenchmarkPop(b *testing.B) {
// 	s := New()
// 	b.N = 200000

// 	for i := 0; i < b.N; i++ {
// 		v := randomdata.Number(0, 65535)
// 		s.Push(v)
// 	}

// 	b.ResetTimer()
// 	b.StartTimer()

// 	for i := 0; i < b.N; i++ {
// 		s.Pop()
// 	}
// }

// func BenchmarkGodsPop(b *testing.B) {
// 	s := arraystack.New()
// 	b.N = 200000

// 	for i := 0; i < b.N; i++ {
// 		v := randomdata.Number(0, 65535)
// 		s.Push(v)
// 	}

// 	b.ResetTimer()
// 	b.StartTimer()

// 	for i := 0; i < b.N; i++ {
// 		s.Pop()
// 	}
// }

// func BenchmarkValues(b *testing.B) {
// 	s := New()
// 	for i := 0; i < b.N; i++ {
// 		v := randomdata.Number(0, 65535)
// 		s.Push(v)
// 	}

// 	b.ResetTimer()
// 	b.StartTimer()

// 	for i := 0; i < b.N; i++ {
// 		s.Values()
// 	}
// }

// func BenchmarkGodsValues(b *testing.B) {
// 	s := arraystack.New()
// 	for i := 0; i < b.N; i++ {
// 		v := randomdata.Number(0, 65535)
// 		s.Push(v)
// 	}

// 	b.ResetTimer()
// 	b.StartTimer()

// 	for i := 0; i < b.N; i++ {
// 		s.Values()
// 	}
// }
