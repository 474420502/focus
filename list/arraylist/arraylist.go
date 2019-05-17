package arraylist

import (
	"fmt"
	"strings"

	"github.com/emirpasic/gods/lists"
	"github.com/emirpasic/gods/utils"
)

func assertListImplementation() {
	var _ lists.List = (*ArrayList)(nil)
}

type ArrayList struct {
	data []interface{}
	size int
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// New instantiates a new list and adds the passed values, if any, to the list
func New(values ...interface{}) *ArrayList {
	list := &ArrayList{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add appends a value at the end of the list
func (list *ArrayList) Add(values ...interface{}) {
	list.growBy(len(values))
	for _, value := range values {
		list.data[list.size] = value
		list.size++
	}
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *ArrayList) Get(index int) (interface{}, bool) {

	if !list.withinRange(index) {
		return nil, false
	}

	return list.data[index], true
}

// Remove removes the element at the given index from the list.
func (list *ArrayList) Remove(index int) {

	if !list.withinRange(index) {
		return
	}

	list.data[index] = nil                                // cleanup reference
	copy(list.data[index:], list.data[index+1:list.size]) // shift to the left by one (slow operation, need ways to optimize this)
	list.size--

	list.shrink()
}

// Contains checks if elements (one or more) are present in the set.
// All elements have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *ArrayList) Contains(values ...interface{}) bool {

	for _, searchValue := range values {
		found := false
		for _, element := range list.data {
			if element == searchValue {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Values returns all elements in the list.
func (list *ArrayList) Values() []interface{} {
	newElements := make([]interface{}, list.size, list.size)
	copy(newElements, list.data[:list.size])
	return newElements
}

//IndexOf returns index of provided element
func (list *ArrayList) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.data {
		if element == value {
			return index
		}
	}
	return -1
}

// Empty returns true if list does not contain any elements.
func (list *ArrayList) Empty() bool {
	return list.size == 0
}

// Size returns number of elements within the list.
func (list *ArrayList) Size() int {
	return list.size
}

// Clear removes all elements from the list.
func (list *ArrayList) Clear() {
	list.size = 0
	list.data = []interface{}{}
}

// Sort sorts values (in-place) using.
func (list *ArrayList) Sort(comparator utils.Comparator) {
	if len(list.data) < 2 {
		return
	}
	utils.Sort(list.data[:list.size], comparator)
}

// Swap swaps the two values at the specified positions.
func (list *ArrayList) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.data[i], list.data[j] = list.data[j], list.data[i]
	}
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *ArrayList) Insert(index int, values ...interface{}) {

	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(values...)
		}
		return
	}

	l := len(values)
	list.growBy(l)
	list.size += l
	copy(list.data[index+l:], list.data[index:list.size-l])
	copy(list.data[index:], values)
}

// Set the value at specified index
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *ArrayList) Set(index int, value interface{}) {

	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(value)
		}
		return
	}

	list.data[index] = value
}

// String returns a string representation of container
func (list *ArrayList) String() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list.data[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (list *ArrayList) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

func (list *ArrayList) resize(cap int) {
	newElements := make([]interface{}, cap, cap)
	copy(newElements, list.data)
	list.data = newElements
}

// Expand the array if necessary, i.e. capacity will be reached if we add n elements
func (list *ArrayList) growBy(n int) {
	// When capacity is reached, grow by a factor of growthFactor and add number of elements
	currentCapacity := cap(list.data)
	if list.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

// Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity
func (list *ArrayList) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	// Shrink when size is at shrinkFactor * capacity
	currentCapacity := cap(list.data)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}
