package tried

import (
	"testing"

	"github.com/Pallinder/go-randomdata"
)

func TestTried_PutAndGet1(t *testing.T) {
	tried := New()

	tried.Put(TriedString("asdf"))
	tried.Put(TriedString("hehe"), "hehe")
	tried.Put(TriedString("xixi"), 3)

	var result interface{}

	result = tried.Get(TriedString("asdf"))
	if result != tried {
		t.Error("result should be 3")
	}

	result = tried.Get(TriedString("xixi"))
	if result != 3 {
		t.Error("result should be 3")
	}

	result = tried.Get(TriedString("hehe"))
	if result != "hehe" {
		t.Error("result should be hehe")
	}

	result = tried.Get(TriedString("haha"))
	if result != nil {
		t.Error("result should be nil")
	}

	result = tried.Get(TriedString("b"))
	if result != nil {
		t.Error("result should be nil")
	}
}

func TestTried_Traversal(t *testing.T) {
	tried := New()
	tried.Put(TriedString("asdf"))
	tried.Put(TriedString("abdf"), "ab")
	tried.Put(TriedString("hehe"), "hehe")
	tried.Put(TriedString("xixi"), 3)

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

	var data []TriedString
	b.N = 10000
	count := 1000

	for i := 0; i < b.N; i++ {
		data = append(data, TriedString(randomdata.RandStringRunes(10)+randomdata.RandStringRunes(4)))
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

	var data []TriedString
	b.N = 10000
	count := 1000

	for i := 0; i < b.N; i++ {
		data = append(data, TriedString(randomdata.RandStringRunes(10)+randomdata.RandStringRunes(4)))
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
