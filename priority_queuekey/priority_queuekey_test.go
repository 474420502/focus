package pqueuekey

import (
	"log"
	"testing"

	"github.com/474420502/focus/compare"
	"github.com/davecgh/go-spew/spew"
)

func TestQueuePush(t *testing.T) {
	pq := New(compare.Int)
	for _, v := range []int{32, 10, 53, 78, 90, 1, 4} {
		pq.Push(v, v)
		if v, ok := pq.Top(); ok {
		} else {
			t.Error(v)
		}
	}

	if v, ok := pq.Top(); ok {
		if v != 90 {
			t.Error(v)
		}
	} else {
		t.Error(v)
	}

}

func TestQueuePop(t *testing.T) {
	pq := New(compare.Int)
	for _, v := range []int{32, 10, 53, 78, 90, 1, 4} {
		pq.Push(v, v)
		if v, ok := pq.Top(); ok {
		} else {
			t.Error(v)
		}
	}

	l := []int{90, 78, 53, 32, 10, 4, 1}
	for _, lv := range l {
		if v, ok := pq.Pop(); ok {
			if v != lv {
				t.Error(v)
			}
		} else {
			t.Error(v)
		}
	}

	if v, ok := pq.Pop(); ok {
		t.Error(v)
	}
}

func TestQueueGet(t *testing.T) {
	pq := New(compare.Int)
	l := []int{32, 10, 53, 78, 90, 1, 4}
	for _, v := range l {
		pq.Push(v, v)
	}

	if v, ok := pq.Get(0); ok {
		t.Error(v)
	}

	if v, ok := pq.Get(70); ok {
		t.Error(v)
	}

	for _, v := range l {
		if gv, ok := pq.Get(v); ok {
			if gv != v {
				t.Error("Get value is error, value is", gv)
			}
		}
	}

}

func TestQueueGetRange(t *testing.T) {
	pq := New(compare.Int)
	l := []int{32, 10, 53, 78, 90, 1, 4}
	for _, v := range l {
		pq.Push(v, v)
	}

	var result string
	result = spew.Sprint(pq.GetRange(10, 40))
	if result != "[10 32]" {
		t.Error(result)
	}

	result = spew.Sprint(pq.GetRange(1, 90))
	if result != "[1 4 10 32 53 78 90]" {
		t.Error(result)
	}

	result = spew.Sprint(pq.GetRange(0, 90))
	if result != "[1 4 10 32 53 78 90]" {
		t.Error(result)
	}

	result = spew.Sprint(pq.GetRange(1, 100))
	if result != "[1 4 10 32 53 78 90]" {
		t.Error(result)
	}

	result = spew.Sprint(pq.GetRange(5, 88))
	if result != "[10 32 53 78]" {
		t.Error(result)
	}
}

func TestQueueGetAround(t *testing.T) {
	pq := New(compare.Int)
	l := []int{32, 10, 53, 78, 90, 1, 4}
	for _, v := range l {
		pq.Push(v, v)
	}

	var result string
	result = spew.Sprint(pq.GetAround(53))
	if result != "[78 53 32]" {
		t.Error(result)
	}

	result = spew.Sprint(pq.GetAround(52))
	if result != "[53 <nil> 32]" {
		t.Error(result)
	}

	result = spew.Sprint(pq.GetAround(1))
	if result != "[4 1 <nil>]" {
		t.Error(result)
	}

	result = spew.Sprint(pq.GetAround(90))
	if result != "[<nil> 90 78]" {
		t.Error(result)
	}

	result = spew.Sprint(pq.GetAround(0))
	if result != "[1 <nil> <nil>]" {
		t.Error(result)
	}

	result = spew.Sprint(pq.GetAround(100))
	if result != "[<nil> <nil> 90]" {
		t.Error(result)
	}
}

