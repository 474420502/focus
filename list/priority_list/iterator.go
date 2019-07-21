package plist

type Iterator struct {
	pl  *PriorityList
	cur *Node
}

func (iter *Iterator) Value() interface{} {
	return iter.cur.value
}

func (iter *Iterator) Prev() bool {
	if iter.cur == iter.pl.head {
		return false
	}
	iter.cur = iter.cur.prev
	return iter.cur != iter.pl.head
}

func (iter *Iterator) Next() bool {
	if iter.cur == iter.pl.tail {
		return false
	}
	iter.cur = iter.cur.next
	return iter.cur != iter.pl.tail
}

func (iter *Iterator) ToHead() {
	iter.cur = iter.pl.head
}

func (iter *Iterator) ToTail() {
	iter.cur = iter.pl.tail
}

type CircularIterator struct {
	pl  *PriorityList
	cur *Node
}

func (iter *CircularIterator) Value() interface{} {
	return iter.cur.value
}

func (iter *CircularIterator) Prev() bool {
	if iter.pl.size == 0 {
		return false
	}

	if iter.cur == iter.pl.head {
		iter.cur = iter.pl.tail.prev
		return true
	}

	iter.cur = iter.cur.prev
	if iter.cur == iter.pl.head {
		iter.cur = iter.pl.tail.prev
	}

	return true
}

func (iter *CircularIterator) Next() bool {
	if iter.pl.size == 0 {
		return false
	}

	if iter.cur == iter.pl.tail {
		iter.cur = iter.pl.head.next
		return true
	}

	iter.cur = iter.cur.next
	if iter.cur == iter.pl.tail {
		iter.cur = iter.pl.head.next
	}

	return true
}

func (iter *CircularIterator) ToHead() {
	iter.cur = iter.pl.head
}

func (iter *CircularIterator) ToTail() {
	iter.cur = iter.pl.tail
}
