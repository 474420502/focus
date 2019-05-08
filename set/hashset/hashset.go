package hashset

import (
	"fmt"
	"strings"
)

var nullItem = struct{}{}

// HashSet
type HashSet struct {
	hm map[interface{}]struct{}
}

// New
func New() *HashSet {
	return &HashSet{hm: make(map[interface{}]struct{})}
}

// Add
func (set *HashSet) Add(items ...interface{}) {
	for _, item := range items {
		if _, ok := set.hm[item]; !ok {
			set.hm[item] = nullItem
		}
	}
}

// Remove
func (set *HashSet) Remove(items ...interface{}) {
	for _, item := range items {
		delete(set.hm, item)
	}
}

// Values
func (set *HashSet) Values() []interface{} {
	values := make([]interface{}, set.Size())
	count := 0
	for item := range set.hm {
		values[count] = item
		count++
	}
	return values
}

// Contains
func (set *HashSet) Contains(item interface{}) bool {
	if _, contains := set.hm[item]; contains {
		return true
	}
	return false
}

// Empty
func (set *HashSet) Empty() bool {
	return set.Size() == 0
}

// Clear
func (set *HashSet) Clear() {
	set.hm = make(map[interface{}]struct{})
}

// Size
func (set *HashSet) Size() int {
	return len(set.hm)
}

// String
func (set *HashSet) String() string {
	content := "["
	items := []string{}
	for k := range set.hm {
		items = append(items, fmt.Sprintf("%v", k))
	}
	content += strings.Join(items, ",")
	content += "]"
	return content
}
