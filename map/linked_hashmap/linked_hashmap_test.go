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

	lhm.Put(5, 5)
	if lhm.String() != "[4 1 2 3 5]" {
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

func TestInsert(t *testing.T) {
	lhm := New()
	for i := 0; i < 5; i++ {
		lhm.Insert(0, i, i)
	}

	if lhm.String() != "[4 3 2 1 0]" {
		t.Error(lhm.String())
	}

	if !lhm.Insert(2, 5, 5) {
		t.Error("Insert 2 5 5 error check it")
	}

	if lhm.String() != "[4 3 5 2 1 0]" {
		t.Error(lhm.String())
	}

	if !lhm.Insert(lhm.Size(), 6, 6) {
		t.Error("Insert Size()   error check it")
	}

	if lhm.String() != "[4 3 5 2 1 0 6]" {
		t.Error(lhm.String())
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

	if lhm.Size() != 0 {
		t.Error(lhm.Size())
	}

	for i := 0; i < 10; i++ {
		lhm.PushFront(i, i)
	}

	for i := 0; i < 10; i++ {
		if i >= 5 {
			lhm.Remove(i)
		}
	}

	if lhm.String() != "[4 3 2 1 0]" {
		t.Error(lhm.String())
	}

	// RemoveIndex [4 3 2 1 0]

	if value, _ := lhm.RemoveIndex(2); value != 2 {
		t.Error("[4 3 2 1 0] remove index 2, value is 2, but now is", value)
	}

	// [4 3 1 0]
	if value, _ := lhm.RemoveIndex(2); value != 1 {
		t.Error("[4 3 1 0] remove index 2, value is 1, but now is", value)
	}

	// [4 3 0]
	if value, _ := lhm.RemoveIndex(2); value != 0 {
		t.Error("[4 3 0] remove index 2, value is 0, but now is", value)
	}

	// [4 3]
	if value, _ := lhm.RemoveIndex(2); value != nil {
		t.Error("[4 3] remove index 2, value is nil, but now is", value)
	}

	// [4 3]
	if value, _ := lhm.RemoveIndex(0); value != 4 {
		t.Error("[4 3] remove index 0, value is 4, but now is", value)
	}
}
