package arrayn

type SizeN struct {
	Sizes []int
}

type ProductN struct {
	Values []int
}

type DimensionSize struct {
	Sizes []int
}

type Node struct {
	size int
	data interface{}
}

type ArrayN struct {
	dims    []int
	product []int

	dimN int
	data *Node // []*Node

	cap int
}

func New() *ArrayN {
	return NewWithCap(8, 8, 8)
}

func NewWithCap(dims ...int) *ArrayN {
	arr := &ArrayN{dimN: len(dims), dims: dims}
	arr.product = make([]int, len(dims)-1, len(dims)-1)
	for i := 0; i < len(dims)-1; i++ {
		pvalue := 1
		for n := i + 1; n < len(dims); n++ {
			pvalue *= dims[n]
		}
		arr.product[i] = pvalue
	}
	// arr.data = make([]*Node, arr.dims[0], arr.dims[0])
	arr.cap = 1
	for _, d := range arr.dims {
		arr.cap *= d
	}
	return arr
}

func (arr *ArrayN) collectValues(curDim int, cur *Node, result *[]interface{}) {
	if cur == nil {
		return
	}

	if curDim == 1 {
		for _, v := range cur.data.([]interface{}) {
			if v != nil {
				*result = append(*result, v)
			} else {
				*result = append(*result, nil)
			}
		}
		return
	}

	for _, n := range cur.data.([]*Node) {
		if n != nil {
			arr.collectValues(curDim-1, n, result)
		} else {
			total := 1
			for i := len(arr.dims) - curDim + 1; i < len(arr.dims); i++ {
				total *= arr.dims[i]
			}

			for i := 0; i < total; i++ {
				*result = append(*result, nil)
			}
		}
	}
}

func (arr *ArrayN) Values() (result []interface{}) {
	arr.collectValues(arr.dimN, arr.data, &result)
	return
}

func (arr *ArrayN) set(curDim int, curidx int, pdata **Node, parent *Node) (*Node, int) {

	sidx := arr.dimN - curDim

	if *pdata == nil {
		if parent != nil {
			parent.size++
		}
		if curDim > 1 {
			*pdata = &Node{data: make([]*Node, arr.dims[sidx], arr.dims[sidx])}
		} else {
			*pdata = &Node{data: make([]interface{}, arr.dims[sidx], arr.dims[sidx])}
			return *pdata, curidx
		}
	}

	cur := *pdata
	if curDim == 1 {
		return cur, curidx
	}

	nidx := curidx % arr.product[sidx]
	dimindex := curidx / arr.product[sidx]

	return arr.set(curDim-1, nidx, &cur.data.([]*Node)[dimindex], cur)
}

func (arr *ArrayN) Cap() int {
	return arr.cap
}

func (arr *ArrayN) Grow(size int) {
	arr.dims[0] += size

	pvalue := 1
	for n := 1; n < len(arr.dims); n++ {
		pvalue *= arr.dims[n]
	}
	arr.product[0] = pvalue

	tempdata := arr.data.data.([]*Node)

	newData := make([]*Node, arr.dims[0], arr.dims[0])
	copy(newData, tempdata)
	arr.data.data = newData

	arr.cap = 1
	for _, d := range arr.dims {
		arr.cap *= d
	}
}

func (arr *ArrayN) Set(idx int, value interface{}) {
	n, nidx := arr.set(arr.dimN, idx, &arr.data, nil)
	n.data.([]interface{})[nidx] = value
}

func (arr *ArrayN) get(curDim int, curidx int, pdata **Node) (*Node, int) {
	sidx := arr.dimN - curDim

	if *pdata == nil {
		return nil, 0
	}

	cur := *pdata
	if curDim == 1 {
		return cur, curidx
	}

	nidx := curidx % arr.product[sidx]
	dimindex := curidx / arr.product[sidx]
	return arr.get(curDim-1, nidx, &cur.data.([]*Node)[dimindex])
}

func (arr *ArrayN) Get(idx int) (interface{}, bool) {
	n, nidx := arr.get(arr.dimN, idx, &arr.data)
	if n != nil {
		v := n.data.([]interface{})[nidx]
		return v, v != nil
	}
	return nil, false
}

func (arr *ArrayN) del(curDim int, curidx int, pdata **Node) (interface{}, bool) {

	cur := *pdata
	if cur == nil {
		return nil, false
	}

	if curDim == 1 {
		values := cur.data.([]interface{})

		v := values[curidx]
		if v != nil {
			cur.size--
			values[curidx] = nil
			if cur.size == 0 {
				return v, true
			}
		}
		return v, false
	}

	sidx := arr.dimN - curDim

	nidx := curidx % arr.product[sidx]
	dimindex := curidx / arr.product[sidx]
	curdata := cur.data.([]*Node)

	v, ok := arr.del(curDim-1, nidx, &curdata[dimindex])
	if ok {
		cur.size--
		curdata[dimindex] = nil
		if cur.size == 0 {
			return v, true
		}
	}

	return v, false
}

func (arr *ArrayN) Del(idx int) (interface{}, bool) {
	v, _ := arr.del(arr.dimN, idx, &arr.data)
	if v != nil {
		return v, true
	}
	return nil, false
}
