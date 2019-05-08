package hashmap

import "fmt"

type HashMap struct {
	hm map[interface{}]interface{}
}

// New instantiates a hash map.
func New() *HashMap {
	return &HashMap{hm: make(map[interface{}]interface{})}
}

// Put inserts element into the map.
func (hm *HashMap) Put(key interface{}, value interface{}) {
	hm.hm[key] = value
}

func (hm *HashMap) Get(key interface{}) (value interface{}, isfound bool) {
	value, isfound = hm.hm[key]
	return
}

func (hm *HashMap) Remove(key interface{}) {
	delete(hm.hm, key)
}

func (hm *HashMap) Empty() bool {
	return len(hm.hm) == 0
}

func (hm *HashMap) Size() int {
	return len(hm.hm)
}

func (hm *HashMap) Keys() []interface{} {
	keys := make([]interface{}, len(hm.hm))
	count := 0
	for key := range hm.hm {
		keys[count] = key
		count++
	}
	return keys
}

func (hm *HashMap) Values() []interface{} {
	values := make([]interface{}, len(hm.hm))
	count := 0
	for _, value := range hm.hm {
		values[count] = value
		count++
	}
	return values
}

func (hm *HashMap) Clear() {
	hm.hm = make(map[interface{}]interface{})
}

func (hm *HashMap) String() string {
	str := fmt.Sprintf("%v", hm.hm)
	return str
}
