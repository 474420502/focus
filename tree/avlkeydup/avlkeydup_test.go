package avlkeydup

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"strconv"
	"testing"

	"github.com/474420502/focus/compare"
	"github.com/davecgh/go-spew/spew"
	"github.com/emirpasic/gods/trees/avltree"
)

func loadTestData() []int {
	data, err := ioutil.ReadFile("../../l.log")
	if err != nil {
		log.Println(err)
	}
	var l []int
	decoder := gob.NewDecoder(bytes.NewReader(data))
	decoder.Decode(&l)
	return l
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

func TestGetAround(t *testing.T) {
	tree := New(compare.Int)
	for _, v := range []int{7, 14, 14, 14, 16, 17, 20, 30, 21, 40, 50, 3, 40, 40, 40, 15} {
		tree.Put(v, v)
	}

	var Result string

	Result = spew.Sprint(tree.GetAround(14))
	if Result != "[7 14 15]" {
		t.Error(tree.Values())
		t.Error("17 is root, tree.GetAround(14)) is error", Result)
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
	if Result != "[30 40 50]" {
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
// 				tree.Put(v, v)
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
		tree.Put(v, v)
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
// 	for c := 0; c < 5000; c++ {
// 		tree := New(compare.Int)
// 		gods := avltree.NewWithIntComparator()
// 		var l []int
// 		m := make(map[int]int)

// 		for i := 0; len(l) < 100; i++ {
// 			v := randomdata.Number(0, 100000)
// 			if _, ok := m[v]; !ok {
// 				m[v] = v
// 				l = append(l, v)
// 				tree.Put(v, v)
// 				gods.Put(v, v)
// 			}
// 		}

// 		for i := 0; i < 100; i++ {
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
// 				tree.Put(v, v)
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
// 	tree := New(utils.IntComparator)

// 	l := loadTestData()

// 	for _, v := range l {
// 		tree.Put(v, v)
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
// 	tree := New(utils.IntComparator)

// 	l := loadTestData()

// 	for _, v := range l {
// 		tree.Put(v, v)
// 	}

// 	ll := tree.Values()
// 	b.N = len(ll)
// 	b.ResetTimer()
// 	b.StartTimer()

// 	for i := 0; i < len(ll); i++ {
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

// 	ll := tree.Values()
// 	b.N = len(ll)
// 	b.ResetTimer()
// 	b.StartTimer()

// 	for i := 0; i < len(ll); i++ {
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

// 	ll := tree.Values()
// 	b.N = len(ll)

// 	b.ResetTimer()
// 	b.StartTimer()

// 	for i := 0; i < len(ll); i++ {
// 		tree.Remove(l[i])
// 	}
// }

// func BenchmarkGet(b *testing.B) {

// 	tree := New(compare.Int)

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

// func BenchmarkGodsRBGet(b *testing.B) {
// 	tree := redblacktree.NewWithIntComparator()

// 	l := loadTestData()
// 	b.N = len(l)

// 	b.ResetTimer()
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		tree.Get(l[i])
// 	}
// }

// func BenchmarkGodsAvlGet(b *testing.B) {
// 	tree := avltree.NewWithIntComparator()

// 	l := loadTestData()
// 	b.N = len(l)

// 	b.ResetTimer()
// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		tree.Get(l[i])
// 	}
// }

func BenchmarkPut(b *testing.B) {

	d := loadTestData()
	var l [][]byte
	for _, v := range d {
		l = append(l, []byte(strconv.Itoa(v)))
	}

	b.ResetTimer()
	b.StartTimer()

	b.N = len(l)
	tree := New(compare.ByteArray)

	for _, v := range l {
		tree.Put(v, v)
	}

	// b.Log(tree.count)
}

func BenchmarkGodsAVLPut(b *testing.B) {

	d := loadTestData()
	var l [][]byte
	for _, v := range d {
		l = append(l, []byte(strconv.Itoa(v)))
	}

	b.ResetTimer()
	b.StartTimer()

	b.N = len(l)
	tree := avltree.NewWith(compare.ByteArray)

	for _, v := range l {
		tree.Put(v, v)
	}

	// b.Log(tree.count)
}

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
