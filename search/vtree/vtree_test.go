package vtree

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/Pallinder/go-randomdata"
)

func TestSeek(t *testing.T) {
	tree := New()

	for i := 0; i < 1000; i += 2 {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}

	for i := 0; i < 100; i += 2 {
		istr := strconv.Itoa(i)
		iter := tree.Seek([]byte(istr))
		if iter.Next() {
			vstr := string(iter.Value())
			if !strings.HasPrefix(vstr, istr) {
				t.Error(vstr)
			}
		}
	}

	if v, ok := tree.Get([]byte("50")); ok {
		if string(v) != "50" {
			t.Error("value error")
		}
	} else {
		t.Error("value error")
	}

	if v, ok := tree.Get([]byte("1")); !ok {
		if string(v) == "1" {
			t.Error("value error")
		}
	} else {
		t.Error("value error")
	}
}

func TestPutSimple(t *testing.T) {
	tree := New()
	for i := 0; i < 100; i++ {
		istr := strconv.Itoa(i)
		if !tree.Put([]byte(istr), []byte(istr)) {
			t.Error("str is cover", istr)
		}
	}

	var result [][]byte

	iter := tree.SeekRange([]byte("40"), []byte("80"))
	for iter.NextLimit() {
		result = append(result, iter.Key())
	}

	for _, v := range result {
		tree.Remove(v)
	}

	if len(tree.GetRange([]byte("40"), []byte("80"))) != 0 {
		t.Error(tree.debugString())
	}
}

func TestPut(t *testing.T) {
	tree := New()

	for i := 0; i < 1000; i += 2 {
		istr := strconv.Itoa(i)
		if !tree.PutNotCover([]byte(istr), []byte(istr)) {
			t.Error("str is cover", istr)
		}
	}

	for i := 0; i < 100; i++ {
		istr := strconv.Itoa(i)
		if i%2 == 0 {
			if tree.PutNotCover([]byte(istr), []byte(istr)) {
				t.Error(i)
			}
		} else {
			if !tree.PutNotCover([]byte(istr), []byte(istr)) {
				t.Error("str is cover", istr)
			}
		}

	}

	t.Run("Method Values", func(t *testing.T) {
		tree = New()
		for i := 0; i < 10; i++ {
			istr := strconv.Itoa(i)
			tree.PutString(istr, istr)
		}
		for i, value := range tree.Values() {
			if string(value) != strconv.Itoa(i) {
				t.Error("LDR is error")
			}
		}

		var result = []string{}
		tree.Traversal(func(k, v []byte) bool {
			result = append(result, string(k))
			return true
		}, LRD)

		if len(result) != 10 {
			t.Error(tree.debugString())
		}

		sort.Strings(result)
		if fmt.Sprint(result) != "[0 1 2 3 4 5 6 7 8 9]" {
			t.Error(result)
		}

		result = []string{}
		tree.Traversal(func(k, v []byte) bool {
			result = append(result, string(k))
			return true
		}, RDL)

		if len(result) != 10 {
			t.Error(tree.debugString())
		}

	})

	t.Run("Traversal", func(t *testing.T) {
		ret, _ := testTraversal(t, RDL)
		if ret != "[9 8 7 6 5 4 3 2 1 0]" {
			t.Error(ret)
		}
	})

	t.Run("Traversal", func(t *testing.T) {
		testTraversal(t, RLD)
	})

	t.Run("Traversal", func(t *testing.T) {
		testTraversal(t, DLR)
	})

	t.Run("Traversal", func(t *testing.T) {
		testTraversal(t, LRD)
	})

	t.Run("Traversal", func(t *testing.T) {
		testTraversal(t, DRL)
	})

	t.Run("Traversal", func(t *testing.T) {
		testTraversal(t, BFSLR)
		// t.Error(ret, dret)
	})

	t.Run("Traversal", func(t *testing.T) {
		testTraversal(t, BFSRL)
	})
}

func testTraversal(t *testing.T, m TraversalMethod) (string, string) {
	tree := New()
	for i := 0; i < 10; i++ {
		istr := strconv.Itoa(i)
		tree.PutString(istr, istr)
	}

	var result = []string{}
	tree.Traversal(func(k, v []byte) bool {
		result = append(result, string(k))
		return true
	}, m)

	if len(result) != 10 {
		t.Error(tree.debugString())
	}

	ret := fmt.Sprint(result)

	sort.Strings(result)
	if fmt.Sprint(result) != "[0 1 2 3 4 5 6 7 8 9]" {
		t.Error(result)
	}

	return ret, tree.debugString()
}

