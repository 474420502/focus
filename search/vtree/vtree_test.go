package vtree

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"testing"
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

	iter := tree.SeekRange([]byte("key-"), []byte("key-aaa"))

	var result [][]byte
	for iter.NextLimit() {
		result = append(result, iter.Value())
	}

	for _, v := range result {
		tree.Remove(v)
	}

	if len(tree.GetRange([]byte("key-"), []byte("key-aaa"))) != 0 {
		t.Error("remove is error")
	}
}

func TestRemoveRange(t *testing.T) {
	tree := New()
	for i := 0; i < 50; i++ {
		istr := strconv.Itoa(i)
		tree.Put([]byte(istr), []byte(istr))
	}

	t.Error(tree.debugString())
}
