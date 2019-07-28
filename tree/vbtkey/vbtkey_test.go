package vbtkey

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/474420502/focus/compare"
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

func TestIndexRange(t *testing.T) {
	tree := New(compare.Int)
	l := []int{7, 14, 14, 14, 16, 17, 20, 30, 21, 40, 50, 3, 40, 40, 40, 15}
	for _, v := range l {
		tree.Put(v, v)
	}
	// [3 7 14 14 14 15 16 17 20 21 30 40 40 40 40 50]
	// t.Error(tree.Values(), tree.Size())

	var result string
	result = spew.Sprint(tree.IndexRange(0, 5))
	if result != "[3 7 14 14 14 15] true" {
		t.Error(result)
	}

	result = spew.Sprint(tree.IndexRange(2, 5))
	if result != "[14 14 14 15] true" {
		t.Error(result)
	}

	result = spew.Sprint(tree.IndexRange(10, 100))
	if result != "[30 40 40 40 40 50] false" {
		t.Error(result)
	}

	result = spew.Sprint(tree.IndexRange(15, 0)) // size = 16, index max = 15
	if result != "[50 40 40 40 40 30 21 20 17 16 15 14 14 14 7 3] true" {
		t.Error(result)
	}

	result = spew.Sprint(tree.IndexRange(16, 0)) // size = 16, index max = 15
	if result != "[50 40 40 40 40 30 21 20 17 16 15 14 14 14 7 3] false" {
		t.Error(result)
	}

	result = spew.Sprint(tree.IndexRange(5, 1)) // size = 16, index max = 15
	if result != "[15 14 14 14 7] true" {
		t.Error(result)
	}

	result = spew.Sprint(tree.IndexRange(-1, -5)) // size = 16, index max = 15
	if result != "[50 40 40 40 40] true" {
		t.Error(result)
	}

	result = spew.Sprint(tree.IndexRange(-1, -16)) // size = 16, index max = 0 - 15 (-1,-16)
	if result != "[50 40 40 40 40 30 21 20 17 16 15 14 14 14 7 3] true" {
		t.Error(result)
	}

	result = spew.Sprint(tree.IndexRange(-1, -17)) // size = 16, index max = 0 - 15 (-1,-16)
	if result != "[50 40 40 40 40 30 21 20 17 16 15 14 14 14 7 3] false" {
		t.Error(result)
	}

	result = spew.Sprint(tree.IndexRange(-5, -1)) // size = 16, index max = 0 - 15 (-1,-16)
	if result != "[40 40 40 40 50] true" {
		t.Error(result)
	}
}

func TestGetAround(t *testing.T) {
	tree := New(compare.Int)
	for _, v := range []int{7, 14, 14, 14, 16, 17, 20, 30, 21, 40, 50, 3, 40, 40, 40, 15} {
		tree.Put(v, v)
	}

	var Result string

	Result = spew.Sprint(tree.GetAround(17))
	if Result != "[16 17 20]" {
		t.Error(tree.Values())
		t.Error("17 is root, tree.GetAround(17)) is error", Result)
		t.Error(tree.debugString())
	}

	Result = spew.Sprint(tree.GetAround(3))
	if Result != "[<nil> 3 7]" {
		t.Error(tree.Values())
		t.Error("tree.GetAround(3)) is error", Result)
		t.Error(tree.debugString())
	}

	Result = spew.Sprint(tree.GetAround(40))
	if Result != "[30 40 40]" {
		t.Error(tree.Values())
		t.Error("tree.GetAround(40)) is error", Result)
		t.Error(tree.debugString())
	}

	Result = spew.Sprint(tree.GetAround(50))
	if Result != "[40 50 <nil>]" {
		t.Error(tree.Values())
		t.Error("tree.GetAround(50)) is error", Result)
		t.Error(tree.debugString())
	}

	Result = spew.Sprint(tree.GetAround(18))
	if Result != "[17 <nil> 20]" {
		t.Error(tree.Values())
		t.Error("18 is not in list, tree.GetAround(18)) is error", Result)
		t.Error(tree.debugString())
	}

	Result = spew.Sprint(tree.GetAround(5))
	if Result != "[3 <nil> 7]" {
		t.Error(tree.Values())
		t.Error("5 is not in list, tree.GetAround(5)) is error", Result)
		t.Error(tree.debugString())
	}

	Result = spew.Sprint(tree.GetAround(2))
	if Result != "[<nil> <nil> 3]" {
		t.Error(tree.Values())
		t.Error("2 is not in list, tree.GetAround(2)) is error", Result)
		t.Error(tree.debugString())
	}

	Result = spew.Sprint(tree.GetAround(100))
	if Result != "[50 <nil> <nil>]" {
		t.Error(tree.Values())
		t.Error("50 is not in list, tree.GetAround(50)) is error", Result)
		t.Error(tree.debugString())
	}

}

