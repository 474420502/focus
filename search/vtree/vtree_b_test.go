package vtree

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// BenchmarkRemove-12    	   26240	     45248 ns/op	    1280 B/op	     400 allocs/op
func BenchmarkRemove(t *testing.B) {
	for n := 0; n < t.N; n++ {
		t.StopTimer()
		tree := New()

		for i := 0; i < 1000; i++ {
			istr := strconv.Itoa(i)
			tree.Put([]byte(istr), []byte(istr))
		}

		t.StartTimer()

		for i := 430; i < 830; i++ {
			istr := strconv.Itoa(i)
			tree.Remove([]byte(istr))
		}

	}

	// t.StopTimer()
}

// BenchmarkRemoveRange-12    	  420288	      2712 ns/op	     896 B/op	      36 allocs/op
func BenchmarkRemoveRange(t *testing.B) {

	for n := 0; n < t.N; n++ {
		t.StopTimer()
		tree := New()

		for i := 0; i < 1000; i++ {
			istr := strconv.Itoa(i)
			tree.Put([]byte(istr), []byte(istr))
		}

		t.StartTimer()
		// t.Log(strconv.Itoa(min), strconv.Itoa(max))
		tree.RemoveRange([]byte(strconv.Itoa(430)), []byte(strconv.Itoa(830)))

	}
	t.Log(t.N)
	// t.StopTimer()
}
