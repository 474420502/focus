package linkedhashmap

import (
	"reflect"
	"testing"
)

func TestPush(t *testing.T) {
	lhm := New()
	lhm.PushFront(1, "1")
	lhm.PushBack("2", 2)
	var values []interface{}
	values = lhm.Values()

	var testType reflect.Type

	if testType = reflect.TypeOf(values[0]); testType.String() != "string" {
		t.Error(testType)
	}

	if testType = reflect.TypeOf(values[1]); testType.String() != "int" {
		t.Error(testType)
	}

	// 1 2
	lhm.PushFront(4, "4") // 4 1 2
	lhm.PushBack("3", 3)  // 4 1 2 3

	if lhm.String() != "[4 1 2 3]" {
		t.Error(lhm.String())
	}
}

func TestBase(t *testing.T) {
	lhm := New()
	for i := 0; i < 10; i++ {
		lhm.PushBack(i, i)
	}

	if lhm.Empty() {
		t.Error("why lhm Enpty, check it")
	}

	if lhm.Size() != 10 {
		t.Error("why lhm Size != 10, check it")
	}

	lhm.Clear()
	if !lhm.Empty() {
		t.Error("why lhm Clear not Empty, check it")
	}

	if lhm.Size() != 0 {
		t.Error("why lhm Size != 0, check it")
	}
}

func TestGet(t *testing.T) {
	lhm := New()
	for i := 0; i < 10; i++ {
		lhm.PushBack(i, i)
	}

	for i := 0; i < 10; i++ {
		lhm.PushBack(i, i)
	}

	if lhm.Size() != 10 {
		t.Error("why lhm Size != 10, check it")
	}

	for i := 0; i < 10; i++ {
		if v, ok := lhm.Get(i); !ok || v != i {
			t.Error("ok is ", ok, " get value is ", v)
		}
	}
}

func TestRemove(t *testing.T) {
	lhm := New()
	for i := 0; i < 10; i++ {
		lhm.PushBack(i, i)
	}

	var resultStr = "[0 1 2 3 4 5 6 7 8 9]"
	for i := 0; i < 10; i++ {
		if lhm.String() != resultStr {
			t.Error(lhm.String(), resultStr)
		}

		lhm.Remove(i)
		if lhm.Size() != uint(9-i) {
			t.Error("why lhm Size != ", uint(9-i), ", check it")
		}

		resultStr = resultStr[0:1] + resultStr[3:]
	}
}
