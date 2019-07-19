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
	//growthFactor = float32(2.0)  // growth by 100%
	//shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

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
		return
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

	lcap := uint(len(l.data))
	nSize := lcap << 1
	temp := make([]interface{}, nSize, nSize)

	ghidx := lcap / 2
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
	for l.tidx+psize > uint(len(l.data)) { // [0 1 2 3 4 5 6]
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

	if idx >= l.size {
		return nil, false
	}

	offset := l.hidx + 1 + idx

	isfound = true
	result = l.data[offset]
	// l.data[offset] = nil // cleanup reference

	if l.size-l.tidx > l.hidx {
		copy(l.data[offset:], l.data[offset+1:l.tidx]) // shift to the left by one (slow operation, need ways to optimize this)
		l.tidx--
	} else {
		copy(l.data[l.hidx+2:], l.data[l.hidx+1:offset])
		l.hidx++
	}

	l.size--
	l.shrink()

	return
}

func (l *ArrayList) Contains(values ...interface{}) bool {

	for _, searchValue := range values {
		found := false
		for _, element := range l.data[l.hidx+1 : l.tidx] {
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

func (l *ArrayList) Values() []interface{} {
	newElements := make([]interface{}, l.size, l.size)
	copy(newElements, l.data[l.hidx+1:l.tidx])
	return newElements
}
