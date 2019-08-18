package tried

import (
	"testing"
)

func TestTried_PutAndGet1(t *testing.T) {
	tried := New()
	tried.Put("asdf")
	tried.Put("hehe", "hehe")
	tried.Put("xixi", 3)

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
	tried.Put("abdf", "ab")
	tried.Put("hehe", "hehe")
	tried.Put("xixi", 3)

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
