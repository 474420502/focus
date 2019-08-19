package tried

import (
	"bytes"
	"encoding/gob"
	"os"
	"testing"

	"github.com/Pallinder/go-randomdata"
)

func TestTried_NewWith(t *testing.T) {
	tried := NewWithWordType(WordIndex32to126)
	words := "~ 23fd "
	tried.Put(words)
	if tried.Get(words) == nil {
		t.Error("should be not nil")
	}

	tried = NewWithWordType(WordIndexLower)
	words = "az"
	tried.Put(words)
	if tried.Get(words) == nil {
		t.Error("should be not nil")
	}

	tried = NewWithWordType(WordIndexUpper)
	words = "AZ"
	tried.Put(words)
	if tried.Get(words) == nil {
		t.Error("should be not nil")
	}

	tried = NewWithWordType(WordIndexUpperLower)
	words = "AZazsdfsd"
	tried.Put(words)
	if tried.Get(words) == nil {
		t.Error("should be not nil")
	}
}

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

func TesStoreData(t *testing.T) {
	var l []string
	const N = 1000000
	for i := 0; i < N; i++ {
		var content []rune
		for c := 0; c < randomdata.Number(5, 15); c++ {
			char := randomdata.Number(0, 26) + 'a'
			content = append(content, rune(byte(char)))
		}
		l = append(l, (string(content)))
	}

	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	encoder.Encode(l)
	lbytes := result.Bytes()
	f, _ := os.OpenFile("tried.log", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	f.Write(lbytes)
}

func Load() []string {
	var result []string
	f, _ := os.Open("tried.log")
	gob.NewDecoder(f).Decode(&result)
	return result
}

func BenchmarkTried_Put(b *testing.B) {

	var data []string
	b.N = 1000000
	count := 10

	// for i := 0; i < b.N; i++ {
	// 	var content []rune
	// 	for c := 0; c < randomdata.Number(5, 15); c++ {
	// 		char := randomdata.Number(0, 26) + 'a'
	// 		content = append(content, rune(byte(char)))
	// 	}
	// 	data = append(data, (string(content)))
	// }

	data = Load()

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
	b.StopTimer()
	var data []string
	b.N = 1000000
	count := 10

	// for i := 0; i < b.N; i++ {
	// 	var content []rune
	// 	for c := 0; c < randomdata.Number(5, 15); c++ {
	// 		char := randomdata.Number(0, 26) + 'a'
	// 		content = append(content, rune(byte(char)))
	// 	}
	// 	data = append(data, string(content))
	// }
	data = Load()

	b.N = b.N * count

	tried := New()
	for _, v := range data {
		tried.Put(v)
	}

	b.StartTimer()
	for c := 0; c < count; c++ {
		for _, v := range data {
			tried.Get(v)
		}
	}
}
