package linkedlist

import (
	"testing"

	"github.com/Pallinder/go-randomdata"
	"github.com/davecgh/go-spew/spew"
)

func TestPushFront(t *testing.T) {
	l := New()
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}
	var result string
	result = spew.Sprint(l.Values())
	if result != "[4 3 2 1 0]" {
		t.Error(result)
	}

	l.PushFront(0)
	result = spew.Sprint(l.Values())
	if result != "[0 4 3 2 1 0]" {
		t.Error(result)
	}
}

func TestPushBack(t *testing.T) {
	l := New()
	for i := 0; i < 5; i++ {
		l.PushBack(i)
	}
	var result string
	result = spew.Sprint(l.Values())
	if result != "[0 1 2 3 4]" {
		t.Error(result)
	}

	l.PushBack(0)
	result = spew.Sprint(l.Values())
	if result != "[0 1 2 3 4 0]" {
		t.Error(result)
	}
}

func TestPopFront(t *testing.T) {
	l := New()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}
	// var result string

	for i := 4; i >= 0; i-- {
		if v, ok := l.PopFront(); ok {
			if v != i {
				t.Error("[4 3 2 1 0] PopFront value should be ", i, ", but is ", v)
			}
		} else {
			t.Error("PopFront is not ok")
		}

		if l.Size() != uint(i) {
			t.Error("l.Size() is error, is", l.Size())
		}
	}
}

func TestPopBack(t *testing.T) {
	l := New()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}
	// var result string

	for i := 0; i < 5; i++ {
		if v, ok := l.PopBack(); ok {
			if v != i {
				t.Error("[4 3 2 1 0] PopFront value should be ", i, ", but is ", v)
			}
		} else {
			t.Error("PopFront is not ok")
		}

		if l.Size() != uint(5-i-1) {
			t.Error("l.Size() is error, is", l.Size())
		}
	}

}

func TestInsert2(t *testing.T) {
	l := New()

	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(0, i)
	}

	// step1: [0 0] -> step2: [1 0 0 0] front
	if l.String() != "[4 0 3 0 2 0 1 0 0 0]" {
		t.Error(l.String())
	}

	if !l.Insert(l.Size(), 5) {
		t.Error("should be true")
	}

	// step1: [0 0] -> step2: [4 0 3 0 2 0 1 0 0 0] front size is 10, but you can insert 11. equal to PushBack [4 0 3 0 2 0 1 0 0 0 5]
	if l.String() != "[4 0 3 0 2 0 1 0 0 0 5]" {
		t.Error(l.String())
	}
}

func TestInsert1(t *testing.T) {
	l1 := New()
	l2 := New()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l1.Insert(0, i)
		l2.PushFront(i)
	}

	var result1, result2 string
	result1 = spew.Sprint(l1.Values())
	result2 = spew.Sprint(l2.Values())
	if result1 != result2 {
		t.Error(result1, result2)
	}

	for i := 0; i < 5; i++ {
		l1.Insert(l1.Size(), i)
		l2.PushBack(i)
		// t.Error(l1.Values(), l2.Values())
	}

	result1 = spew.Sprint(l1.Values())
	result2 = spew.Sprint(l2.Values())
	if result1 != result2 {
		t.Error(result1, result2)
	}

	if result1 != "[4 3 2 1 0 0 1 2 3 4]" {
		t.Error("result should be [4 3 2 1 0 0 1 2 3 4]\n but result is", result1)
	}

	l1.Insert(1, 99)
	result1 = spew.Sprint(l1.Values())
	if result1 != "[4 99 3 2 1 0 0 1 2 3 4]" {
		t.Error("[4 3 2 1 0 0 1 2 3 4] insert with index 1, should be [4 99 3 2 1 0 0 1 2 3 4]\n but result is", result1)
	}

	l1.Insert(9, 99)
	result1 = spew.Sprint(l1.Values())
	if result1 != "[4 99 3 2 1 0 0 1 2 99 3 4]" {
		t.Error("[4 99 3 2 1 0 0 1 2 3 4] insert with index 9, should be [4 99 3 2 1 0 0 1 2 99 3 4]\n but result is", result1)
	}

	l1.Insert(12, 99)
	result1 = spew.Sprint(l1.Values())
	if result1 != "[4 99 3 2 1 0 0 1 2 99 3 4 99]" {
		t.Error("[4 99 3 2 1 0 0 1 2 99 3 4] insert with index 12, should be [4 99 3 2 1 0 0 1 2 99 3 4 99]\n but result is", result1)
	}
}

