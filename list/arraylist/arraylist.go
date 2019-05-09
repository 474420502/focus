package arraylist

type ArrayList struct {
	data []interface{}

	headIndex uint
	nextIndex uint

	size uint

	reserveHead  uint
	reserveLimit uint

	growSize   uint
	shrinkSize uint
}

func New() *ArrayList {
	al := &ArrayList{}
	al.reserveHead = 2
	al.reserveLimit = 256
	al.headIndex = al.reserveHead
	al.nextIndex = al.headIndex

	al.data = make([]interface{}, 8, 8)
	return al
}

func (l *ArrayList) Size() uint {
	return l.size
}

func (l *ArrayList) grow() {
	newsize := uint(len(l.data)) << 1

	l.reserveHead = al.reserveHead << 1

	l.headIndex = al.reserveHead
	l.nextIndex = al.headIndex

	l.data = make([]interface{}, 8, 8)
}

// Add add value to the tail of list
func (l *ArrayList) Add(v interface{}) {

	if l.nextIndex >= uint(len(l.data)) {
		l.grow()
	}

	l.size++
	l.data[l.nextIndex] = v

	// grow
}

// Push push is equal to  add
func (l *ArrayList) Push(v interface{}) {
	l.data = append(l.data, v)
}

func (l *ArrayList) Set(idx uint, value interface{}) {
	l.data[idx] = value
}

func (l *ArrayList) Get(idx uint) (result interface{}, isfound bool) {
	if idx >= l.Size() {
		return nil, false
	}
	return l.data[idx], true
}

func (l *ArrayList) Pop() (result interface{}, found bool) {
	if l.Size() == 0 {
		return nil, false
	}
	rindex := len(l.data) - 1
	result = l.data[rindex]
	l.data = l.data[0:rindex]
	return result, true
}

func (l *ArrayList) Remove(idx uint) (rvalue interface{}, isfound bool) {

	if idx >= l.Size() {
		return nil, false
	}

	rvalue = l.data[idx]
	l.data = append(l.data[0:idx], l.data[idx+1:])

	return rvalue, true
}

func (l *ArrayList) Values() (result []interface{}) {
	values := make([]interface{}, l.Size(), l.Size())
	copy(values, l.data)
	return values
}

func (l *ArrayList) Traversal(every func(index int, cur interface{}) bool) {
	for i, cur := range l.data {
		if !every(i, cur) {
			return
		}
	}
}
