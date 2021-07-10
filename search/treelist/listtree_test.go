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
		log.Println(string(node1.Key.([]byte)), string(node2.key.([]byte)))
		log.Panic(node2)
	}

	TraverseGodAVL(node1.Children[0], node2.children[0])
	TraverseGodAVL(node1.Children[1], node2.children[1])
}

func TestCase1(t *testing.T) {

	var replay = []int{}
	var status = []int{0, 0, 0}

	// var isShow = true

	for n := 0; n < 10000; n++ {
		tree := New()
		avl := avlkeydup.New(compare.ByteArray)
		var record []int64
		for i := 0; i < 10000; i++ {
			var r int
			if i < len(replay) {
				r = replay[i]
			} else {
				r = randomdata.Number(100, 10000000)
				// r = i
			}

			k := []byte(strconv.FormatInt(int64(r), 10))
			record = append(record, int64(r))
			// log.Println("put:", r)
			avl.Put(k, k)
			// log.Println(record)
			// log.Println(avl.String())
			// log.Println(avl.RotateLog)
			tree.Put(k, k)

			// for _, v := range avl.Values() {
			// 	if _, ok := tree.Get(v.([]byte)); !ok {
			// 		log.Println(string(v.([]byte)))
			// 		log.Panic("")
			// 	}
			// }

			// log.Println(tree.debugString(false))

			// if CompatorByte(tree.getRoot().key, avl.Root.Key.([]byte)) != 0 {
			// 	log.Println(string(tree.getRoot().key), string(avl.Root.Key.([]byte)))
			// }

			// TraverseGodAVL(avl.Root, tree.root.children[0])

			// h1 := getAVLHeight(avl)
			// h2 := tree.getHeight()
			// if avl.Count != int(tree.Count) && isShow {
			// 	log.Println(h1, h2, avl.Count, tree.Count)
			// 	log.Println(tree.debugString(false))
			// 	isShow = false
			// }

			// if getAVLHeight(avl)-tree.getHeight() >= 2 || int64(avl.Count) < tree.Count {
			// 	log.Println(h1, h2, avl.Count, tree.Count)
			// 	break
			// 	// log.Println(tree.debugString(false))
			// }

		}

		h1 := getAVLHeight(avl)
		h2 := tree.getHeight()

		if h1 < h2 {
			status[1]++
			status[2] = h2 - h1
		} else {
			status[0]++
			status[2] = 0
		}
		log.Println(h1, h2, avl.Count, tree.Count, status)
		// isShow = true
	}

	// t.Error(tree.root.children[0].size)

	// t.Error(tree.debugString())
	// t.Error(avl.String())
	// t.Error(tree.debugString())
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
	// godsavl := avltree.NewWith(compare.ByteArray)
	// myavl := avlkeydup.New(compare.ByteArray)
	for _, v := range l {
		tree.Put(v, v)
		// godsavl.Put(v, v)
		// myavl.Put(v, v)
	}
	b.StopTimer()

	b.Log(tree.Size(), "rotate count:", tree.Count, tree.getHeight()) // BenchmarkPut-12    	 5000000	      1287 ns/op	      94 B/op	       1 allocs/op
	// b.Log(godsavl.Size(), getGodsAVLHeight(godsavl))
	// b.Log(myavl.Size(), "rotate count:", myavl.Count, getAVLHeight(myavl)) //    rotate count: 3325439 27

	// b.Log(tree.Count, tree.Size(), tree.getHeight(), getGodsAVLHeight(godsavl), getAVLHeight(myavl))
}

func BenchmarkAVLPut(b *testing.B) {

	d := loadTestData()
	// var dict map[int]bool = make(map[int]bool)

	var l [][]byte
	for _, v := range d {
		l = append(l, []byte(strconv.Itoa(v)))
	}

	b.ResetTimer()
	b.StartTimer()

	b.N = len(l)

	myavl := avlkeydup.New(compare.ByteArray)
	for _, v := range l {

		myavl.Put(v, v)
	}
	b.StopTimer()

	b.Log(myavl.Size(), "rotate count:", myavl.Count, getAVLHeight(myavl)) // 990148 690663 24 1367616 25
}

func BenchmarkGet(b *testing.B) {
	// myavl := avlkeydup.New(compare.ByteArray)
	tree := New()
	d := loadTestData()
	// var dict map[int]bool = make(map[int]bool)

	var l [][]byte
	for _, v := range d {
		k := []byte(strconv.Itoa(v))
		l = append(l, []byte(strconv.Itoa(v)))
		// if _, ok := dict[v]; !ok {
		// 	l = append(l, []byte(strconv.Itoa(v)))
		// 	dict[v] = true
		// }
		// myavl.Put(k, k)
		tree.Put(k, k)
	}

	b.ResetTimer()
	b.StartTimer()

	b.N = len(l)
	// tree := New()
	// godsavl := avltree.NewWith(compare.ByteArray)

	for _, v := range l {
		tree.Get(v)
		// godsavl.Put(v, v)
		// myavl.Get(v)
	}

	// b.Log(tree.Size(), tree.Count, tree.getHeight())
	// b.Log(godsavl.Size(), getGodsAVLHeight(godsavl))
	// b.Log(myavl.Size(), myavl.Count, getAVLHeight(myavl)) // 990148 690663 24

	// b.Log(tree.Count, tree.Size(), tree.getHeight(), getGodsAVLHeight(godsavl), getAVLHeight(myavl))
}

func BenchmarkAVLGet(b *testing.B) {
	myavl := avlkeydup.New(compare.ByteArray)
	// tree := New()
	d := loadTestData()
	// var dict map[int]bool = make(map[int]bool)

	var l [][]byte
	for _, v := range d {
		k := []byte(strconv.Itoa(v))
		l = append(l, []byte(strconv.Itoa(v)))

		myavl.Put(k, k)

	}

	b.ResetTimer()
	b.StartTimer()

	b.N = len(l)
	// tree := New()
	// godsavl := avltree.NewWith(compare.ByteArray)

	for _, v := range l {
		// tree.Get(v)
		// godsavl.Put(v, v)
		myavl.Get(v)
	}

	// b.Log(tree.Size(), tree.Count, tree.getHeight())
	// b.Log(godsavl.Size(), getGodsAVLHeight(godsavl))
	// b.Log(myavl.Size(), myavl.Count, getAVLHeight(myavl)) // 990148 690663 24

	// b.Log(tree.Count, tree.Size(), tree.getHeight(), getGodsAVLHeight(godsavl), getAVLHeight(myavl))
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
