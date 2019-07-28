package heap

import (
	"github.com/474420502/focus/compare"
)

type Heap struct {
	size     int
	elements []interface{}
	Compare  compare.Compare
}

func New(Compare compare.Compare) *Heap {
	h := &Heap{Compare: Compare}
	h.elements = make([]interface{}, 16, 16)
	return h
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Values() []interface{} {
	return h.elements[0:h.size]
}

func (h *Heap) grow() {
	ecap := len(h.elements)
	if h.size >= ecap {
		ecap = ecap << 1
		grow := make([]interface{}, ecap, ecap)
		copy(grow, h.elements)
		h.elements = grow
	}
}

func (h *Heap) Empty() bool {
	return h.size < 1
}

func (h *Heap) Clear() {
	h.size = 0
}

func (h *Heap) Reborn() {
	h.size = 0
	h.elements = make([]interface{}, 16, 16)
}

func (h *Heap) Top() (interface{}, bool) {
	if h.size != 0 {
		return h.elements[0], true
	}
	return nil, false
}

func (h *Heap) Put(v interface{}) {
	if v == nil {
		return
	}

	h.grow()

	curidx := h.size
	h.size++
	// up
	for curidx != 0 {
		pidx := (curidx - 1) >> 1
		pvalue := h.elements[pidx]
		if h.Compare(v, pvalue) > 0 {
			h.elements[curidx] = pvalue
			curidx = pidx
		} else {
			break
		}
	}
	h.elements[curidx] = v
}

func (h *Heap) slimming() {

	elen := len(h.elements)
	if elen >= 32 {
		ecap := elen >> 1
		if h.size <= ecap {
			ecap = elen - (ecap >> 1)
			slimming := make([]interface{}, ecap, ecap)
			copy(slimming, h.elements)
			h.elements = slimming
		}
	}

}

func (h *Heap) Pop() (interface{}, bool) {

	if h.size == 0 {
		return nil, false
	}

	curidx := 0
	top := h.elements[curidx]
	h.size--

	h.slimming()

	if h.size == 0 {
		return top, true
	}

	downvalue := h.elements[h.size]
	var cidx, c1, c2 int
	var cvalue1, cvalue2, cvalue interface{}
	// down
	for {
		cidx = curidx << 1

		c2 = cidx + 2
		if c2 < h.size {
			cvalue2 = h.elements[c2]

			c1 = cidx + 1
			cvalue1 = h.elements[c1]

			if h.Compare(cvalue1, cvalue2) >= 0 {
				cidx = c1
				cvalue = cvalue1
			} else {
				cidx = c2
				cvalue = cvalue2
			}
		} else {

			c1 = cidx + 1
			if c1 < h.size {
				cvalue1 = h.elements[c1]
				cidx = c1
				cvalue = cvalue1
			} else {
				break
			}

		}

		if h.Compare(cvalue, downvalue) > 0 {
			h.elements[curidx] = cvalue
			curidx = cidx
		} else {
			break
		}
	}
	h.elements[curidx] = downvalue
	return top, true
}
