package lastack

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