func TestQueueRemove(t *testing.T) {
	pq := New(compare.Int)
	l := []int{32, 10, 53, 78, 90, 1, 4}
	for _, v := range l {
		pq.Push(v, v)
	}

	content := ""
	for _, v := range l {
		pq.Remove(v)
		content += spew.Sprint(pq.Values())
	}

	if content != "[90 78 53 10 4 1][90 78 53 4 1][90 78 4 1][90 4 1][4 1][4][]" {
		t.Error(content)
	}
}

func TestQueueRemoveIndex(t *testing.T) {
	pq := New(compare.Int)
	l := []int{32, 10, 53, 78, 90, 1, 4}
	for _, v := range l {
		pq.Push(v, v)
	}

	content := ""
	for range l {
		pq.RemoveIndex(0)
		content += spew.Sprint(pq.Values())
	}

	if content != "[78 53 32 10 4 1][53 32 10 4 1][32 10 4 1][10 4 1][4 1][1][]" {
		t.Error(content)
	}

	if n, ok := pq.RemoveIndex(0); ok {
		t.Error("pq is not exist elements", n)
	}

}

func TestQueueIndex(t *testing.T) {
	pq := New(compare.Int)
	for _, v := range []int{32, 10, 53, 78, 90, 1, 4} {
		pq.Push(v, v)
	}

	l := []int{90, 78, 53, 32, 10, 4, 1}
	for i, lv := range l {

		if v, ok := pq.Index(len(l) - i - 1); ok {
			if v != l[len(l)-i-1] {
				t.Error(v)
			}
		} else {
			t.Error(i, "index is not exist")
		}

		if v, ok := pq.Index(i); ok {
			if v != lv {
				t.Error(v)
			}
		} else {
			t.Error(i, "index is not exist")
		}
	}

	if v, ok := pq.Index(-1); ok {
		if v != 1 {
			t.Error(v)
		}
	} else {
		t.Error("-1 index is not exist")
	}

	if v, ok := pq.Index(pq.Size()); ok {
		t.Error("index is exits", pq.Size(), v)
	}

	if v, ok := pq.Index(pq.Size() - 1); !ok {
		if v != 1 {
			t.Error("the last value is 1 not is ", v)
		}
	}

	if v, ok := pq.Index(-10); ok {
		t.Error("-10 index is exits", v)
	}
}

// func BenchmarkQueueGet(b *testing.B) {

// 	l := loadTestData()

// 	pq := New(compare.Int)
// 	for _, v := range l {
// 		pq.Push(v, v)
// 	}

// 	execCount := 5
// 	b.N = len(l) * execCount

// 	b.ResetTimer()
// 	b.StartTimer()

// ALL:
// 	for i := 0; i < execCount; i++ {
// 		for _, v := range l {
// 			if gv, ok := pq.Get(v); !ok {
// 				b.Error(gv)
// 				break ALL
// 			}
// 		}
// 	}
// }

// func BenchmarkQueueRemove(b *testing.B) {
// 	l := loadTestData()

// 	pq := New(compare.Int)
// 	for _, v := range l {
// 		pq.Push(v, v)
// 	}

// 	b.N = len(l)
// 	b.ResetTimer()
// 	b.StartTimer()

// 	for _, v := range l {
// 		pq.Remove(v)
// 	}
// }

// func BenchmarkQueueIndex(b *testing.B) {

// 	l := loadTestData()

// 	pq := New(compare.Int)
// 	for _, v := range l {
// 		pq.Push(v, v)
// 	}

// 	execCount := 2
// 	b.N = len(l) * execCount

// 	b.ResetTimer()
// 	b.StartTimer()

// ALL:
// 	for i := 0; i < execCount; i++ {
// 		for idx := range l {
// 			if v, ok := pq.Index(idx); !ok {
// 				b.Error(v)
// 				break ALL
// 			}
// 		}
// 	}
// }

// func BenchmarkPriorityPush(b *testing.B) {

// 	l := loadTestData()
// 	execCount := 5
// 	b.N = len(l) * execCount

// 	b.ResetTimer()
// 	b.StartTimer()

// 	for i := 0; i < execCount; i++ {
// 		pq := New(compare.Int)
// 		for _, v := range l {
// 			pq.Push(v, v)
// 		}
// 	}
// }