// // for test error case

// func TestPutComparatorRandom(t *testing.T) {

// 	for n := 0; n < 300000; n++ {
// 		tree := New(compare.Int)
// 		godsavl := avltree.NewWithIntComparator()

// 		content := ""
// 		m := make(map[int]int)
// 		for i := 0; len(m) < 10; i++ {
// 			v := randomdata.Number(0, 65535)
// 			if _, ok := m[v]; !ok {
// 				m[v] = v
// 				content += spew.Sprint(v) + ","
// 				tree.Put(v, v)
// 				godsavl.Put(v, v)
// 			}
// 		}

// 		s1 := spew.Sprint(tree.Values())
// 		s2 := spew.Sprint(godsavl.Values())

// 		if s1 != s2 {
// 			t.Error(godsavl.String())
// 			t.Error(tree.debugString())
// 			t.Error(content, n)
// 			break
// 		}
// 	}
// }

func TestGet(t *testing.T) {
	tree := New(compare.Int)
	for _, v := range []int{2383, 7666, 3055, 39016, 57092, 27897, 36513, 1562, 22574, 23202} {
		tree.Put(v, v)
	}

	for _, v := range []int{2383, 7666, 3055, 39016, 57092, 27897, 36513, 1562, 22574, 23202} {
		v, ok := tree.Get(v)
		if !ok {
			t.Error("the val not found ", v)
		}
	}

	if v, ok := tree.Get(10000); ok {
		t.Error("the val(1000) is not in tree, but is found", v)
	}
}

func TestGetRange(t *testing.T) {
	tree := New(compare.Int)
	for _, v := range []int{5, 6, 8, 10, 13, 17, 1, 2, 40, 30} {
		tree.Put(v, v)
	}

	// t.Error(tree.debugString())
	// t.Error(tree.getArountNode(20))
	// t.Error(tree.Values())

	result := tree.GetRange(0, 20)
	if spew.Sprint(result) != "[1 2 5 6 8 10 13 17]" {
		t.Error(result)
	}

	result = tree.GetRange(-5, -1)
	if spew.Sprint(result) != "[]" {
		t.Error(result)
	}

	result = tree.GetRange(7, 20)
	if spew.Sprint(result) != "[8 10 13 17]" {
		t.Error(result)
	}

	result = tree.GetRange(30, 40)
	if spew.Sprint(result) != "[30 40]" {
		t.Error(result)
	}

	result = tree.GetRange(30, 60)
	if spew.Sprint(result) != "[30 40]" {
		t.Error(result)
	}

	result = tree.GetRange(40, 40)
	if spew.Sprint(result) != "[40]" {
		t.Error(result)
	}

	result = tree.GetRange(50, 60)
	if spew.Sprint(result) != "[]" {
		t.Error(result)
	}

	result = tree.GetRange(50, 1)
	if spew.Sprint(result) != "[40 30 17 13 10 8 6 5 2 1]" {
		t.Error(result)
	}

	result = tree.GetRange(30, 20)
	if spew.Sprint(result) != "[30]" {
		t.Error(result)
	}

}

func TestTravalsal(t *testing.T) {
	tree := New(compare.Int)
	for _, v := range []int{5, 6, 8, 10, 13, 17, 1, 2, 40, 30} {
		tree.Put(v, v)
	}

	i := 0
	var result []interface{}
	tree.Traversal(func(k, v interface{}) bool {
		result = append(result, k)
		i++
		if i >= 10 {
			return false
		}
		return true
	})

	if spew.Sprint(result) != "[1 2 5 6 8 10 13 17 30 40]" {
		t.Error(result)
	}

}

