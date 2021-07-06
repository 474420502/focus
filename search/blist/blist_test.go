package blist

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/474420502/focus/compare"
	"github.com/474420502/focus/tree/avlkeydup"
	"github.com/Pallinder/go-randomdata"
	"github.com/emirpasic/gods/trees/avltree"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func TestPut(t *testing.T) {
	bl := New()
	bl.IsDebug = -1
	var dict = make(map[int]bool)

	for i := 0; i < 20; i++ {
		// for _, n := range []int{42, 15, 35, 34} {
		n := randomdata.Number(0, 100)
		if _, ok := dict[n]; !ok {
			dict[n] = true
		} else {
			i--
			continue
		}

		k := []byte(strconv.FormatInt(int64(n), 10))

		log.Println("put:", string(k))
		bl.Put(k, k)
		log.Println(bl.debugString())

		if bl.root.size == 10 {
			bl.IsDebug = 0
		}

		if bl.IsDebug > 0 {
			log.Println("isDebug:", bl.IsDebug)
			log.Println(bl.debugString())
			bl.IsDebug = -1
		}

	}

}

func TestPut2(t *testing.T) {
	bl := New()
	tree := avlkeydup.New(compare.Int)
	gods := avltree.NewWith(compare.Int)

	for i := 0; i <= 40; i++ {
		var k []byte
		if i == 12 {
			tree.Put(41, 41)
			gods.Put(41, 41)
			k = []byte(strconv.FormatInt(int64(41), 10))
		} else {
			tree.Put(i, i)
			gods.Put(i, i)
			k = []byte(strconv.FormatInt(int64(i), 10))
		}

		bl.Put(k, k)
		if i == 12 {
			log.Println("put:", string(k))
			log.Println(bl.debugString())
			log.Println("")
			log.Println(tree.String())
			log.Println(gods.String())
			// bl.IsDebug = 0
		}

		if i == 11 {
			log.Println("put:", string(k))
			log.Println(bl.debugString())
			log.Println("")
			log.Println(tree.String())
			log.Println(gods.String())
			// bl.IsDebug = -1
		}
	}

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

func BenchmarkCase1(b *testing.B) {

	data := loadTestData()
	bl := New()
	var l [][]byte
	for i := range data {
		l = append(l, []byte(strconv.Itoa(data[i])))
	}

	// b.SkipNow()
	b.ReportAllocs()
	b.ResetTimer()
	b.StartTimer()

	for _, v := range l {
		bl.Put(v, v)
	}

	b.N = len(l)
}

func TestBenchmarkCase1(t *testing.T) {

	data := loadTestData()
	bl := New()
	var l [][]byte
	for i := range data {
		l = append(l, []byte(strconv.Itoa(data[i])))
	}

	// b.SkipNow()
	now := time.Now()

	for _, v := range l {
		bl.Put(v, v)
	}

	log.Println(time.Since(now).Nanoseconds() / int64(len(l)))
}
