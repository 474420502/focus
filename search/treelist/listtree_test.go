package listtree

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"strconv"
	"testing"

	"github.com/474420502/focus/compare"
	"github.com/474420502/focus/tree/avlkeydup"
	"github.com/Pallinder/go-randomdata"
	"github.com/emirpasic/gods/trees/avltree"
)

func TraverseGodAVL(node1 *avltree.Node, node2 *Node) {

	if node1 == nil && node2 == nil {
		return
	}

	// log.Println(string(node1.Key.([]byte)), string(node2.key))

	if CompatorMath(node1.Key.([]byte), node2.key) != 0 {
		log.Println(string(node1.Key.([]byte)), string(node2.key))
		log.Panic(node2)
	}

	TraverseGodAVL(node1.Children[0], node2.children[0])
	TraverseGodAVL(node1.Children[1], node2.children[1])
}

func TestCase1(t *testing.T) {

	for n := 0; n < 1000000; n++ {
		tree := New()
		avl := avlkeydup.New(compare.ByteArray)
		for i := int64(0); i < 100; i++ {
			r := randomdata.Number(0, 1000)
			k := []byte(strconv.FormatInt(int64(r), 10))
			log.Println(r)
			avl.Put(k, k)
			log.Println(avl.String())
			tree.Put(k, k)

			// for _, v := range avl.Values() {
			// 	if _, ok := tree.Get(v.([]byte)); !ok {
			// 		log.Println(string(v.([]byte)))
			// 		log.Panic("")
			// 	}
			// }

			log.Println(tree.debugString(true))

			if CompatorByte(tree.getRoot().key, avl.Root.Key.([]byte)) != 0 {
				log.Println(tree.root.key, avl.Root.Key)
			}

			// TraverseGodAVL(avl.Root, tree.root.children[0])
			h1 := getAVLHeight(avl)
			h2 := tree.getHeight()
			if getAVLHeight(avl) != tree.getHeight() {
				log.Println(h1, h2)
			}
		}

	}

	// t.Error(tree.root.children[0].size)

	// t.Error(tree.debugString())
	// t.Error(avl.String())
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

func getGodsAVLHeight(tree *avltree.Tree) int {
	root := tree.Root

	var height = 1

	var traverse func(cur *avltree.Node, h int)
	traverse = func(cur *avltree.Node, h int) {

		if cur == nil {
			return
		}

		if h > height {
			height = h
		}

		traverse(cur.Children[0], h+1)
		traverse(cur.Children[1], h+1)
	}

	traverse(root, 1)

	return height
}

func getAVLHeight(tree *avlkeydup.Tree) int {
	root := tree.Root

	var height = 1

	var traverse func(cur *avlkeydup.Node, h int)
	traverse = func(cur *avlkeydup.Node, h int) {

		if cur == nil {
			return
		}

		if h > height {
			height = h
		}

		traverse(cur.Children[0], h+1)
		traverse(cur.Children[1], h+1)
	}

	traverse(root, 1)

	return height
}

func BenchmarkPut(b *testing.B) {

	d := loadTestData()
	// var dict map[int]bool = make(map[int]bool)

	var l [][]byte
	for _, v := range d {
		l = append(l, []byte(strconv.Itoa(v)))
		// if _, ok := dict[v]; !ok {
		// 	l = append(l, []byte(strconv.Itoa(v)))
		// 	dict[v] = true
		// }
	}

	b.ResetTimer()
	b.StartTimer()

	b.N = len(l)
	tree := New()
	godsavl := avltree.NewWith(compare.ByteArray)
	myavl := avlkeydup.New(compare.ByteArray)
	for _, v := range l {
		tree.Put(v, v)
		godsavl.Put(v, v)
		myavl.Put(v, v)
	}

	b.Log(tree.Count, tree.Size(), tree.getHeight(), getGodsAVLHeight(godsavl), getAVLHeight(myavl))
}
