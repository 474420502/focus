package plist

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"testing"

	"focus/compare"

	"github.com/davecgh/go-spew/spew"
)

func loadTestData() []int {
	data, err := ioutil.ReadFile("../l.log")
	if err != nil {
		log.Println(err)
	}
	var l []int
	decoder := gob.NewDecoder(bytes.NewReader(data))
	decoder.Decode(&l)
	return l
}

func TestInsert(t *testing.T) {
	pl := New(compare.Int)
	for i := 0; i < 10; i++ {
		pl.Push(i)
	}

	if pl.size != 10 {
		t.Error(pl.size)
	}

	if pl.String() != "9 8 7 6 5 4 3 2 1 0" {
		t.Error(pl.String())
	}

	if pl.RString() != "0 1 2 3 4 5 6 7 8 9" {
		t.Error(pl.RString())
	}

	for i := 0; i < 10; i++ {
		pl.Push(i)
	}

	if pl.String() != "9 9 8 8 7 7 6 6 5 5 4 4 3 3 2 2 1 1 0 0" {
		t.Error(pl.String())
	}

	if pl.RString() != "0 0 1 1 2 2 3 3 4 4 5 5 6 6 7 7 8 8 9 9" {
		t.Error(pl.RString())
	}
}

func TestIterator(t *testing.T) {
	pl := New(compare.Int)
	for i := 0; i < 10; i++ {
		pl.Push(i)
	}

	iter := pl.Iterator()

	for i := 0; iter.Next(); i++ {
		if iter.Value() != 9-i {
			t.Error("iter.Next() ", iter.Value(), "is not equal ", 9-i)
		}
	}

	if iter.cur != iter.pl.tail {
		t.Error("current point is not equal tail ", iter.pl.tail)
	}

	for i := 0; iter.Prev(); i++ {
		if iter.Value() != i {
			t.Error("iter.Prev() ", iter.Value(), "is not equal ", i)
		}
	}
}

func TestCircularIterator(t *testing.T) {
	pl := New(compare.Int)
	for i := 0; i < 10; i++ {
		pl.Push(i)
	}

	iter := pl.CircularIterator()

	for i := 0; i != 10; i++ {
		iter.Next()
		if iter.Value() != 9-i {
			t.Error("iter.Next() ", iter.Value(), "is not equal ", 9-i)
		}
	}

	if iter.cur != iter.pl.tail.prev {
		t.Error("current point is not equal tail ", iter.pl.tail.prev)
	}

	if iter.Next() {
		if iter.Value() != 9 {
			t.Error("iter.Value() != ", iter.Value())
		}
	}

	iter.ToTail()
	for i := 0; i != 10; i++ {
		iter.Prev()
		if iter.Value() != i {
			t.Error("iter.Prev() ", iter.Value(), "is not equal ", i)
		}
	}

	if iter.cur != iter.pl.head.next {
		t.Error("current point is not equal tail ", iter.pl.tail.prev)
	}

	if iter.Prev() {
		if iter.Value() != 0 {
			t.Error("iter.Value() != ", iter.Value())
		}
	}
}

func TestGet(t *testing.T) {
	pl := New(compare.Int)
	for i := 0; i < 10; i++ {
		pl.Push(i)
	}

	for _, v := range []int{0, 9, 5, 7} {
		if g, ok := pl.Index(v); ok {
			if g != (9 - v) {
				t.Error(v, "Get == ", g)
			}
		}
	}

	if n, ok := pl.Index(10); ok {
		t.Error("index 10  is over size", n)
	}
}

func TestTop(t *testing.T) {
	pl := New(compare.Int)
	for i := 0; i < 10; i++ {
		pl.Push(i)
	}

	i := 0
	for n, ok := pl.Pop(); ok; n, ok = pl.Pop() {
		if (9 - i) != n {
			t.Error("value is not equal ", i)
		}
		if top, tok := pl.Top(); tok {
			if (9 - i - 1) != top {
				t.Error("top is error cur i = ", i, "top is ", top)
			}
		}

		i++
	}

	if pl.Size() != 0 {
		t.Error("list size is not zero")
	}
}

func TestPop(t *testing.T) {
	pl := New(compare.Int)
	for i := 0; i < 10; i++ {
		pl.Push(i)
	}

	i := 0
	for n, ok := pl.Pop(); ok; n, ok = pl.Pop() {
		if (9 - i) != n {
			t.Error("value is not equal ", i)
		}
		i++
	}

	if pl.Size() != 0 {
		t.Error("list size is not zero")
	}

	for i := 9; i >= 0; i-- {
		pl.Push(i)
	}

	i = 0
	for n, ok := pl.Pop(); ok; n, ok = pl.Pop() {
		if (9 - i) != n {
			t.Error("value is not equal ", i)
		}
		i++
	}

	if pl.Size() != 0 {
		t.Error("list size is not zero")
	}
}

func TestRemove(t *testing.T) {
	pl := New(compare.Int)
	for i := 0; i < 10; i++ {
		pl.Push(i)
	}

	pl.RemoveWithIndex(0)
	if g, ok := pl.Index(0); ok {
		if g != 8 {
			t.Error(g)
		}
	}

	pl.RemoveWithIndex(-1)
	if g, ok := pl.Index(-1); ok {
		if g != 1 {
			t.Error(g)
		}
	}

}

func TestTraversal(t *testing.T) {
	l := New(compare.Int)
	for i := 0; i < 5; i++ {
		l.Push(i)
	}

	var result []interface{}

	l.Traversal(func(v interface{}) bool {
		result = append(result, v)
		return true
	})

	if spew.Sprint(result) != "[4 3 2 1 0]" {
		t.Error(result)
	}

	l.Push(7)
	l.Push(6)
	result = nil
	l.Traversal(func(v interface{}) bool {
		result = append(result, v)
		return true
	})

	if spew.Sprint(result) != "[7 6 4 3 2 1 0]" {
		t.Error(result)
	}
}

func TestContains(t *testing.T) {
	ll := New(compare.Int)
	for i := 0; i < 10; i++ {
		ll.Push(i)
	}

	for i := 0; i < 10; i++ {
		if !ll.Contains(i) {
			t.Error(i)
		}
	}

	for i := 10; i < 20; i++ {
		if ll.Contains(i) {
			t.Error(i)
		}
	}

	if spew.Sprint(ll.Values()) != "[9 8 7 6 5 4 3 2 1 0]" {
		t.Error(spew.Sprint(ll.Values()))
	}

}

// func BenchmarkGet(b *testing.B) {
// 	pl := New(compare.Int)

// 	l := loadTestData()

// 	for _, v := range l {
// 		pl.Push(v)
// 	}

// 	b.ResetTimer()
// 	b.StartTimer()
// 	b.N = len(l)

// 	for i := 0; i < b.N; i++ {
// 		if i%2 == 0 {
// 			pl.Get(i)
// 		}
// 	}

// }
// func BenchmarkInsert(b *testing.B) {

// 	l := loadTestData()

// 	b.ResetTimer()
// 	b.StartTimer()

// 	execCount := 1
// 	b.N = len(l) * execCount
// 	for i := 0; i < execCount; i++ {
// 		pl := New(compare.Int)
// 		for _, v := range l {
// 			pl.Push(v)
// 		}
// 	}
// }
