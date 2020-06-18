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
	tree := catdogdoc()
	t.Error(tree.debugString())
}
