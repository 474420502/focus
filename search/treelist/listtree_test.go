package listtree

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"strconv"
	"testing"

	"github.com/emirpasic/gods/trees/avltree"
)

func TestCase1(t *testing.T) {
	tree := New()
	avl := avltree.NewWithIntComparator()
	for i := int64(0); i < 100; i++ {
		k := []byte(strconv.FormatInt(i, 10))
		tree.Put(k, k)
		avl.Put(int(i), int(i))
		log.Println(string(k), tree.debugString())
		// t.Error(avl.String())
	}

	// t.Error(tree.root.children[0].size)

	t.Error(tree.debugString())
	t.Error(avl.String())
	// t.Error(tree.debugString())
}

func loadTestData() []int {
	data, err := ioutil.ReadFile("../../l.log")
	if err != nil {
		log.Println(err)
	}
	var l []int
	decoder := gob.NewDecoder(bytes.NewReader(data))
	decoder.Decode(&l)
	return l
}

func BenchmarkPut(b *testing.B) {

	d := loadTestData()
	var l [][]byte
	for _, v := range d {
		l = append(l, []byte(strconv.Itoa(v)))
	}

	b.ResetTimer()
	b.StartTimer()

	b.N = len(l)
	tree := New()

	for _, v := range l {
		tree.Put(v, v)
	}

	b.Log(tree.Count)
}
