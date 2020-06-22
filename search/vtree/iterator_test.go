package vtree

import (
	"strconv"
	"strings"
	"testing"
)

func catdogdoc() *Tree {
	tree := New()
	for i := 0; i < 50; i++ {
		v := "cat-" + strconv.Itoa(i)
		tree.PutString(v, v)
	}

	for i := 0; i < 50; i++ {
		v := "dog-" + strconv.Itoa(i)
		tree.PutString(v, v)
	}

	for i := 0; i < 10; i++ {
		v := "doc-" + strconv.Itoa(i)
		tree.PutString(v, v)
	}

	return tree
}

func TestCasePrev(t *testing.T) {
	tree := catdogdoc()
	iter := tree.SeekPrefix([]byte("doc-"))

	checksize := 10
	for iter.Next() {
		checksize--
	}
	if checksize != 0 {
		t.Error("size is error")
	}

	checksize = 10
	for iter.Prev() {
		checksize--
		if !strings.HasPrefix(string(iter.Value()), "doc-") {
			t.Error("prefix error")
		}
	}

	if checksize != 0 {
		t.Error("size is error")
	}
}

func TestCaseRangePrev(t *testing.T) {
	tree := catdogdoc()
	iter := tree.SeekRange([]byte("doc-"), []byte("doc-zzz"))

	checksize := 10
	for iter.Next() {
		checksize--
	}
	if checksize != 0 {
		t.Error("size is error")
	}

	checksize = 10
	for iter.Prev() {
		checksize--
		if !strings.HasPrefix(string(iter.Value()), "doc-") {
			t.Error("prefix error")
		}
	}

	if checksize != 0 {
		t.Error("size is error")
	}

	iter = tree.SeekRange([]byte("doc-"), []byte("doc-10"))
	for iter.Next() {
		if !(string(iter.Value()) == "doc-0" || string(iter.Value()) == "doc-1") {
			t.Error("seek error")
		}
	}
}

func TestRangeCount(t *testing.T) {
	// tree := New()
	// for i := 2; i < 1000; i += 2 {
	// 	sv := strconv.Itoa(i)
	// 	tree.PutString(sv, sv)
	// }

	// for i := 0; i < 1000; i++ {
	// 	sv := strconv.Itoa(i)
	// 	nnext := tree.seekNodeNextEx([]byte(sv))
	// 	nprev := tree.seekNodePrevEx([]byte(sv))
	// 	if nnext == nil {
	// 		t.Error("nnext: nil . sv == ", sv)
	// 	}
	// 	if nprev == nil {
	// 		t.Error("nprev: nil . sv == ", sv)
	// 	}

	// 	if nnext == nil {
	// 		continue
	// 	}

	// 	if nprev == nil {
	// 		continue
	// 	}

	// 	seekv := string(nnext.Value())
	// 	iter := nnext.IteratorRange(tree)
	// 	iter.Prev()
	// 	iter.Prev()
	// 	seekvNext2Prev := string(iter.Value())

	// 	seekvPrev := string(nprev.Value())
	// 	if i%2 != 0 {
	// 		if seekvPrev != seekvNext2Prev {
	// 			t.Error(sv, "seek:", seekv, "prev:", seekvPrev, "next2prev:", seekvNext2Prev)
	// 		}
	// 	} else {
	// 		if seekv != seekvPrev {
	// 			t.Error(sv, "seek:", seekv, "prev:", seekvPrev, "next2prev:", seekvNext2Prev)
	// 		}
	// 	}

	// }

	// t.Error(tree.debugString())
}
