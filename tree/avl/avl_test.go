package avl

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"testing"

	"github.com/474420502/focus/compare"
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

func TestIteratorHeadTail(t *testing.T) {
	tree := New(compare.Int)
	for _, v := range []int{1, 2, 7, 4, 5, 6, 7, 14, 15, 20, 30, 21, 3} {
		tree.Put(v)
	}
	// ` AVLTree
	// │       ┌── 30
	// │       │   └── 21
	// │   ┌── 20
	// │   │   └── 15
	// └── 14
	// 	   │       ┌── 7
	// 	   │   ┌── 7
	// 	   │   │   └── 6
	// 	   └── 5
	// 		   │   ┌── 4
	// 		   │   │   └── 3
	// 		   └── 2
	// 			   └── 1`

	iter := tree.Iterator()
	iter.Prev()
	if iter.Value() != 14 {
		t.Error("iter.Value() != ", 14, " value =", iter.Value())
	}

	iter.ToHead()
	if iter.Value() != 1 {
		t.Error("iter.Value() != ", 14, " value =", iter.Value())
	}

	iter.ToTail()
	if iter.Value() != 30 {
		t.Error("iter.Value() != ", 30, " value =", iter.Value())
	}
}

func TestIterator(t *testing.T) {
	tree := New(compare.Int)
	for _, v := range []int{1, 2, 7, 4, 5, 6, 7, 14, 15, 20, 30, 21, 3} {
		// t.Error(v)
		tree.Put(v)

	}
	// ` AVLTree
	// │       ┌── 30
	// │       │   └── 21
	// │   ┌── 20
	// │   │   └── 15
	// └── 14
	// 	   │       ┌── 7
	// 	   │   ┌── 7
	// 	   │   │   └── 6
	// 	   └── 5
	// 		   │   ┌── 4
	// 		   │   │   └── 3
	// 		   └── 2
	// 			   └── 1`

	iter := tree.Iterator() // root start point
	l := []int{14, 15, 20, 21, 30}

	for i := 0; iter.Next(); i++ {
		if iter.Value().(int) != l[i] {
			t.Error("iter Next error", iter.Value(), l[i])
		}
	}

	iter.Next()
	if iter.Value().(int) != 30 {
		t.Error("Next == false", iter.Value(), iter.Next(), iter.Value())
	}

	l = []int{21, 20, 15, 14, 7, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; iter.Prev(); i++ { // cur is 30 next is 21
		if iter.Value().(int) != l[i] {
			t.Error(iter.Value())
		}
	}

	if iter.Prev() != false {
		t.Error("Prev is error, cur is tail, val = 1 Prev return false")
	}
	if iter.Value().(int) != 1 { // cur is 1
		t.Error("next == false", iter.Value(), iter.Prev(), iter.Value())
	}

	if iter.Next() != true && iter.Value().(int) != 2 {
		t.Error("next to prev is error")
	}
}

func TestGetRange(t *testing.T) {
	tree := New(compare.Int)
	for _, v := range []int{5, 6, 8, 10, 13, 17, 1, 2, 40, 30} {
		tree.Put(v)
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

func TestGetAround(t *testing.T) {
	tree := New(compare.Int)
	for _, v := range []int{7, 14, 14, 14, 16, 17, 20, 30, 21, 40, 50, 3, 40, 40, 40, 15} {
		tree.Put(v)
	}

	var Result string

	Result = spew.Sprint(tree.GetAround(14))
	if Result != "[7 14 14]" {
		t.Error(tree.Values())
		t.Error("14 is root, tree.GetAround(14)) is error", Result)
		t.Error(tree.debugString())
	}

	Result = spew.Sprint(tree.GetAround(17))
	if Result != "[16 17 20]" {
		t.Error(tree.Values())
		t.Error("tree.GetAround(17)) is error", Result)
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

// for test error case
func TestPutStable(t *testing.T) {
	// f, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	// log.SetOutput(f)
	// 0-1 3 | 2-3 7-8 | 4-7 12-16 | 8-15 20-32 | 16-31 33-58 l := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 18, 19, 20, 21, 22, 30, 41, 41, 41}

	// tree := New(compare.Int)
	// for i := 0; i < 10; i++ {
	// 	tree.Put(randomdata.Number(0, 100))
	// }
	// t.Error(tree.debugString())

	// t.Error(tree.debugString(), tree.TraversalBreadth(), "\n", "-----------")
}

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
// 				tree.Put(v)
// 				godsavl.Put(v, v)
// 			}
// 		}

// 		if tree.String() != godsavl.String() {
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
		tree.Put(v)
	}

	result := `
│           ┌── 57092
│       ┌── 39016
│       │   └── 36513
│   ┌── 27897
│   │   │   ┌── 23202
│   │   └── 22574
└── 7666
    │   ┌── 3055
    └── 2383
        └── 1562
`

	s1 := tree.String()
	s2 := "AVLTree" + result
	if s1 != s2 {
		t.Error(s1, s2)
	}

	for _, v := range []int{2383, 7666, 3055, 39016, 57092, 27897, 36513, 1562, 22574, 23202} {
		v, ok := tree.Get(v)
		if !ok {
			t.Error("the val not found ", v)
		}
	}

	if v, ok := tree.Get(10000); ok {
		t.Error("the val(10000) is not in tree, but is found", v)
	}

}

// func TestRemoveAll(t *testing.T) {

// ALL:
// 	for c := 0; c < 50000; c++ {
// 		tree := New(compare.Int)
// 		gods := avltree.NewWithIntComparator()
// 		var l []int
// 		m := make(map[int]int)

// 		for i := 0; len(l) < 50; i++ {
// 			v := randomdata.Number(0, 100000)
// 			if _, ok := m[v]; !ok {
// 				m[v] = v
// 				l = append(l, v)
// 				tree.Put(v)
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
// 				t.Error(s1)
// 				t.Error(s2)
// 				break ALL
// 			}
// 		}
// 	}

// }

// func TestRemove(t *testing.T) {

// ALL:
// 	for N := 0; N < 500000; N++ {
// 		tree := New(compare.Int)
// 		gods := avltree.NewWithIntComparator()

// 		var l []int
// 		m := make(map[int]int)

// 		for i := 0; len(l) < 10; i++ {
// 			v := randomdata.Number(0, 100)
// 			if _, ok := m[v]; !ok {
// 				l = append(l, v)
// 				m[v] = v
// 				tree.Put(v)
// 				gods.Put(v, v)
// 			}
// 		}

// 		src1 := tree.String()
// 		src2 := gods.String()

// 		for i := 0; i < 10; i++ {
// 			tree.Remove(l[i])
// 			gods.Remove(l[i])
// 			if spew.Sprint(gods.Values()) != spew.Sprint(tree.Values()) && tree.size != 0 {
// 				// if gods.String() != tree.String() && gods.Size() != 0 && tree.size != 0 {
// 				t.Error(src1)
// 				t.Error(src2)
// 				t.Error(tree.debugString())
// 				t.Error(gods.String())
// 				t.Error(l[i])
// 				// t.Error(tree.TraversalDepth(-1))
// 				// t.Error(gods.Values())
// 				break ALL
// 			}
// 		}
// 	}
// }

// func BenchmarkIterator(b *testing.B) {
// 	tree := New(compare.Int)

// 	l := loadTestData()

// 	for _, v := range l {
// 		tree.Put(v)
// 	}

// 	b.ResetTimer()
// 	b.StartTimer()
// 	b.N = 0
// 	iter := tree.Iterator()
// 	for iter.Next() {
// 		b.N++
// 	}
// 	for iter.Prev() {
// 		b.N++
// 	}
// 	for iter.Next() {
// 		b.N++
// 	}
// 	for iter.Prev() {
// 		b.N++
// 	}

// }

// func BenchmarkRemove(b *testing.B) {
// 	tree := New(compare.Int)

// 	l := loadTestData()

// 	b.N = len(l)
// 	for _, v := range l {
// 		tree.Put(v)
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
// 		tree.Put(l[i])
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
// 		tree.Put(l[i], l[i])
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
// 		tree.Put(l[i], l[i])
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
// 			tree.Put(v)
// 		}
// 	}
// 	// b.Log(tree.count)
// }

// func BenchmarkGodsRBPut(b *testing.B) {
// 	tree := redblacktree.NewWithIntComparator()

// 	l := loadTestData()

// 	b.ResetTimer()
// 	b.StartTimer()

// 	b.N = len(l)
// 	for _, v := range l {
// 		tree.Put(v, v)
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
