package linkedhashmap

import (
	"fmt"

	linkedlist "github.com/474420502/focus/list/linked_list"
	"github.com/davecgh/go-spew/spew"
)

// LinkedHashmap
type LinkedHashmap struct {
	list *linkedlist.LinkedList
	hmap map[interface{}]interface{}
}

// New
func New() *LinkedHashmap {
	lhmap := &LinkedHashmap{list: linkedlist.New(), hmap: make(map[interface{}]interface{})}
	return lhmap
}

// PushBack if key exists, push value replace the value is exists. size is unchanging
func (lhmap *LinkedHashmap) PushBack(key interface{}, value interface{}) {
	if _, ok := lhmap.hmap[key]; !ok {
		lhmap.list.PushBack(key)
	}
	lhmap.hmap[key] = value
}

// PushFront if key exists, push value replace the value is exists. size is unchanging
func (lhmap *LinkedHashmap) PushFront(key interface{}, value interface{}) {
	if _, ok := lhmap.hmap[key]; !ok {
		lhmap.list.PushFront(key)
	}
	lhmap.hmap[key] = value
}

// Insert 如果成功在该位置返回True, 否则返回false
func (lhmap *LinkedHashmap) Insert(idx uint, key interface{}, value interface{}) bool {
	if _, ok := lhmap.hmap[key]; !ok {

		return true
	}
	return false
}

// Get
func (lhmap *LinkedHashmap) Get(key interface{}) (interface{}, bool) {
	value, ok := lhmap.hmap[key]
	return value, ok
}

// Clear
func (lhmap *LinkedHashmap) Clear() {
	lhmap.list.Clear()
	lhmap.hmap = make(map[interface{}]interface{})
}

// Remove if key not exists reture nil, false.
func (lhmap *LinkedHashmap) Remove(key interface{}) (interface{}, bool) {
	if v, ok := lhmap.hmap[key]; ok {
		delete(lhmap.hmap, key)
		lhmap.list.RemoveIf(func(idx uint, lkey interface{}) linkedlist.RemoveState {
			if lkey == key {
				return linkedlist.RemoveAndBreak
			}
			return linkedlist.UnremoveAndContinue
		})
		return v, true
	}
	return nil, false
}

// RemoveIndex
func (lhmap *LinkedHashmap) RemoveIndex(idx uint) (interface{}, bool) {
	if lhmap.list.Size() <= idx {
		panic(fmt.Sprintf("out of list range, size is %d, idx is %d", lhmap.list.Size(), idx))
	}

	if _, ok := lhmap.hmap[key]; ok {
		delete(lhmap.hmap, key)
		lhmap.list.RemoveIf(func(idx uint, lkey interface{}) linkedlist.RemoveState {
			if lkey == key {
				return linkedlist.RemoveAndBreak
			}
			return linkedlist.UnremoveAndContinue
		})
	}
}

// Empty returns true if map does not contain any elements
func (lhmap *LinkedHashmap) Empty() bool {
	return lhmap.Size() == 0
}

// Size returns number of elements in the map.
func (lhmap *LinkedHashmap) Size() uint {
	return lhmap.list.Size()
}

// Keys returns all keys left to right (head to tail)
func (lhmap *LinkedHashmap) Keys() []interface{} {
	return lhmap.list.Values()
}

// Values returns all values in-order based on the key.
func (lhmap *LinkedHashmap) Values() []interface{} {
	values := make([]interface{}, lhmap.Size())
	count := 0
	lhmap.list.Traversal(func(key interface{}) bool {
		values[count] = lhmap.hmap[key]
		count++
		return true
	})
	return values
}

// String returns a string
func (lhmap *LinkedHashmap) String() string {
	return spew.Sprint(lhmap.Values())
}