func TestInsertIf(t *testing.T) {
	l := New()

	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.Insert(0, i)
	}

	// "[4 3 2 1 0]" 插入两个11
	for i := 0; i < 2; i++ {
		l.InsertIf(func(idx uint, value interface{}) InsertState {
			if value == 3 {
				return InsertFront
			}
			return UninsertAndContinue
		}, 11)
	}

	var result string

	result = spew.Sprint(l.Values())
	if result != "[4 3 11 11 2 1 0]" {
		t.Error("result should be [4 3 11 11 2 1 0], reuslt is", result)
	}

	// "[4 3 2 1 0]"
	for i := 0; i < 2; i++ {
		l.InsertIf(func(idx uint, value interface{}) InsertState {
			if value == 0 {
				return InsertBack
			}
			return UninsertAndContinue
		}, 11)
	}

	result = spew.Sprint(l.Values())
	if result != "[4 3 11 11 2 1 11 11 0]" {
		t.Error("result should be [4 3 11 11 2 1 11 11 0], reuslt is", result)
	}

	// "[4 3 2 1 0]"
	for i := 0; i < 2; i++ {
		l.InsertIf(func(idx uint, value interface{}) InsertState {
			if value == 0 {
				return InsertFront
			}
			return UninsertAndContinue
		}, 11)
	}

	result = spew.Sprint(l.Values())
	if result != "[4 3 11 11 2 1 11 11 0 11 11]" {
		t.Error("result should be [4 3 11 11 2 1 11 11 0 11 11], reuslt is", result)
	}

	// t.Error(l.Values())
}

func TestFind(t *testing.T) {
	l := New()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	if v, isfound := l.Find(func(idx uint, value interface{}) bool {
		if idx == 1 {
			return true
		}
		return false
	}); isfound {
		if v != 3 {
			t.Error("[4 3 2 1 0] index 1 shoud be 3 but value is", v)
		}
	} else {
		t.Error("should be found")
	}

	if v, isfound := l.Find(func(idx uint, value interface{}) bool {
		if idx == 5 {
			return true
		}
		return false
	}); isfound {
		t.Error("should not be found, but v is found, ", v)
	}
}

func TestFindMany(t *testing.T) {
	l := New()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	if values, isfound := l.FindMany(func(idx uint, value interface{}) int {
		if idx >= 1 {
			return 1
		}
		return 0
	}); isfound {
		var result string
		result = spew.Sprint(values)
		if result != "[3 2 1 0]" {
			t.Error("result should be [3 2 1 0], reuslt is", result)
		}
	} else {
		t.Error("should be found")
	}

	if values, isfound := l.FindMany(func(idx uint, value interface{}) int {
		if idx%2 == 0 {
			return 1
		}
		return 0
	}); isfound {
		var result string
		result = spew.Sprint(values)
		if result != "[4 2 0]" {
			t.Error("result should be [3 2 1 0], reuslt is", result)
		}
	} else {
		t.Error("should be found")
	}

	if values, isfound := l.FindMany(func(idx uint, value interface{}) int {
		if value == 0 || value == 2 || value == 4 || value == 7 {
			return 1
		}
		return 0
	}); isfound {
		var result string
		result = spew.Sprint(values)
		if result != "[4 2 0]" {
			t.Error("result should be [4 2 0], reuslt is", result)
		}
	} else {
		t.Error("should be found")
	}

	if values, isfound := l.FindMany(func(idx uint, value interface{}) int {
		if value.(int) <= 2 {
			return -1
		}

		if value.(int) <= 4 && value.(int) > 2 {
			return 1
		}

		return 0
	}); isfound {
		var result string
		result = spew.Sprint(values)
		if result != "[4 3]" {
			t.Error("result should be [4 2 0], reuslt is", result)
		}
	} else {
		t.Error("should be found")
	}
	// if v, isfound := l.Find(func(idx uint, value interface{}) bool {
	// 	if idx == 5 {
	// 		return true
	// 	}
	// 	return false
	// }); isfound {
	// 	t.Error("should not be found, but v is found, ", v)
	// }
}