// func TestRemoveAll(t *testing.T) {
// ALL:
// 	for c := 0; c < 5000; c++ {
// 		tree := New(compare.Int)
// 		gods := avltree.NewWithIntComparator()
// 		var l []int
// 		m := make(map[int]int)

// 		for i := 0; len(l) < 50; i++ {
// 			v := randomdata.Number(0, 100000)
// 			if _, ok := m[v]; !ok {
// 				m[v] = v
// 				l = append(l, v)
// 				tree.Put(v, v)
// 				gods.Put(v, v)
// 			}
// 		}

// 		for i := 0; i < 50; i++ {

// 			tree.Remove(l[i])
// 			gods.Remove(l[i])

// 			s1 := spew.Sprint(tree.Values())
// 			s2 := spew.Sprint(gods.Values())
// 			if s1 != s2 {
// 				t.Error("avl remove error", "avlsize = ", tree.Size())
// 				t.Error(tree.root, i, l[i])
// 				t.Error(s1)
// 				t.Error(s2)
// 				break ALL
// 			}
// 		}
// 	}
// }

// func TestRemove(t *testing.T) {

// ALL:
// 	for N := 0; N < 5000; N++ {
// 		tree := New(compare.Int)
// 		gods := avltree.NewWithIntComparator()

// 		var l []int
// 		m := make(map[int]int)

// 		for i := 0; len(l) < 20; i++ {
// 			v := randomdata.Number(0, 100)
// 			if _, ok := m[v]; !ok {
// 				l = append(l, v)
// 				m[v] = v
// 				tree.Put(v, v)
// 				gods.Put(v, v)
// 			}
// 		}

// 		src1 := tree.String()
// 		src2 := gods.String()

// 		for i := 0; i < 20; i++ {
// 			tree.Remove(l[i])
// 			gods.Remove(l[i])
// 			if tree.root != nil && spew.Sprint(gods.Values()) != spew.Sprint(tree.Values()) {
// 				t.Error(src1)
// 				t.Error(src2)
// 				t.Error(tree.debugString())
// 				t.Error(gods.String())
// 				t.Error(l[i])
// 				break ALL
// 			}
// 		}
// 	}
// }

// func BenchmarkGetRange(b *testing.B) {

// }

// func BenchmarkIndexRange(b *testing.B) {
// 	tree := New(compare.Int)
// 	l := loadTestData()
// 	b.N = len(l)

// 	for _, v := range l {
// 		tree.Put(v, v)
// 	}

// 	b.ResetTimer()
// 	b.StartTimer()

// 	execCount := 5
// 	b.N = len(l) * execCount

// 	for i := 0; i < execCount; i++ {
// 		for range l {
// 			tree.IndexRange(i, i+49)
// 		}
// 	}
// }

// func BenchmarkSkipListGet(b *testing.B) {
// 	sl := skiplist.New(skiplist.Int)
// 	l := loadTestData()
// 	b.N = len(l)

// 	for _, v := range l {
// 		sl.Set(v, v)
// 	}

// 	b.ResetTimer()
// 	b.StartTimer()

// 	execCount := 5
// 	b.N = len(l) * execCount

// 	for i := 0; i < execCount; i++ {
// 		for _, v := range l {
// 			sl.Get(v)
// 			// e := sl.Get(v)
// 			// var result [50]interface{}
// 			// for i := 0; i < 50 && e != nil; i++ {
// 			// 	result[i] = e.Value
// 			// 	e = e.Next()
// 			// }
// 		}
// 	}
// }

// func BenchmarkSkipListSet(b *testing.B) {

// 	l := loadTestData()
// 	execCount := 1
// 	b.N = len(l) * execCount
// 	for i := 0; i < execCount; i++ {
// 		sl := skiplist.New(skiplist.Int)
// 		for _, v := range l {
// 			sl.Set(v, v)
// 		}
// 	}
// }

// func BenchmarkIterator(b *testing.B) {
// 	tree := New(compare.Int)

// 	l := loadTestData()
// 	b.N = len(l)

// 	for _, v := range l {
// 		tree.Put(v, v)
// 	}

// 	b.ResetTimer()
// 	b.StartTimer()
// 	iter := tree.Iterator()
// 	b.N = 0
// 	for iter.Next() {
// 		b.N++
// 	}
// 	for iter.Prev() {
// 		b.N++
// 	}
// 	for iter.Next() {
// 		b.N++
// 	}
// 	b.Log(b.N, len(l))
// }

