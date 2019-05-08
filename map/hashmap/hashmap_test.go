package hashmap

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"log"
	"testing"
)

func loadTestData() []int {
	log.SetFlags(log.Lshortfile)

	data, err := ioutil.ReadFile("../l.log")
	if err != nil {
		log.Println(err)
	}
	var l []int
	decoder := gob.NewDecoder(bytes.NewReader(data))
	decoder.Decode(&l)
	return l
}
func TestCount(t *testing.T) {
	hm := New()
	// for i := 0; i < 100000; i++ {
	// 	hm.Put(i, i)
	// }
	for i := 0; i < 100000; i++ {
		hm.Put(i, i)
	}
	// t.Error(hm.Get(4))
}

// var executeCount = 5
// var compareSize = 100000

// func BenchmarkPut(b *testing.B) {
// 	b.StopTimer()

// 	l := loadTestData()
// 	hm := New()
// 	b.N = len(l) * executeCount

// 	// for i := 0; i < len(l); i++ {
// 	// 	v := l[i]
// 	// 	hm.Put(v, v)
// 	// }

// 	b.StartTimer()

// 	for c := 0; c < executeCount; c++ {
// 		for i := 0; i < len(l); i++ {
// 			v := l[i]
// 			hm.Put(v, v)
// 		}
// 	}

// 	//b.Log(len(hm.table), hm.size)
// 	//PrintMemUsage()
// }

// func BenchmarkGoPut(b *testing.B) {

// 	l := loadTestData()

// 	hm := make(map[int]int)
// 	b.N = len(l) * executeCount
// 	for c := 0; c < executeCount; c++ {
// 		for i := 0; i < len(l); i++ {
// 			v := l[i]
// 			hm[v] = v
// 		}
// 	}

// 	//b.Log(len(m))
// 	//PrintMemUsage()
// }

// func BenchmarkGet(b *testing.B) {

// 	b.StopTimer()
// 	l := loadTestData()
// 	hm := New()
// 	b.N = len(l) * executeCount
// 	for i := 0; i < len(l); i++ {
// 		v := l[i]
// 		hm.Put(v, v)
// 	}
// 	b.StartTimer()

// 	for i := 0; i < b.N; i++ {
// 		hm.Get(i)
// 	}

// 	//b.Log(len(hm.table), hm.size)
// 	//PrintMemUsage()
// }

// func BenchmarkGoGet(b *testing.B) {
// 	b.StopTimer()
// 	l := loadTestData()
// 	hm := make(map[int]int)
// 	b.N = len(l) * executeCount

// 	for i := 0; i < len(l); i++ {
// 		v := l[i]
// 		hm[v] = v
// 	}

// 	b.StartTimer()
// 	for i := 0; i < b.N; i++ {
// 		if _, ok := hm[i]; !ok {

// 		}
// 	}
// 	//b.Log(len(m))
// 	//PrintMemUsage()
// }
