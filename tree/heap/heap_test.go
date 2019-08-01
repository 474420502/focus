package heap

import (
	"sort"
	"testing"

	"github.com/474420502/focus/compare"
	"github.com/Pallinder/go-randomdata"
)

func TestHeapGrowSlimming(t *testing.T) {
	h := New(compare.Int)
	var results []int
	for i := 0; i < 100; i++ {
		v := randomdata.Number(0, 100)
		results = append(results, v)
		h.Put(v)
	}
	sort.Slice(results, func(i, j int) bool {
		if results[i] > results[j] {
			return true
		}
		return false
	})

	if h.Size() != 100 || h.Empty() {
		t.Error("size != 100")
	}

	for i := 0; !h.Empty(); i++ {
		v, _ := h.Pop()
		if results[i] != v {
			t.Error("heap is error")
		}
	}

	if h.Size() != 0 {
		t.Error("size != 0")
	}

	h.Put(1)
	h.Put(5)
	h.Put(2)

	if h.Values()[0] != 5 {
		t.Error("top is not equal to 5")
	}

	h.Clear()
	h.Reborn()

	if !h.Empty() {
		t.Error("clear reborn is error")
	}

}

func TestHeapPushTopPop(t *testing.T) {
	h := New(compare.Int)
	l := []int{9, 5, 15, 2, 3}
	ol := []int{15, 9, 5, 3, 2}
	for _, v := range l {
		h.Put(v)
	}

	for _, tv := range ol {
		if v, isfound := h.Top(); isfound {
			if !(isfound && v == tv) {
				t.Error(v)
			}
		}

		if v, isfound := h.Pop(); isfound {
			if !(isfound && v == tv) {
				t.Error(v)
			}
		}
	}

	if h.Size() != 0 {
		t.Error("heap size is not equals to zero")
	}
}

// func Int(k1, k2 interface{}) int {
// 	c1 := k1.(int)
// 	c2 := k2.(int)
// 	switch {
// 	case c1 > c2:
// 		return -1
// 	case c1 < c2:
// 		return 1
// 	default:
// 		return 0
// 	}
// }

// func TestPush(t *testing.T) {

// 	for i := 0; i < 1000000; i++ {
// 		h := New(Int)

// 		gods := binaryheap.NewWithIntComparator()
// 		for c := 0; c < 20; c++ {
// 			v := randomdata.Number(0, 100)
// 			h.Push(v)
// 			gods.Push(v)
// 		}

// 		r1 := spew.Sprint(h.Values())
// 		r2 := spew.Sprint(gods.Values())
// 		if r1 != r2 {
// 			t.Error(r1)
// 			t.Error(r2)
// 			break
// 		}
// 	}

// }

// func TestPop(t *testing.T) {

// 	for i := 0; i < 200000; i++ {
// 		h := New(Int)

// 		// m := make(map[int]int)
// 		gods := binaryheap.NewWithIntComparator()
// 		for c := 0; c < 40; c++ {
// 			v := randomdata.Number(0, 100)
// 			// if _, ok := m[v]; !ok {
// 			h.Push(v)
// 			gods.Push(v)
// 			// 	m[v] = v
// 			// }

// 		}

// 		// t.Error(h.Values())
// 		// t.Error(gods.Values())
// 		for c := 0; c < randomdata.Number(5, 10); c++ {
// 			v1, _ := h.Pop()
// 			v2, _ := gods.Pop()

// 			if v1 != v2 {
// 				t.Error(h.Values(), v1)
// 				t.Error(gods.Values(), v2)
// 				return
// 			}
// 		}

// 		r1 := spew.Sprint(h.Values())
// 		r2 := spew.Sprint(gods.Values())
// 		if r1 != r2 {
// 			t.Error(r1)
// 			t.Error(r2)
// 			break
// 		}
// 	}
// }

// func BenchmarkPush(b *testing.B) {

// 	l := loadTestData()

// 	b.ResetTimer()
// 	execCount := 50
// 	b.N = len(l) * execCount

// 	for c := 0; c < execCount; c++ {
// 		b.StopTimer()
// 		h := New(Int)
// 		b.StartTimer()
// 		for _, v := range l {
// 			h.Push(v)
// 		}
// 	}
// }

// func BenchmarkPop(b *testing.B) {

// 	h := New(Int)

// 	l := loadTestData()

// 	b.ResetTimer()
// 	execCount := 20
// 	b.N = len(l) * execCount

// 	for c := 0; c < execCount; c++ {
// 		b.StopTimer()
// 		for _, v := range l {
// 			h.Push(v)
// 		}
// 		b.StartTimer()
// 		for h.size != 0 {
// 			h.Pop()
// 		}
// 	}
// }

// func BenchmarkGodsPop(b *testing.B) {

// 	h := binaryheap.NewWithIntComparator()

// 	l := loadTestData()

// 	b.ResetTimer()
// 	execCount := 20
// 	b.N = len(l) * execCount

// 	for c := 0; c < execCount; c++ {
// 		b.StopTimer()
// 		for _, v := range l {
// 			h.Push(v)
// 		}
// 		b.StartTimer()
// 		for h.Size() != 0 {
// 			h.Pop()
// 		}
// 	}

// }

// func BenchmarkGodsPush(b *testing.B) {
// 	l := loadTestData()

// 	b.ResetTimer()
// 	execCount := 50
// 	b.N = len(l) * execCount

// 	for c := 0; c < execCount; c++ {
// 		b.StopTimer()
// 		h := binaryheap.NewWith(Int)
// 		b.StartTimer()
// 		for _, v := range l {
// 			h.Push(v)
// 		}
// 	}
// }

// func loadTestData() []int {
// 	data, err := ioutil.ReadFile("../l.log")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	var l []int
// 	decoder := gob.NewDecoder(bytes.NewReader(data))
// 	decoder.Decode(&l)
// 	return l
// }
