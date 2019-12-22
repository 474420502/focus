package pqueuekey

import (
	"github.com/474420502/focus/compare"
	"github.com/davecgh/go-spew/spew"
)

type PriorityQueue struct {
	queue *vbTree
}

func (pq *PriorityQueue) Iterator() *Iterator {
	return NewIterator(pq.queue.top)
}

func New(Compare compare.Compare) *PriorityQueue {
	return &PriorityQueue{queue: newVBT(Compare)}
}

func (pq *PriorityQueue) Size() int {
	return pq.queue.Size()
}

func (pq *PriorityQueue) Push(key, value interface{}) {
	pq.queue.Put(key, value)
}

func (pq *PriorityQueue) Top() (result interface{}, ok bool) {
	if pq.queue.top != nil {
		return pq.queue.top.value, true
	}
	return nil, false
}

func (pq *PriorityQueue) Pop() (result interface{}, ok bool) {
	if pq.queue.top != nil {
		result = pq.queue.top.value
		pq.queue.removeNode(pq.queue.top)
		return result, true
	}
	return nil, false
}

func (pq *PriorityQueue) Index(idx int) (interface{}, bool) {
	return pq.queue.Index(idx)
}

func (pq *PriorityQueue) IndexNode(idx int) (*Node, bool) {
	n := pq.queue.indexNode(idx)
	return n, n != nil
}

func (pq *PriorityQueue) Get(key interface{}) (interface{}, bool) {
	return pq.queue.Get(key)
}

func (pq *PriorityQueue) GetNode(key interface{}) (*Node, bool) {
	return pq.queue.GetNode(key)
}

func (pq *PriorityQueue) GetAround(key interface{}) [3]interface{} {
	return pq.queue.GetAround(key)
}

func (pq *PriorityQueue) GetAroundNode(key interface{}) [3]*Node {
	return pq.queue.getArounNode(key)
}

func (pq *PriorityQueue) GetRange(k1, k2 interface{}) []interface{} {
	return pq.queue.GetRange(k1, k2)
}

func (pq *PriorityQueue) RemoveIndex(idx int) (interface{}, bool) {
	return pq.queue.RemoveIndex(idx)
}

func (pq *PriorityQueue) Remove(key interface{}) (interface{}, bool) {
	return pq.queue.Remove(key)
}

func (pq *PriorityQueue) RemoveNode(node *Node) {
	pq.queue.removeNode(node)
}

func (pq *PriorityQueue) Values() []interface{} {
	return pq.queue.Values()
}

func (pq *PriorityQueue) String() string {
	return spew.Sprint(pq.queue.Values())
}