func TestIndex(t *testing.T) {
	l := New()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	if v, ok := l.Index(4); ok {
		if v != 0 {
			t.Error("[4 3 2 1 0] Index 4 value is 0, but v is ", v)
		}
	} else {
		t.Error("not ok is error")
	}

	if v, ok := l.Index(1); ok {
		if v != 3 {
			t.Error("[4 3 2 1 0] Index 1 value is 3, but v is ", v)
		}
	} else {
		t.Error("not ok is error")
	}

	if v, ok := l.Index(0); ok {
		if v != 4 {
			t.Error("[4 3 2 1 0] Index 1 value is 4, but v is ", v)
		}
	} else {
		t.Error("not ok is error")
	}

	if _, ok := l.Index(5); ok {
		t.Error("[4 3 2 1 0] Index 5, out of range,ok = true is error")
	}

}

func TestRemove(t *testing.T) {
	l := New()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	l.Remove(0)
	var result string
	result = spew.Sprint(l.Values())
	if result != "[3 2 1 0]" {
		t.Error("should be [3 2 1 0] but result is", result)
	}

	l.Remove(3)
	result = spew.Sprint(l.Values())
	if result != "[3 2 1]" {
		t.Error("should be [3 2 1] but result is", result)
	}

	l.Remove(2)
	result = spew.Sprint(l.Values())
	if result != "[3 2]" {
		t.Error("should be [3 2 1] but result is", result)
	}

	l.Remove(1)
	result = spew.Sprint(l.Values())
	if result != "[3]" {
		t.Error("should be [3 2 1] but result is", result)
	}

	l.Remove(0)
	result = spew.Sprint(l.Values())
	if result != "<nil>" && l.Size() == 0 && len(l.Values()) == 0 {
		t.Error("should be [3 2 1] but result is", result, "Size is", l.Size())
	}

	if _, rvalue := l.Remove(3); rvalue != false {
		t.Error("l is empty")
	}

}

func TestRemoveIf(t *testing.T) {
	l := New()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	if result, ok := l.RemoveIf(func(idx uint, value interface{}) RemoveState {
		if value == 0 {
			return RemoveAndContinue
		}
		return UnremoveAndContinue
	}); ok {
		if result[0] != 0 {
			t.Error("result should is", 0)
		}
	} else {
		t.Error("should be ok")
	}

	if result, ok := l.RemoveIf(func(idx uint, value interface{}) RemoveState {
		if value == 4 {
			return RemoveAndContinue
		}
		return UnremoveAndContinue
	}); ok {
		if result[0] != 4 {
			t.Error("result should is", 4)
		}
	} else {
		t.Error("should be ok")
	}

	var result string
	result = spew.Sprint(l.Values())
	if result != "[3 2 1]" {
		t.Error("should be [3 2 1] but result is", result)
	}

	if result, ok := l.RemoveIf(func(idx uint, value interface{}) RemoveState {
		if value == 4 {
			return RemoveAndContinue
		}
		return UnremoveAndContinue
	}); ok {
		t.Error("should not be ok and result is nil")
	} else {
		if result != nil {
			t.Error("should be nil")
		}
	}

	result = spew.Sprint(l.Values())
	if result != "[3 2 1]" {
		t.Error("should be [3 2 1] but result is", result)
	}

	l.RemoveIf(func(idx uint, value interface{}) RemoveState {
		if value == 3 || value == 2 || value == 1 {
			return RemoveAndContinue
		}
		return UnremoveAndContinue
	})

	result = spew.Sprint(l.Values())
	if result != "<nil>" {
		t.Error("result should be <nil>, but now result is", result)
	}

	if results, ok := l.RemoveIf(func(idx uint, value interface{}) RemoveState {
		if value == 3 || value == 2 || value == 1 {
			return RemoveAndContinue
		}
		return UnremoveAndContinue
	}); ok {
		t.Error("why  ok")
	} else {
		if results != nil {
			t.Error(results)
		}
	}

}