// func BenchmarkPriorityPop(b *testing.B) {

// 	l := loadTestData()

// 	pq := New(compare.Int)
// 	for _, v := range l {
// 		pq.Push(v, v)
// 	}

// 	b.N = len(l)
// 	b.ResetTimer()
// 	b.StartTimer()

// 	for i := 0; i < b.N; i++ {
// 		pq.Pop()
// 	}
// }

func TestPriorityQueue_Iterator(t *testing.T) {
	pq := New(compare.Int)
	for i := 0; i < 5; i++ {
		pq.Push(i, i)
	}

	pq.Push(-1, -1)
	pq.Push(10, 10)

	result := pq.String()
	if result != "[10 4 3 2 1 0 -1]" {
		t.Error("should be [10 4 3 2 1 0 -1]")
	}

	iter := pq.Iterator()
	iter.ToHead()

	values := pq.Values()
	for i := 0; ; i++ {
		if values[i] != iter.Value() {
			t.Error(values[i], " != ", iter.Value())
		}

		if !iter.Next() {
			break
		}
	}
}

func TestPriorityQueue_Iterator2(t *testing.T) {
	pq := New(compare.Int)
	for i := 0; i < 5; i++ {
		pq.Push(i, i)
	}

	iter := pq.Iterator()
	iter.ToHead()

	n, _ := pq.IndexNode(0)
	if n.value != 4 {
		t.Error(n)
	}

	if v, _ := pq.Top(); v != 4 {
		t.Error("Top != 4, and is ", v)
	}

	if v := iter.GetNext(n, 2).value; v != 2 {
		t.Error("iter.GetNext(n, 2) != 2, and is ", v)
	}

	pq = New(compare.Int)
	for i := 100; i >= 0; i-- {
		pq.Push(i, i)
	}
	if v, _ := pq.Top(); v != 100 {
		t.Error("Top != 100, and is ", v)
	}

	for pq.Size() >= 50 {
		pq.Pop()
	}

	if v, _ := pq.Top(); v != 48 {
		t.Error("Top != 48, and is ", v)
	}

	pq = New(compare.Int)
	for i := 0; i < 100; i++ {
		pq.Push(i, i)
	}
	if v, _ := pq.Top(); v != 99 {
		t.Error("Top != 99, and is ", v)
	}

	for pq.Size() >= 50 {
		pq.Pop()
	}

	if v, _ := pq.Top(); v != 48 {
		t.Error("Top != 49, and is ", v)
	}
}

func TestMain(t *testing.T) {
	pq := New(compare.Int)
	pq.Push(1, 1)
	pq.Push(4, 4)
	pq.Push(5, 5)
	pq.Push(6, 6)
	pq.Push(2, 2) // pq.Values() = [6 5 4 2 1]
	log.Println(pq.Values())
	value, _ := pq.Pop() // value = 6
	log.Println(value)
	value, _ = pq.Get(1) // value = 1 pq.Values() = [5 4 2 1]
	log.Println(value)
	value, _ = pq.Get(0) // value = nil , Get equal to Seach Key
	log.Println(value)
	value, _ = pq.Index(0) // value = 5, compare.Int the order from big to small
	log.Println(value)
	values := pq.GetRange(2, 5) // values = [2 4 5]
	log.Println(values)
	values = pq.GetRange(5, 2) // values = [5 4 2]
	log.Println(values)
	values = pq.GetRange(100, 2) // values = [5 4 2]
	log.Println(values)
	values3 := pq.GetAround(5) // values3 = [<nil>, 5, 4]
	log.Println(values3)

	iter := pq.Iterator() // Next 大到小 从root节点起始
	log.Println(pq.String())
	// log.Println(iter.Value()) 直接使用会报错,
	iter.ToHead()
	log.Println(iter.Value())              // 起始最大值. true 5
	log.Println(iter.Prev(), iter.Value()) // false 5

	// Prev 大到小
	log.Println(iter.Next(), iter.Value()) // true 4

}
