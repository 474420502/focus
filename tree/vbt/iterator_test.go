package vbt

import (
	"sort"
	"testing"

	"focus/compare"
)

func TestIerator(t *testing.T) {
	tree := New(compare.Int)
	l := []int{5, 10, 100, 30, 40, 70, 45, 35, 23}
	for _, v := range l {
		tree.Put(v)
	}

	sort.Ints(l)

	iter := tree.Iterator()
	iter.ToHead()
	for i := 0; iter.Next(); i++ {

		if iter.Value() != l[i] {
			t.Error(iter.Value(), l[i])
		}
	}
	iter.ToTail()
	iter.Prev()
	for i := len(l) - 1; iter.Next(); i-- {

		if iter.Value() != l[i] {
			t.Error(iter.Value(), l[i])
		}
	}
}
