package heap

import (
	"github.com/474420502/focus/compare"
)

type Tree struct {
	size     int
	elements []interface{}
	Compare  compare.Compare
}

func New(Compare compare.Compare) *Tree {
	h := &Tree{Compare: Compare}
	h.elements = make([]interface{}, 16, 16)
	return h
}

func (h *Tree) Size() int {
	return h.size
}

func (h *Tree) Values() []interface{} {
	return h.elements[0:h.size]
}

func (h *Tree) grow() {
	ecap := len(h.elements)
	if h.size >= ecap {
		ecap = ecap << 1
		grow := make([]interface{}, ecap, ecap)
		copy(grow, h.elements)
		h.elements = grow
	}
}

func (h *Tree) Empty() bool {
	return h.size < 1
}

func (h *Tree) Clear() {
	h.size = 0
}

func (h *Tree) Reborn() {
	h.size = 0
	h.elements = make([]interface{}, 16, 16)
}

func (h *Tree) Top() (interface{}, bool) {
	if h.size != 0 {
		return h.elements[0], true
	}
	return nil, false
}

func (h *Tree) Put(v interface{}) {
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

func (h *Tree) slimming() {

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

func (h *Tree) Pop() (interface{}, bool) {

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
	// down
	for {
		cidx = curidx << 1

		c1 = cidx + 1
		c2 = cidx + 2

		if c2 < h.size {
			if h.Compare(h.elements[c1], h.elements[c2]) >= 0 {
				cidx = c1
			} else {
				cidx = c2
			}
		} else {
			cidx = c1
			if c1 >= h.size {
				break
			}
		}

		if h.Compare(h.elements[cidx], downvalue) > 0 {
			h.elements[curidx] = h.elements[cidx]
			curidx = cidx
		} else {
			break
		}
	}
	h.elements[curidx] = downvalue
	return top, true
}
