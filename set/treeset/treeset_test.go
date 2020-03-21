package treeset

import (
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"

	"focus/compare"
	"focus/tree/avldup"
)

func TestTreeSet_Add(t *testing.T) {
	type fields struct {
		tree *avldup.Tree
	}
	type args struct {
		items []interface{}
	}
	tests := []struct {
		name   string
		result string
		fields fields
		args   args
	}{
		{name: "add int", result: "(1, 3, 5)", args: args{items: []interface{}{1, 5, 3, 3, 5}}},
		{name: "add -int", result: "(-5, 1, 5, 3132)", args: args{items: []interface{}{-5, -5, 3132, 3132, 5, 1, 1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := New(compare.Int)
			set.Add(tt.args.items...)
			if set.String() != tt.result {
				t.Error(set.String(), " != ", tt.result)
			}
		})
	}

	tests2 := []struct {
		name   string
		result string
		fields fields
		args   args
	}{
		{name: "add String 1", result: "(1, 3, 5)", args: args{items: []interface{}{"1", "5", "3", "3", "5"}}},
		{name: "add String 2", result: "(-5, 1, 3132, 5)", args: args{items: []interface{}{"-5", "-5", "3132", "3132", "5", "1", "1", "1"}}},
		{name: "add String 3", result: "(a, aa, b, bc)", args: args{items: []interface{}{"a", "b", "aa", "aa", "bc"}}},
		{name: "add String 4", result: "(他, 你, 我, 我我)", args: args{items: []interface{}{"我", "你", "他", "我", "我我"}}},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			set := New(compare.String)
			set.Add(tt.args.items...)
			if set.String() != tt.result {
				t.Error(set.String(), " != ", tt.result)
			}

			vstr := spew.Sprint(set.Values())

			if vstr[1:len(vstr)-1] != strings.ReplaceAll(tt.result[1:len(tt.result)-1], ",", "") {
				t.Error(vstr[1:len(vstr)-1], tt.result[1:len(tt.result)-1])
			}
		})
	}
}

func TestTreeSet_Remove(t *testing.T) {
	type fields struct {
		tree *avldup.Tree
	}
	type args struct {
		addItems    []interface{}
		removeItems []interface{}
	}
	tests := []struct {
		name   string
		result string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{name: "remove 1", result: "()",
			args: args{
				addItems:    []interface{}{5, 7, 5, 3, 2},
				removeItems: []interface{}{5, 7, 3, 2}},
		},

		{name: "remove 2", result: "(5)",
			args: args{
				addItems:    []interface{}{5, 7, 5, 3, 2},
				removeItems: []interface{}{7, 3, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := New(compare.Int)
			set.Add(tt.args.addItems...)
			set.Remove(tt.args.removeItems...)

			if set.String() != tt.result {
				t.Error(set.String(), " != ", tt.result)
			}
		})
	}
}

func TestTreeSet_Base(t *testing.T) {
	set := New(compare.Int)
	set.Add(5, 4, 3, 5)

	if !set.Contains(5) {
		t.Error("Contains 5 is true")
	}

	set.Clear()

	if set.Contains(5) {
		t.Error("Contains 5 is false")
	}

	if !set.Empty() {
		t.Error("should be Empty")
	}

	set.Add(1, 1, 1, 2, 2, 3, 7, 9, 10)

	var results []interface{}
	results = set.GetRange(4, 10)
	if spew.Sprint(results) != "[7 9 10]" {
		t.Error(results)
	}

	results = set.GetRange(1, 10)
	if spew.Sprint(results) != "[1 2 3 7 9 10]" {
		t.Error(results)
	}

	results3 := set.GetAround(3)
	if spew.Sprint(results3) != "[2 3 7]" {
		t.Error(results3)
	}

	results3 = set.GetAround(1)
	if spew.Sprint(results3) != "[<nil> 1 2]" {
		t.Error(results3)
	}
}

func TestTreeSet_Iterator(t *testing.T) {
	set := New(compare.Int)
	set.Add(5, 4, 3, 5)

	iter := set.Iterator()
	iter.ToHead()

	// if not call Next Prev will error
	// 5 4 3
	// if iter.Value() != nil {
	// 	t.Error(iter.Value())
	// }

	iter.Next()
	if iter.Value() != 3 {
		t.Error(iter.Value())
	}

	iter.Next()
	if iter.Value() != 4 {
		t.Error(iter.Value())
	}

	iter.ToTail()
	iter.Prev()
	if iter.Value() != 5 {
		t.Error(iter.Value())
	}

	iter.Prev()
	if iter.Value() != 4 {
		t.Error(iter.Value())
	}

	iter.ToHead()
	iter.Prev()
	if iter.Value() != 3 {
		t.Error(iter.Value())
	}

	iter.ToTail()
	iter.Next()
	if iter.Value() != 5 {
		t.Error(iter.Value())
	}
}
