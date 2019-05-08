package array2

import (
	"testing"

	"github.com/davecgh/go-spew/spew"

	"github.com/Pallinder/go-randomdata"
)

func testSet1(t *testing.T) {
	arr := NewWithCap(4, 4)
	l := []int{0, 6, 5, 15}
	for _, v := range l {
		arr.Set(v, v)
	}

	var result string
	result = spew.Sprint(arr.debugValues())
	if result != "[0 {} {} {} {} 5 6 {} <nil> <nil> <nil> <nil> {} {} {} 15]" {
		t.Error(result)
	}

	defer func() {
		if err := recover(); err == nil {
			t.Error("err == nil, but array the set is out of range")
		}
	}()

	arr.Set(16, 16)
}

func testSet2(t *testing.T) {
	arr := NewWithCap(5, 3)
	l := []int{0, 6, 5, 14}
	for _, v := range l {
		arr.Set(v, v)
	}

	var result string
	result = spew.Sprint(arr.debugValues())
	if result != "[0 {} {} {} {} 5 6 {} {} <nil> <nil> <nil> {} {} 14]" {
		t.Error(result)
	}

	defer func() {
		if err := recover(); err == nil {
			t.Error("err == nil, but array the set is out of range")
		}
	}()

	arr.Set(16, 16)
}

func TestArray2Set(t *testing.T) {
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

	defer func() {
		if err := recover(); err == nil {
			t.Error("err == nil, but array the get is out of range")
		}
	}()

	arr.Get(64)
}

func testArray2Get2(t *testing.T) {
	arr := NewWithCap(9, 8)
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

	defer func() {
		if err := recover(); err == nil {
			t.Error("err == nil, but array the get is out of range")
		}
	}()

	arr.Get(72)
}

func TestArray2Get(t *testing.T) {
	testArray2Get1(t)
	testArray2Get2(t)
}

func TestArray2Del(t *testing.T) {
	arr := NewWithCap(3, 6)
	l := []int{0, 6, 5, 15}
	for _, v := range l {
		arr.Set(v, v)
	}
	// default  [0 {} {} {} {} 5 6 {} <nil> <nil> <nil> <nil> {} {} {} 15]
	var result string

	arr.Del(0)
	result = spew.Sprint(arr.debugValues())
	if result != "[{} {} {} {} {} 5 6 {} {} {} {} {} {} {} {} 15 {} {}]" {
		t.Error(arr.data)
		t.Error(result)
	}

	arr.Del(5)
	result = spew.Sprint(arr.debugValues())
	if result != "[<nil> <nil> <nil> <nil> <nil> <nil> 6 {} {} {} {} {} {} {} {} 15 {} {}]" {
		t.Error(arr.data)
		t.Error(result)
	}

	arr.Del(6)
	result = spew.Sprint(arr.debugValues())
	if result != "[<nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> {} {} {} 15 {} {}]" {
		t.Error(result)
	}

	arr.Del(15)
	result = spew.Sprint(arr.debugValues())
	if result != "[<nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil> <nil>]" {
		t.Error(result)
	}

	defer func() {
		if err := recover(); err == nil {
			t.Error("err == nil, but array the del is out of range")
		}
	}()

	arr.Del(18)
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

	result = spew.Sprint(arr.debugValues())
	if result != "[0 {} {} {} {} 5 6 {} <nil> <nil> <nil> <nil> {} {} {} 15]" {
		t.Error(result)
	}
}

func BenchmarkArray2Set(b *testing.B) {

	arr := NewWithCap(1000, 100)
	b.N = 500000000

	b.StopTimer()
	var l []int
	for i := 0; i < b.N/10; i++ {
		l = append(l, randomdata.Number(0, 65535))
	}
	b.StartTimer()
	for c := 0; c < 10; c++ {
		for i := 0; i < b.N/10; i++ {
			arr.Set(l[i], i)
		}
	}

}

func BenchmarkArray2Get(b *testing.B) {

	arr := NewWithCap(1000, 100)
	b.N = 500000000

	b.StopTimer()

	for i := 0; i < 105535; i++ {
		v := randomdata.Number(0, 65535)
		arr.Set(v, v)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		arr.Get(i % 65535)
	}

}

func BenchmarkArray2Del(b *testing.B) {

	arr := NewWithCap(1000, 100)
	b.N = 500000000

	b.StopTimer()
	for i := 0; i < 105535; i++ {
		v := randomdata.Number(0, 65535)
		arr.Set(v, v)
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		arr.Del(i % 65535)
	}

}
