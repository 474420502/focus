package arraylist

import "log"

type ArrayList struct {
	data []interface{}
	hidx uint // [ nil(hdix) 1 nil(tidx) ]
	tidx uint
	size uint

	growthSize uint
	shrinkSize uint
}

const (
	listMaxLimit = uint(1) << 63
	listMinLimit = uint(8)
	initCap      = uint(8)
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// New instantiates a new list and adds the passed values, if any, to the list
func New() *ArrayList {
	l := &ArrayList{}
	l.data = make([]interface{}, initCap, initCap)
	l.tidx = initCap / 2
	l.hidx = l.tidx - 1
	return l
}

func (l *ArrayList) Clear() {
	l.data = make([]interface{}, 8, 8)
	l.tidx = initCap / 2
	l.hidx = l.tidx - 1
	l.size = 0
}

func (l *ArrayList) Empty() bool {
	return l.size == 0
}

func (l *ArrayList) Size() uint {
	return l.size
}

func (l *ArrayList) shrink() {

	if l.size <= listMinLimit {
		log.Panic("list size is over listMaxLimit", listMinLimit)
	}

	if l.size <= l.shrinkSize {
		nSize := l.shrinkSize - l.shrinkSize>>1
		temp := make([]interface{}, nSize, nSize)

		ghidx := l.size / 2
		gtidx := ghidx + l.size + 1
		copy(temp[ghidx+1:], l.data[l.hidx+1:l.tidx])
		l.data = temp
		l.hidx = ghidx
		l.tidx = gtidx
	}

}

// 后续需要优化 growth 策略
func (l *ArrayList) growth() {

	if l.size >= listMaxLimit {
		log.Panic("list size is over listMaxLimit", listMaxLimit)
	}

	nSize := l.size << 1
	temp := make([]interface{}, nSize, nSize)

	ghidx := l.size / 2
	gtidx := ghidx + l.size + 1
	copy(temp[ghidx+1:], l.data[l.hidx+1:l.tidx])
	l.data = temp
	l.hidx = ghidx
	l.tidx = gtidx

}

func (l *ArrayList) PushFront(values ...interface{}) {
	psize := uint(len(values))
	for l.hidx+1-psize > listMaxLimit {
		l.growth()
		// panic("growth -1")
	}

	for _, v := range values {
		l.data[l.hidx] = v
		l.hidx--
	}
	l.size += psize
}

func (l *ArrayList) PushBack(values ...interface{}) {
	psize := uint(len(values))
	for l.tidx+psize > uint(len(l.data)) {
		l.growth()
	}

	for _, v := range values {
		l.data[l.tidx] = v
		l.tidx++
	}
	l.size += psize
}

func (l *ArrayList) PopFront() (result interface{}, found bool) {
	if l.size != 0 {
		l.size--
		l.hidx++
		return l.data[l.hidx], true
	}
	return nil, false
}

func (l *ArrayList) PopBack() (result interface{}, found bool) {
	if l.size != 0 {
		l.size--
		l.tidx--
		return l.data[l.tidx], true
	}
	return nil, false
}

func (l *ArrayList) Index(idx uint) (interface{}, bool) {
	if idx < l.size {
		return l.data[idx+l.hidx+1], true
	}
	return nil, false
}

func (l *ArrayList) Remove(idx uint) (result interface{}, isfound bool) {

	if idx < l.size {
		return nil, false
	}

	offset := l.hidx + 1 + idx

	isfound = true
	result = l.data[offset]
	l.data[offset] = nil // cleanup reference

	copy(l.data[offset:], l.data[idx+1:l.size]) // shift to the left by one (slow operation, need ways to optimize this)
	l.size--
	l.shrink()

	return
}

// func (list *ArrayList) Contains(values ...interface{}) bool {

// 	for _, searchValue := range values {
// 		found := false
// 		for _, element := range list.data {
// 			if element == searchValue {
// 				found = true
// 				break
// 			}
// 		}
// 		if !found {
// 			return false
// 		}
// 	}
// 	return true
// }

func (l *ArrayList) Values() []interface{} {
	newElements := make([]interface{}, l.size, l.size)
	copy(newElements, l.data[l.hidx+1:l.tidx])
	return newElements
}

// func (list *ArrayList) IndexOf(value interface{}) int {
// 	if list.size == 0 {
// 		return -1
// 	}
// 	for index, element := range list.data {
// 		if element == value {
// 			return index
// 		}
// 	}
// 	return -1
// }

// func (list *ArrayList) Empty() bool {
// 	return list.size == 0
// }

// func (list *ArrayList) Size() int {
// 	return list.size
// }

// // Clear removes all elements from the list.
// func (list *ArrayList) Clear() {
// 	list.size = 0
// 	list.data = []interface{}{}
// }

// // Sort sorts values (in-place) using.
// func (list *ArrayList) Sort(comparator utils.Comparator) {
// 	if len(list.data) < 2 {
// 		return
// 	}
// 	utils.Sort(list.data[:list.size], comparator)
// }

// // Swap swaps the two values at the specified positions.
// func (list *ArrayList) Swap(i, j int) {
// 	if list.withinRange(i) && list.withinRange(j) {
// 		list.data[i], list.data[j] = list.data[j], list.data[i]
// 	}
// }

// // Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// // Does not do anything if position is negative or bigger than list's size
// // Note: position equal to list's size is valid, i.e. append.
// func (list *ArrayList) Insert(index int, values ...interface{}) {

// 	if !list.withinRange(index) {
// 		// Append
// 		if index == list.size {
// 			list.Add(values...)
// 		}
// 		return
// 	}

// 	l := len(values)
// 	list.growBy(l)
// 	list.size += l
// 	copy(list.data[index+l:], list.data[index:list.size-l])
// 	copy(list.data[index:], values)
// }

// // Set the value at specified index
// // Does not do anything if position is negative or bigger than list's size
// // Note: position equal to list's size is valid, i.e. append.
// func (list *ArrayList) Set(index int, value interface{}) {

// 	if !list.withinRange(index) {
// 		// Append
// 		if index == list.size {
// 			list.Add(value)
// 		}
// 		return
// 	}

// 	list.data[index] = value
// }

// // String returns a string representation of container
// func (list *ArrayList) String() string {
// 	str := "ArrayList\n"
// 	values := []string{}
// 	for _, value := range list.data[:list.size] {
// 		values = append(values, fmt.Sprintf("%v", value))
// 	}
// 	str += strings.Join(values, ", ")
// 	return str
// }

// // Check that the index is within bounds of the list
// func (list *ArrayList) withinRange(index int) bool {
// 	return index >= 0 && index < list.size
// }

// func (list *ArrayList) resize(cap int) {
// 	newElements := make([]interface{}, cap, cap)
// 	copy(newElements, list.data)
// 	list.data = newElements
// }

// // Expand the array if necessary, i.e. capacity will be reached if we add n elements
// func (list *ArrayList) growBy(n int) {
// 	// When capacity is reached, grow by a factor of growthFactor and add number of elements
// 	currentCapacity := cap(list.data)
// 	if list.size+n >= currentCapacity {
// 		newCapacity := int(growthFactor * float32(currentCapacity+n))
// 		list.resize(newCapacity)
// 	}
// }

// // Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity
// func (list *ArrayList) shrink() {
// 	if shrinkFactor == 0.0 {
// 		return
// 	}
// 	// Shrink when size is at shrinkFactor * capacity
// 	currentCapacity := cap(list.data)
// 	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
// 		list.resize(list.size)
// 	}
// }
