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

func BenchmarkRemove(t *testing.B) {
	for n := 0; n < t.N; n++ {
		t.StopTimer()
		tree := New()

		for i := 0; i < 1000; i += 2 {
			istr := strconv.Itoa(i)
			tree.Put([]byte(istr), []byte(istr))
		}

		values := tree.Values()
		for i := len(values) - 1; i > 0; i-- {
			num := rand.Intn(i + 1)
			values[i], values[num] = values[num], values[i]
		}

		t.StartTimer()

		for _, v := range values[0 : len(values)/2] {
			tree.Remove(v)
		}
	}

	// t.StopTimer()
}

func BenchmarkRemoveRange(t *testing.B) {

	for n := 0; n < t.N; n++ {
		t.StopTimer()
		tree := New()

		for i := 0; i < 1000; i += 2 {
			istr := strconv.Itoa(i)
			tree.Put([]byte(istr), []byte(istr))
		}

		values := tree.Values()

		min := rand.Intn(1000)
		max := min + len(values)/2

		t.StartTimer()
		t.Log(strconv.Itoa(min), strconv.Itoa(max))
		tree.RemoveRange([]byte(strconv.Itoa(min)), []byte(strconv.Itoa(max)))

	}

	// t.StopTimer()
}