func TestSeekRange(t *testing.T) {
	tree := New()
	for i := 0; i < 50; i++ {
		istr := "key-" + strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}
	for i := 0; i < 50; i++ {
		istr := "xxx-" + strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}

	for i := 0; i < 50; i++ {
		istr := "xixi-" + strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}

	iter := tree.SeekRange([]byte("key-"), []byte("key-50"))
	if iter.NextLimit() {
		ivalue := string(iter.Value())
		if ivalue != "key-0" {
			t.Error(ivalue)
		}
	}

	iter = tree.SeekRange([]byte("xxx-50"), []byte("xxx-"))
	if iter.NextLimit() {
		t.Error(string(iter.Value()))
	}

}

func TestSeekOnlyOne(t *testing.T) {

	// t.Error(tree.debugString())
	t.Run("case1", func(t *testing.T) {
		tree := New()
		for i := 0; i < 1; i += 2 {
			istr := strconv.Itoa(i)
			tree.Put([]byte(istr), []byte(istr))
		}

		iter := tree.Seek([]byte("23"))

		if iter.Prev() {
			ivalue := string(iter.Value())
			if ivalue != "0" {
				t.Error(ivalue)
			}
		} else {
			t.Error("not exists prev")
		}
	})

	t.Run("case2", func(t *testing.T) {
		tree := New()
		for i := 0; i < 4; i += 2 {
			istr := strconv.Itoa(i)
			tree.Put([]byte(istr), []byte(istr))
		}

		iter := tree.Seek([]byte("23"))

		if iter.Prev() {
			ivalue := string(iter.Value())
			if ivalue != "2" {
				t.Error(ivalue)
			}
		} else {
			t.Error("not exists prev")
		}
	})

}

func TestRemove(t *testing.T) {

	tree := New()
	// defer func() {
	// 	if err := recover(); err != nil {

	// 		t.Error(tree.debugString())
	// 		panic(err)
	// 	}
	// }()

	for i := 0; i < 50; i++ {
		istr := "key-" + strconv.Itoa(i)
		if istr == "key-11" {
			//log.Println(tree.debugString())
		}
		tree.Put([]byte(istr), []byte(istr))
		// log.Println(tree.debugString())
		// log.Println(istr)
	}
	for i := 0; i < 50; i++ {
		istr := "xxx-" + strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
		//log.Println(tree.debugString())
	}

	for i := 0; i < 50; i++ {
		istr := "xixi-" + strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
		//log.Println(tree.debugString())
	}

	iter := tree.SeekRange([]byte("key-"), []byte("key-aaa"))

	var result [][]byte
	for iter.NextLimit() {
		result = append(result, iter.Key())
	}

	for _, v := range result {

		// if string(v) == "key-11" {
		// 	log.Println(tree.debugString())
		// 	log.Println(string(v))
		// }
		// log.Println(tree.debugString())
		// log.Println("remove:", string(v))
		tree.Remove(v)
		// log.Println(tree.debugString())
		// if string(v) == "key-11" {
		// 	log.Println(tree.Get([]byte("key-11")))

		// 	log.Println(string(v))
		// }
	}

	log.Println(tree.debugString())

	result = tree.GetRange([]byte("key-"), []byte("key-aaa"))
	if len(result) != 0 {
		t.Error(tree.debugString())
		for _, v := range result {
			t.Error(string(v))
		}

		t.Error("remove is error")
	}
}

func getValues(values [][]byte) string {
	var str string = "["
	for _, v := range values {
		str += string(v) + " "
	}
	if len(str) > 1 {
		str = str[0 : len(str)-1]
	}
	str += "]"
	return str
}

func checkValues(t *testing.T, tree *Tree, strvalue string) {
	var str string = getValues(tree.Values())
	if str != strvalue {
		t.Error(tree.debugString())
		panic(errors.New(fmt.Sprint("error, should be\n", strvalue, "\nbut values is\n", str)))
	}
}