func TestRemoveIf2(t *testing.T) {
	l := New()
	// "[4 3 2 1 0]"
	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	for i := 0; i < 5; i++ {
		l.PushFront(i)
	}

	// 只删除一个
	if result, ok := l.RemoveIf(func(idx uint, value interface{}) RemoveState {
		if value == 0 {
			return RemoveAndBreak
		}
		return UnremoveAndContinue
	}); ok {
		if result[0] != 0 {
			t.Error("result should is", 0)
		}
	} else {
		t.Error("should be ok")
	}

	var resultstr string
	resultstr = spew.Sprint(l.Values())
	if resultstr != "[4 3 2 1 4 3 2 1 0]" {
		t.Error("result should is", resultstr)
	}

	// 只删除多个
	if result, ok := l.RemoveIf(func(idx uint, value interface{}) RemoveState {
		if value == 4 {
			return RemoveAndContinue
		}
		return UnremoveAndContinue
	}); ok {

		resultstr = spew.Sprint(result)
		if resultstr != "[4 4]" {
			t.Error("result should is", result)
		}

		resultstr = spew.Sprint(l.Values())
		if resultstr != "[3 2 1 3 2 1 0]" {
			t.Error("result should is", resultstr)
		}

	} else {
		t.Error("should be ok")
	}
}

func TestIterator(t *testing.T) {
	ll := New()
	for i := 0; i < 10; i++ {
		ll.PushFront(i)
	}

	iter := ll.Iterator()

	for i := 0; iter.Next(); i++ {
		if iter.Value() != 9-i {
			t.Error("iter.Next() ", iter.Value(), "is not equal ", 9-i)
		}
	}

	if iter.cur != iter.ll.tail {
		t.Error("current point is not equal tail ", iter.ll.tail)
	}

	for i := 0; iter.Prev(); i++ {
		if iter.Value() != i {
			t.Error("iter.Prev() ", iter.Value(), "is not equal ", i)
		}
	}
}

func TestCircularIterator(t *testing.T) {
	ll := New()
	for i := 0; i < 10; i++ {
		ll.PushFront(i)
	}

	iter := ll.CircularIterator()

	for i := 0; i != 10; i++ {
		iter.Next()
		if iter.Value() != 9-i {
			t.Error("iter.Next() ", iter.Value(), "is not equal ", 9-i)
		}
	}

	if iter.cur != iter.pl.tail.prev {
		t.Error("current point is not equal tail ", iter.pl.tail.prev)
	}

	if iter.Next() {
		if iter.Value() != 9 {
			t.Error("iter.Value() != ", iter.Value())
		}
	}

	iter.MoveToTail()
	for i := 0; i != 10; i++ {
		iter.Prev()
		if iter.Value() != i {
			t.Error("iter.Prev() ", iter.Value(), "is not equal ", i)
		}
	}

	if iter.cur != iter.pl.head.next {
		t.Error("current point is not equal tail ", iter.pl.tail.prev)
	}

	if iter.Prev() {
		if iter.Value() != 0 {
			t.Error("iter.Value() != ", iter.Value())
		}
	}
}

func BenchmarkPushBack(b *testing.B) {

	ec := 5
	cs := 2000000
	b.N = cs * ec

	for c := 0; c < ec; c++ {
		l := New()
		for i := 0; i < cs; i++ {
			l.PushBack(i)
		}
	}
}

func BenchmarkPushFront(b *testing.B) {

	ec := 5
	cs := 2000000
	b.N = cs * ec

	for c := 0; c < ec; c++ {
		l := New()
		for i := 0; i < cs; i++ {
			l.PushFront(i)
		}
	}

}

func BenchmarkInsert(b *testing.B) {

	ec := 10
	cs := 1000
	b.N = cs * ec

	for c := 0; c < ec; c++ {
		l := New()
		for i := 0; i < cs; i++ {
			ridx := randomdata.Number(0, int(l.Size())+1)
			l.Insert(uint(ridx), i)
		}
	}
}
