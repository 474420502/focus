package treeset

import (
	"fmt"
	"strings"

	"474420502.top/eson/structure/avldup"
	"474420502.top/eson/structure/compare"
)

// TreeSet
type TreeSet struct {
	tree *avldup.Tree
}

// New
func New(Compare compare.Compare) *TreeSet {
	return &TreeSet{tree: avldup.New(Compare)}
}

// Add
func (set *TreeSet) Add(items ...interface{}) {
	for _, item := range items {
		set.tree.Put(item)
	}
}

// Remove
func (set *TreeSet) Remove(items ...interface{}) {
	for _, item := range items {
		set.tree.Remove(item)
	}
}

// Values
func (set *TreeSet) Values() []interface{} {
	return set.tree.Values()
}

// Contains
func (set *TreeSet) Contains(item interface{}) bool {
	if _, ok := set.tree.Get(item); ok {
		return true
	}
	return false
}

// Contains the result is [r1,r2], not [r1, r2)
func (set *TreeSet) GetRange(r1, r2 interface{}) (result []interface{}) {
	return set.tree.GetRange(r1, r2)
}

// Contains the result is [r1,item,r2] r1->item->r2 are close-knit
func (set *TreeSet) GetAround(item interface{}) (result [3]interface{}) {
	return set.tree.GetAround(item)
}

// Empty
func (set *TreeSet) Empty() bool {
	return set.Size() == 0
}

// Clear
func (set *TreeSet) Clear() {
	set.tree.Clear()
}

// Size
func (set *TreeSet) Size() int {
	return set.tree.Size()
}

// String
func (set *TreeSet) String() string {
	content := "HashSet\n"
	items := []string{}

	set.tree.Traversal(func(k interface{}) bool {
		items = append(items, fmt.Sprintf("%v", k))
		return true
	})

	content += strings.Join(items, ", ")
	return content
}