func checkSize(t *testing.T, tree *Tree, size int) {
	if tree.Size() != size {
		t.Error(tree.debugString())
		panic(errors.New(fmt.Sprint("tree size is error, should be ", size, " but size is ", tree.Size())))
	}
}

func TestRemoveRange(t *testing.T) {
	tree := New()
	for i := 0; i < 50; i++ {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
		// log.Println(istr)
		// log.Println(tree.debugString())
		// log.Println()
	}

	tree.RemoveRange([]byte("15"), []byte("31"))
	checkSize(t, tree, 50-17-2)
	checkValues(t, tree, "[0 1 10 11 12 13 14 32 33 34 35 36 37 38 39 4 40 41 42 43 44 45 46 47 48 49 5 6 7 8 9]")

	tree = New()
	for i := 0; i < 50; i++ {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}
	checkSize(t, tree, 50)
	tree.RemoveRange([]byte("15"), []byte("40"))
	checkSize(t, tree, 40-15-1-3)
	checkValues(t, tree, "[0 1 10 11 12 13 14 41 42 43 44 45 46 47 48 49 5 6 7 8 9]")
}

func TestRemoveRangeCase1(t *testing.T) {
	tree := New()
	for i := 0; i < 50; i++ {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}

	tree.RemoveRange([]byte("0"), []byte("46"))
	checkValues(t, tree, "[47 48 49 5 6 7 8 9]")
	checkSize(t, tree, 8)

	tree.RemoveRange([]byte("0"), []byte("49"))

	checkValues(t, tree, "[5 6 7 8 9]")
	checkSize(t, tree, 5)
}

func TestRemoveRangeCase2(t *testing.T) {
	tree := New()
	for i := 47; i < 50; i++ {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}

	tree.RemoveRange([]byte("0"), []byte("49"))
	checkSize(t, tree, 0)
	checkValues(t, tree, "[]")

	tree = New()
	for i := 47; i < 50; i++ {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}

	tree.RemoveRange([]byte("48"), []byte("48"))
	checkSize(t, tree, 2)
	checkValues(t, tree, "[47 49]")
}

func TestRemoveRangeCase3(t *testing.T) {
	tree := New()
	for i := 47; i < 50; i++ {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}

	tree.RemoveRange([]byte("0"), []byte("49"))
	checkSize(t, tree, 0)
	checkValues(t, tree, "[]")
}

func TestRemoveRangeCase4(t *testing.T) {
	tree := New()
	for i := 14; i < 25; i++ {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}

	checkSize(t, tree, 11)

	tree.RemoveRange([]byte("19"), []byte("24"))

	checkSize(t, tree, 11-(6))
	checkValues(t, tree, "[14 15 16 17 18]")

	tree = New() // min: 10 max: 12 rmin: 10 rmax: 11
	for i := 10; i < 12; i++ {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}

	tree.RemoveRange([]byte("10"), []byte("11"))
	checkSize(t, tree, 0)
	checkValues(t, tree, "[]")
}

func TestRemoveRangeCase5(t *testing.T) {
	tree := New() // min: 10 max: 12 rmin: 10 rmax: 11
	istr := strconv.Itoa(12)
	tree.Put([]byte(istr), []byte(istr))

	tree.RemoveRange([]byte("0"), []byte("14"))
	checkSize(t, tree, 0)
	checkValues(t, tree, "[]")

	tree.Put([]byte(istr), []byte(istr))
	tree.RemoveRange([]byte("13"), []byte("25"))
	checkSize(t, tree, 1)
	checkValues(t, tree, "[12]")

	tree.Put([]byte(istr), []byte(istr))
	tree.RemoveRange([]byte("0"), []byte("11"))
	checkSize(t, tree, 1)
	checkValues(t, tree, "[12]")
}

func TestRemoveRangeCase6(t *testing.T) {
	tree := New() // min: 10 max: 12 rmin: 10 rmax: 11
	for i := 100; i < 1000; i++ {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}

	values := tree.Values()
	tree.RemoveRange([]byte("0"), []byte("99"))
	checkSize(t, tree, 10)
	checkValues(t, tree, getValues(values[890:]))

	tree = New() // min: 10 max: 12 rmin: 10 rmax: 11
	for i := 100; i < 1000; i++ {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}

	tree.RemoveRange([]byte("1000"), []byte("5000"))
	checkSize(t, tree, 500)

	tree = New() // min: 10 max: 12 rmin: 10 rmax: 11
	for i := 100; i < 1000; i++ {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}

	tree.RemoveRange([]byte("0"), []byte("999"))
	checkSize(t, tree, 0)
	checkValues(t, tree, "[]")
}

