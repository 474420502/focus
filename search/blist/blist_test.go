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
	"github.com/Pallinder/go-randomdata"
	"github.com/emirpasic/gods/trees/avltree"
)

func init() {
	log.SetFlags(log.Llongfile)
}

func TestPut(t *testing.T) {
	bl := New()
	bl.IsDebug = -1

	for ii := 0; ii < 10000; ii++ {
		var dict = make(map[string]bool)
		for i := 0; i < 1000; i++ {

			n := randomdata.Number(0, 100000)
			k := []byte(strconv.FormatInt(int64(n), 10))

			if _, ok := dict[string(k)]; !ok {
				dict[string(k)] = true
			} else {
				i--
				continue
			}

			// k := []byte(strconv.FormatInt(int64(n), 10))

			// log.Println("put:", string(k))
			bl.Put(k, k)
			// log.Println(bl.debugString())

			// if bl.root.size == 10 {
			// 	bl.IsDebug = 0
			// }

			// if bl.IsDebug > 0 {
			// 	log.Println("isDebug:", bl.IsDebug)
			// 	log.Println(bl.debugString())
			// 	bl.IsDebug = -1
			// }

		}

		// for k := range dict {
		// 	if _, ok := bl.Get([]byte(strconv.FormatInt(int64(k), 10))); !ok {
		// 		t.Error("找不到", k)
		// 		break
		// 	}
		// }

		var cur = bl.root
		for cur.children[1] != nil {
			cur = cur.children[1]
		}

		for cur.direct[0] != nil {
			v := cur.direct[0]
			if bl.compartor(cur.key, v.key) <= 0 {
				log.Println("链表错误", string(cur.key), string(v.key))
			}
			cur = v
		}
	}

}

func TestPut2(t *testing.T) {
	bl := New()
	// tree := avlkeydup.New(compare.Int)
	gods := avltree.NewWith(compare.Int)

	// var look = 50

	for i := 0; i <= 100; {
		var k []byte
		// if i == 10002 {
		// 	tree.Put(41, 41)
		// 	gods.Put(41, 41)
		// 	k = []byte(strconv.FormatInt(int64(41), 10))
		// } else {
		// 	tree.Put(i, i)
		// 	gods.Put(i, i)
		// 	k = []byte(strconv.FormatInt(int64(i), 10))
		// }

		if i == 84 {
			i = 63
		} else {
			i += 2
		}
		var input int = i
		// if i < 84 {
		// 	i += 2
		// 	input = i
		// } else {
		// 	reader := bufio.NewReader(os.Stdin)
		// 	fmt.Print("Enter text: ")
		// 	text, _ := reader.ReadString('\n')
		// 	if text == "" {
		// 		i += 2
		// 		input = i
		// 	} else {
		// 		si, err := strconv.Atoi(text)
		// 		if err != nil {
		// 			log.Println(err)
		// 			return
		// 		}
		// 		input = si
		// 	}

		// }

		// tree.Put(i, i)
		gods.Put(input, input)
		k = []byte(strconv.FormatInt(int64(input), 10))
		bl.Put(k, k)

		log.Println("put:", string(k), "count:", bl.Count)
		// log.Println(bl.debugString())
		log.Println(gods.String())

		// if i == look {
		// 	log.Println("put:", string(k), "count:", bl.Count)
		// 	log.Println(bl.debugString())
		// 	log.Println("")
		// 	// log.Println(tree.String())
		// 	// log.Println(gods.String())
		// 	bl.IsDebug = 0
		// }

		// if i == look+1 {
		// 	log.Println("put:", string(k), "count:", bl.Count)
		// 	log.Println(bl.debugString())
		// 	log.Println("")
		// 	// log.Println(tree.String())
		// 	log.Println(gods.String())
		// 	bl.IsDebug = -1
		// }
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
	b.Log(bl.Count)
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

	log.Println(time.Since(now).Nanoseconds()/int64(len(l)), bl.Count)
}