// func BenchmarkRemove(b *testing.B) {
// 	tree := New(compare.Int)

// 	l := loadTestData()

// 	b.N = len(l)
// 	for _, v := range l {
// 		tree.Put(v, v)
// 	}

// 	b.ResetTimer()
// 	b.StartTimer()

// 	for i := 0; i < len(l); i++ {
// 		tree.Remove(l[i])
// 	}
// }

// func BenchmarkGodsRemove(b *testing.B) {
// 	tree := avltree.NewWithIntComparator()

// 	l := loadTestData()

// 	b.N = len(l)
// 	for _, v := range l {
// 		tree.Put(v, v)
// 	}

// 	b.ResetTimer()
// 	b.StartTimer()

// 	for i := 0; i < len(l); i++ {
// 		tree.Remove(l[i])
// 	}
// }

// func BenchmarkGodsRBRemove(b *testing.B) {
// 	tree := redblacktree.NewWithIntComparator()

// 	l := loadTestData()

// 	b.N = len(l)
// 	for _, v := range l {
// 		tree.Put(v, v)
// 	}

// 	b.ResetTimer()
// 	b.StartTimer()

// 	for i := 0; i < len(l); i++ {
// 		tree.Remove(l[i])
// 	}
// }

// func BenchmarkGet(b *testing.B) {

// 	tree := New(compare.Int)

// 	l := loadTestData()
// 	b.N = len(l)
// 	for i := 0; i < b.N; i++ {
// 		tree.Put(l[i], i)
// 	}

// 	b.ResetTimer()
// 	b.StartTimer()

// 	execCount := 10
// 	b.N = len(l) * execCount

// 	for i := 0; i < execCount; i++ {
// 		for _, v := range l {
// 			tree.Get(v)
// 		}
// 	}
// }

// func BenchmarkGodsRBGet(b *testing.B) {
// 	tree := redblacktree.NewWithIntComparator()

// 	l := loadTestData()
// 	b.N = len(l)
// 	for i := 0; i < b.N; i++ {
// 		tree.Put(l[i], i)
// 	}

// 	b.ResetTimer()
// 	b.StartTimer()

// 	execCount := 10
// 	b.N = len(l) * execCount

// 	for i := 0; i < execCount; i++ {
// 		for _, v := range l {
// 			tree.Get(v)
// 		}
// 	}
// }

// func BenchmarkGodsAvlGet(b *testing.B) {
// 	tree := avltree.NewWithIntComparator()

// 	l := loadTestData()
// 	b.N = len(l)
// 	for i := 0; i < b.N; i++ {
// 		tree.Put(l[i], i)
// 	}

// 	b.ResetTimer()
// 	b.StartTimer()

// 	execCount := 10
// 	b.N = len(l) * execCount

// 	for i := 0; i < execCount; i++ {
// 		for _, v := range l {
// 			tree.Get(v)
// 		}
// 	}
// }

// func BenchmarkPut(b *testing.B) {
// 	l := loadTestData()

// 	b.ResetTimer()
// 	b.StartTimer()

// 	execCount := 50
// 	b.N = len(l) * execCount
// 	for i := 0; i < execCount; i++ {
// 		tree := New(compare.Int)
// 		for _, v := range l {
// 			tree.Put(v, v)
// 		}
// 	}
// }

// func TestPutStable(t *testing.T) {

// }

// func BenchmarkGodsRBPut(b *testing.B) {
// 	l := loadTestData()

// 	b.ResetTimer()
// 	b.StartTimer()

// 	execCount := 50
// 	b.N = len(l) * execCount
// 	for i := 0; i < execCount; i++ {
// 		tree := redblacktree.NewWithIntComparator()
// 		for _, v := range l {
// 			tree.Put(v, v)
// 		}
// 	}
// }

// func BenchmarkGodsPut(b *testing.B) {
// 	tree := avltree.NewWithIntComparator()

// 	l := loadTestData()

// 	b.ResetTimer()
// 	b.StartTimer()

// 	b.N = len(l)
// 	for _, v := range l {
// 		tree.Put(v, v)
// 	}
// }