func TestRemoveRangeCase7(t *testing.T) {
	tree := New() // min: 10 max: 12 rmin: 10 rmax: 11
	for i := 0; i < 10; i++ {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}
	tree.RemoveRange([]byte("9"), []byte("14")) // 2 - 9
	checkSize(t, tree, 2)
	checkValues(t, tree, "[0 1]")

	tree = New() // min: 10 max: 12 rmin: 10 rmax: 11
	for i := 0; i < 10; i++ {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}
	tree.RemoveRange([]byte("!"), []byte("0"))
	checkSize(t, tree, 9)
	checkValues(t, tree, "[1 2 3 4 5 6 7 8 9]")
}

func TestRemoveRangeCase8(t *testing.T) {
	tree := New() // min: 10 max: 12 rmin: 10 rmax: 11
	for i := 0; i < 10; i++ {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}

	tree.RemoveRange([]byte("0"), []byte("2"))
	checkSize(t, tree, 7)
	checkValues(t, tree, "[3 4 5 6 7 8 9]")
	tree.RemoveRange([]byte("3"), []byte("8"))
	checkSize(t, tree, 1)
	checkValues(t, tree, "[9]")
}

func TestRemoveRangeForce(t *testing.T) {
	checksize := 1000
	for ; checksize > 0; checksize-- {
		tree := New()

		var min, max int

		for min == max {
			min = randomdata.Number(0, 500)
			max = randomdata.Number(0, 500)
		}

		if min > max {
			min, max = max, min
		}

		size := max - min
		for i := min; i < max; i++ {
			istr := strconv.Itoa(i)
			tree.Put([]byte(istr), []byte(istr))
		}

		var rmin, rmax int

		rmin = randomdata.Number(min, max)
		rmax = randomdata.Number(min, max)

		if rmin > rmax {
			rmin, rmax = rmax, rmin
		}

		t.Log("min:", min, "max:", max, "rmin:", rmin, "rmax:", rmax)

		var result []int
		for i := min; i < max; i++ {
			if i >= rmin && i <= rmax {
				continue
			}
			result = append(result, i)
		}

		checkSize(t, tree, size)
		tree.RemoveRange([]byte(strconv.Itoa(rmin)), []byte(strconv.Itoa(rmax)))
		checkSize(t, tree, size-rmax+rmin-1)
		checkValues(t, tree, fmt.Sprint(result))
	}
}

func TestIndex(t *testing.T) {
	tree := New() // min: 10 max: 12 rmin: 10 rmax: 11
	for i := 0; i < 100; i++ {
		istr := strconv.Itoa(i)
		tree.PutString(istr, istr)
	}

	iter := tree.Index(50)
	if !iter.Next() {
		t.Error("50 error")
	}
	if string(iter.Key()) != "50" {
		t.Error("key = 50")
	}
	if string(iter.Value()) != "50" {
		t.Error("key = 50")
	}
}

func TestIndexNode(t *testing.T) {
	tree := New() // min: 10 max: 12 rmin: 10 rmax: 11
	for i := 0; i < 100; i++ {
		istr := strconv.Itoa(i)
		tree.PutString(istr, istr)
	}

	for i := 0; i < 100; i++ {
		n := tree.IndexNode(i)
		if string(n.Key()) != strconv.Itoa(i) {
			t.Error("error", string(n.Key()), i)
		}
	}

	for i := -1; i >= -100; i-- {
		n := tree.IndexNode(i)
		if string(n.Value()) != strconv.Itoa(100+i) {
			t.Error("error", string(n.Value()), i)
		}
	}

	for i := 0; i < 100; i++ {
		n := tree.IndexNode(i)
		key, _ := tree.IndexKey(i)
		if string(n.Key()) != string(key) {
			t.Error("error", string(key), i)
		}
	}

	for i := 0; i < 100; i++ {
		n := tree.IndexNode(i)
		v, _ := tree.IndexValue(i)
		if string(n.Key()) != string(v) {
			t.Error("error", string(v), i)
		}
	}
}
