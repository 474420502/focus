package tried

import (
	"testing"

	"github.com/Pallinder/go-randomdata"
)

func TestTried_PutAndGet1(t *testing.T) {
	tried := New()

	tried.Put(("asdf"))
	tried.Put(("hehe"), "hehe")
	tried.Put(("xixi"), 3)

	var result interface{}

	result = tried.Get("asdf")
	if result != tried {
		t.Error("result should be 3")
	}

	result = tried.Get("xixi")
	if result != 3 {
		t.Error("result should be 3")
	}

	result = tried.Get("hehe")
	if result != "hehe" {
		t.Error("result should be hehe")
	}

	result = tried.Get("haha")
	if result != nil {
		t.Error("result should be nil")
	}

	result = tried.Get("b")
	if result != nil {
		t.Error("result should be nil")
	}
}

func TestTried_Traversal(t *testing.T) {
	tried := New()
	tried.Put("asdf")
	tried.Put(("abdf"), "ab")
	tried.Put(("hehe"), "hehe")
	tried.Put(("xixi"), 3)

	var result []interface{}
	tried.Traversal(func(idx uint, v interface{}) bool {
		// t.Error(idx, v)
		result = append(result, v)
		return true
	})

	if result[0] != "ab" {
		t.Error(result[0])
	}

	if result[1] != tried {
		t.Error(result[1])
	}

	if result[2] != "hehe" {
		t.Error(result[2])
	}

	if result[3] != 3 {
		t.Error(result[3])
	}
}

func BenchmarkTried_Put(b *testing.B) {

	var data []string
	b.N = 100000
	count := 50

	for i := 0; i < b.N; i++ {
		data = append(data, (randomdata.RandStringRunes(10) + randomdata.RandStringRunes(4)))
	}

	b.ResetTimer()
	b.N = b.N * count
	for c := 0; c < count; c++ {
		tried := New()
		for _, v := range data {
			tried.Put(v)
		}
	}
}

func BenchmarkTried_Get(b *testing.B) {

	var data []string
	b.N = 100000
	count := 50

	for i := 0; i < b.N; i++ {
		data = append(data, (randomdata.RandStringRunes(10) + randomdata.RandStringRunes(4)))
	}

	b.N = b.N * count

	tried := New()
	for _, v := range data {
		tried.Put(v)
	}

	b.ResetTimer()
	for c := 0; c < count; c++ {
		for _, v := range data {
			tried.Get(v)
		}
	}
}
