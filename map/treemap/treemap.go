package treemap

import (
	"strings"

	"github.com/davecgh/go-spew/spew"

	"github.com/474420502/focus/compare"
	"github.com/474420502/focus/tree/avlkeydup"
)

type TreeMap struct {
	avl *avlkeydup.Tree
}

// New instantiates a hash map.
func New(Compare compare.Compare) *TreeMap {
	return &TreeMap{avl: avlkeydup.New(Compare)}
}

// Put inserts element into the map.
func (tmap *TreeMap) Put(key interface{}, value interface{}) {
	tmap.avl.Put(key, value)
}

func (tmap *TreeMap) Get(key interface{}) (value interface{}, isfound bool) {
	value, isfound = tmap.avl.Get(key)
	return
}

func (tmap *TreeMap) Remove(key interface{}) {
	tmap.Remove(key)
}

func (tmap *TreeMap) Empty() bool {
	return tmap.avl.Size() == 0
}

func (tmap *TreeMap) Size() int {
	return tmap.avl.Size()
}

func (tmap *TreeMap) Keys() []interface{} {
	keys := make([]interface{}, tmap.avl.Size())
	count := 0
	tmap.avl.Traversal(func(key, value interface{}) bool {
		keys[count] = key
		count++
		return true
	})
	return keys
}

func (tmap *TreeMap) Values() []interface{} {
	values := make([]interface{}, tmap.avl.Size())
	count := 0
	tmap.avl.Traversal(func(key, value interface{}) bool {
		values[count] = value
		count++
		return true
	})
	return values
}

func (tmap *TreeMap) Clear() {
	tmap.avl.Clear()
}

func (tmap *TreeMap) String() string {
	content := "{"
	tmap.avl.Traversal(func(key, value interface{}) bool {
		content += spew.Sprint(key) + ":" + spew.Sprint(value) + ","
		return true
	})
	content = strings.TrimRight(content, ",") + "}"
	return content
}
