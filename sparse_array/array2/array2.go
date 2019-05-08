package array2

type Array2 struct {
	sizes []int
	ysize int
	xsize int
	data  [][]interface{}

	cap int
}

func New() *Array2 {
	return NewWithCap(8, 8)
}

func NewWithCap(ysize, xsize int) *Array2 {
	arr := &Array2{ysize: ysize, xsize: xsize}
	arr.sizes = make([]int, arr.ysize, arr.ysize)
	arr.data = make([][]interface{}, arr.ysize, arr.ysize)

	arr.cap = arr.ysize * arr.xsize
	return arr
}

func (arr *Array2) debugValues() []interface{} {
	var result []interface{}
	for _, y := range arr.data {
		if y != nil {
			for _, v := range y {
				if v == nil {
					result = append(result, struct{}{})
				} else {
					result = append(result, v)
				}
			}
		} else {
			for i := 0; i < arr.xsize; i++ {
				result = append(result, nil)
			}
		}
	}
	return result
}

func (arr *Array2) Values() []interface{} {
	var result []interface{}
	for _, y := range arr.data {
		if y != nil {
			for _, v := range y {
				if v == nil {
					result = append(result, nil)
				} else {
					result = append(result, v)
				}
			}
		} else {
			for i := 0; i < arr.xsize; i++ {
				result = append(result, nil)
			}
		}
	}
	return result
}

func (arr *Array2) Cap() int {
	return arr.cap
}

func (arr *Array2) Grow(size int) {
	arr.ysize += size
	temp := make([][]interface{}, arr.ysize, arr.ysize)
	copy(temp, arr.data)
	arr.data = temp

	tempsizes := make([]int, arr.ysize, arr.ysize)
	copy(tempsizes, arr.sizes)
	arr.sizes = tempsizes

	arr.cap = arr.ysize * arr.xsize
}

func (arr *Array2) Set(idx int, value interface{}) {
	yindex := idx / arr.xsize
	xindex := idx % arr.xsize

	xdata := arr.data[yindex]
	if xdata == nil {
		xdata = make([]interface{}, arr.xsize, arr.xsize)
		arr.data[yindex] = xdata
	}

	if xdata[xindex] == nil {
		arr.sizes[yindex]++
	}

	xdata[xindex] = value
}

func (arr *Array2) Get(idx int) (interface{}, bool) {
	yindex := idx / arr.xsize
	xindex := idx % arr.xsize

	xdata := arr.data[yindex]
	if xdata == nil {
		return nil, false
	}
	v := xdata[xindex]
	return v, v != nil
}

func (arr *Array2) GetOrSet(idx int, DoSetValue func([]interface{}, int)) (result interface{}, isSet bool) {
	yindex := idx / arr.xsize
	xindex := idx % arr.xsize

	xdata := arr.data[yindex]
	if xdata == nil {
		xdata = make([]interface{}, arr.xsize, arr.xsize)
		arr.data[yindex] = xdata
	}

	result = xdata[xindex]
	if result == nil {
		DoSetValue(xdata, xindex)
		result = xdata[xindex]
		if result == nil {
			panic("DoSetValue Not Set <nil> Value")
		}
		arr.sizes[yindex]++
		return result, true
	}
	return result, false
}

func (arr *Array2) Del(idx int) (interface{}, bool) {
	yindex := idx / arr.xsize
	xindex := idx % arr.xsize

	xdata := arr.data[yindex]
	if xdata == nil {
		return nil, false
	}
	v := xdata[xindex]
	xdata[xindex] = nil

	isnil := v != nil
	if isnil {
		arr.sizes[yindex]--
		if arr.sizes[yindex] == 0 {
			arr.data[yindex] = nil
		}
	}
	return v, isnil
}
