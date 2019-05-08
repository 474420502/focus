package arrayn

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func testSet1(t *testing.T) {
	arr := NewWithCap(2, 2, 2)
	l := []int{0, 1, 7}
	for _, v := range l {
		arr.Set(v, v)
	}

	var result string
	result = spew.Sprint(arr.Values())
	if result != "[0 1 <nil> <nil> <nil> <nil> <nil> 7]" {
		t.Error(result)
	}

	defer func() {
		if err := recover(); err == nil {
			t.Error("err == nil, but array the set is out of range")
		}
	}()

	arr.Set(8, 8)
}

func testSet2(t *testing.T) {
	arr := NewWithCap(2, 2, 3)
	l := []int{0, 6, 5, 11}
	for _, v := range l {
		arr.Set(v, v)
	}

	var result string
	result = spew.Sprint(arr.Values())
	if result != "[0 <nil> <nil> <nil> <nil> 5 6 <nil> <nil> <nil> <nil> 11]" {
		t.Error(arr.data)
		t.Error(result)
	}

	defer func() {
		if err := recover(); err == nil {
			t.Error("err == nil, but array the set is out of range")
		}
	}()

	arr.Set(12, 12)
}

func TestSet(t *testing.T) {
	testSet1(t)
	testSet2(t)
}

func testArray2Get1(t *testing.T) {
	arr := New()
	for i := 0; i < 64; i++ {
		arr.Set(i, i)
	}

	for i := 0; i < 64; i++ {
		if v, ok := arr.Get(i); ok {
			if v != i {
				t.Error("v is equal i, but", v, i)
			}
		} else {
			t.Error("not ok is error")
		}
	}

	if v, ok := arr.Get(8*8*8 - 1); ok {
		t.Error(v)
	}

	defer func() {
		if err := recover(); err == nil {
			t.Error("err == nil, but array the get is out of range")
		}
	}()

	arr.Get(8 * 8 * 8)
}

func testArray2Get2(t *testing.T) {
	arr := NewWithCap(4, 3, 3)
	for i := 0; i < 36; i++ {
		arr.Set(i, i)
	}

	for i := 0; i < 36; i++ {
		if v, ok := arr.Get(i); ok {
			if v != i {
				t.Error("v is equal i, but", v, i)
			}
		} else {
			t.Error("not ok is error")
		}
	}

	defer func() {
		if err := recover(); err == nil {
			t.Error("err == nil, but array the get is out of range")
		}
	}()

	arr.Get(36)
}

func TestArray2Get(t *testing.T) {
	testArray2Get1(t)
	testArray2Get2(t)
}

func TestDel(t *testing.T) {
	arr := NewWithCap(2, 2, 2, 3)
	for i := 0; i < 12; i++ {
		arr.Set(i, i)
	}

	arr.Set(23, 23)

	for i := 0; i < 12; i++ {
		arr.Del(i)
	}

	arr.Del(23)
	var result string

	result = spew.Sprint(arr.Values())
	if result != "[<nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>]" {
		t.Error("result should be all is nil")
	}
}

func TestArray2Grow(t *testing.T) {
	arr := NewWithCap(4, 4)
	l := []int{0, 6, 5, 15}
	for _, v := range l {
		arr.Set(v, v)
	}

	arr.Grow(1)

	if v, ok := arr.Get(15); ok {
		if v != 15 {
			t.Error(v)
		}
	} else {
		t.Error(v)
	}

	arr.Set(19, 19)
	if v, ok := arr.Get(19); ok {
		if v != 19 {
			t.Error(v)
		}
	} else {
		t.Error(v)
	}

	arr.Grow(-1)
	var result string
	result = spew.Sprint(arr.Values())
	if result != "[0 <nil> <nil> <nil> <nil> 5 6 <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> 15]" {
		t.Error(result)
	}
}

// func BenchmarkGoMap(b *testing.B) {
// 	m := make(map[int]bool)
// 	b.N = 50000000
// 	b.StopTimer()
// 	var l []int
// 	for i := 0; i < b.N/10; i++ {
// 		l = append(l, randomdata.Number(0, 100000000))
// 	}
// 	b.StartTimer()
// 	for c := 0; c < 10; c++ {
// 		for i := 0; i < b.N/10; i++ {
// 			m[l[i]] = true
// 		}
// 	}
// }

// func BenchmarkArrayNSet(b *testing.B) {

// 	arr := NewWithCap(1000, 10, 10, 100)
// 	b.N = 10000000

// 	b.StopTimer()
// 	var l []int
// 	for i := 0; i < b.N/10; i++ {
// 		l = append(l, randomdata.Number(0, 10000000))
// 	}
// 	b.StartTimer()
// 	for c := 0; c < 10; c++ {
// 		for i := 0; i < b.N/10; i++ {
// 			arr.Set(l[i], i)
// 		}
// 	}
// }

// func BenchmarkArray3Set(b *testing.B) {

// 	arr := NewWithCap(100, 100, 10)
// 	b.N = 500000000

// 	b.StopTimer()
// 	var l []int
// 	for i := 0; i < b.N/10; i++ {
// 		l = append(l, randomdata.Number(0, 65535))
// 	}
// 	b.StartTimer()
// 	for c := 0; c < 10; c++ {
// 		for i := 0; i < b.N/10; i++ {
// 			arr.Set(l[i], i)
// 		}
// 	}
// }
