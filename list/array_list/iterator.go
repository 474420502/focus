package arraylist

type Iterator struct {
	al     *ArrayList
	cur    uint
	isInit bool
}

func (iter *Iterator) Value() interface{} {
	v, _ := iter.al.Index(iter.cur)
	return v
}

func (iter *Iterator) Prev() bool {

	if iter.isInit == false {
		if iter.al.size != 0 {
			iter.isInit = true
			iter.cur = iter.al.size - 1
			return true
		}
		return false
	}

	if iter.cur <= 0 {
		return false
	}
	iter.cur--
	return true
}

func (iter *Iterator) Next() bool {

	if iter.isInit == false {
		if iter.al.size != 0 {
			iter.isInit = true
			iter.cur = 0
			return true
		}
		return false
	}

	if iter.cur >= iter.al.size-1 {
		return false
	}
	iter.cur++
	return true
}

func (iter *Iterator) ToHead() {
	iter.isInit = true
	iter.cur = 0
}

func (iter *Iterator) ToTail() {
	iter.isInit = true
	iter.cur = iter.al.size - 1
}

type CircularIterator struct {
	al     *ArrayList
	cur    uint
	isInit bool
}

func (iter *CircularIterator) Value() interface{} {
	v, _ := iter.al.Index(iter.cur)
	return v
}

func (iter *CircularIterator) Prev() bool {

	if iter.isInit == false {
		if iter.al.size != 0 {
			iter.isInit = true
			iter.cur = iter.al.size - 1
			return true
		}
		return false
	}

	if iter.al.size == 0 {
		return false
	}

	if iter.cur <= 0 {
		iter.cur = iter.al.size - 1
	} else {
		iter.cur--
	}
	return true
}

func (iter *CircularIterator) Next() bool {

	if iter.isInit == false {
		if iter.al.size != 0 {
			iter.isInit = true
			iter.cur = 0
			return true
		}
		return false
	}

	if iter.al.size == 0 {
		return false
	}

	if iter.cur >= iter.al.size-1 {
		iter.cur = 0
	} else {
		iter.cur++
	}
	return true
}

func (iter *CircularIterator) ToHead() {
	iter.isInit = true
	iter.cur = 0
}

func (iter *CircularIterator) ToTail() {
	iter.isInit = true
	iter.cur = iter.al.size - 1
}
