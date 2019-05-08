package array3

type Array3 struct {
	ysizes    []int
	xsizes    [][]int
	xyproduct int
	zsize     int
	ysize     int
	xsize     int
	data      [][][]interface{}

	cap int
}

func New() *Array3 {
	return NewWithCap(8, 8, 8)
}

func NewWithCap(zsize, ysize, xsize int) *Array3 {
	arr := &Array3{zsize: zsize, ysize: ysize, xsize: xsize}

	arr.ysizes = make([]int, arr.zsize, arr.zsize)

	arr.xsizes = make([][]int, arr.zsize, arr.zsize)
	for i := 0; i < arr.zsize; i++ {
		arr.xsizes[i] = make([]int, arr.ysize, arr.ysize)
	}

	arr.xyproduct = arr.ysize * arr.xsize
	arr.data = make([][][]interface{}, arr.zsize, arr.zsize)

	arr.cap = arr.zsize * arr.xyproduct
	return arr
}

func (arr *Array3) debugValues() []interface{} {
	var result []interface{}
	for _, z := range arr.data {
		if z != nil {
			for _, y := range z {
				if y == nil {
					for i := 0; i < arr.xsize; i++ {
						result = append(result, nil)
					}
				} else {
					for _, x := range y {
						if x == nil {
							result = append(result, struct{}{})
						} else {
							result = append(result, x)
						}
					}
				}
			}
		} else {
			for i := 0; i < arr.ysize*arr.xsize; i++ {
				result = append(result, nil)
			}
		}
	}
	return result
}

func (arr *Array3) Values() []interface{} {
	var result []interface{}
	for _, z := range arr.data {
		if z != nil {

			for _, y := range z {
				if y == nil {
					for i := 0; i < arr.xsize; i++ {
						result = append(result, nil)
					}
				} else {
					for _, x := range y {
						if x == nil {
							result = append(result, nil)
						} else {
							result = append(result, x)
						}
					}
				}
			}
		} else {
			for i := 0; i < arr.ysize*arr.xsize; i++ {
				result = append(result, nil)
			}
		}
	}

	return result
}

func (arr *Array3) Cap() int {
	return arr.cap
}

func (arr *Array3) Grow(size int) {
	zsize := arr.zsize + size
	temp := make([][][]interface{}, zsize, zsize)
	copy(temp, arr.data)
	arr.data = temp

	tempysizes := make([]int, zsize, zsize)
	copy(tempysizes, arr.ysizes)
	arr.ysizes = tempysizes

	tempxsizes := make([][]int, zsize, zsize)
	copy(tempxsizes, arr.xsizes)
	arr.xsizes = tempxsizes

	for i := arr.zsize; i < zsize; i++ {
		arr.xsizes[i] = make([]int, arr.ysize, arr.ysize)
	}

	arr.zsize += size
	arr.cap = arr.zsize * arr.xyproduct
}

func (arr *Array3) Set(idx int, value interface{}) {
	zindex := idx / arr.xyproduct
	nidx := (idx % arr.xyproduct)
	yindex := nidx / arr.xsize
	xindex := nidx % arr.xsize

	ydata := arr.data[zindex]
	if ydata == nil {
		ydata = make([][]interface{}, arr.ysize, arr.ysize)
		arr.data[zindex] = ydata
	}

	xdata := ydata[yindex]
	if xdata == nil {
		xdata = make([]interface{}, arr.xsize, arr.xsize)
		ydata[yindex] = xdata
		arr.ysizes[zindex]++
	}

	v := xdata[xindex]
	if v == nil {
		arr.xsizes[zindex][yindex]++
	}
	xdata[xindex] = value
}

func (arr *Array3) Get(idx int) (interface{}, bool) {
	zindex := idx / arr.xyproduct
	nextsize := (idx % arr.xyproduct)
	yindex := nextsize / arr.xsize
	xindex := nextsize % arr.xsize

	ydata := arr.data[zindex]
	if ydata == nil {
		return nil, false
	}

	xdata := ydata[yindex]
	if xdata == nil {
		return nil, false
	}

	v := xdata[xindex]
	return v, v != nil
}

func (arr *Array3) GetOrSet(idx int, DoSetValue func([]interface{}, int)) (result interface{}, isSet bool) {
	zindex := idx / arr.xyproduct
	nidx := (idx % arr.xyproduct)
	yindex := nidx / arr.xsize
	xindex := nidx % arr.xsize

	ydata := arr.data[zindex]
	if ydata == nil {
		ydata = make([][]interface{}, arr.ysize, arr.ysize)
		arr.data[zindex] = ydata
	}

	xdata := ydata[yindex]
	if xdata == nil {
		xdata = make([]interface{}, arr.xsize, arr.xsize)
		ydata[yindex] = xdata
		arr.ysizes[zindex]++
	}

	result = xdata[xindex]
	if result == nil {
		DoSetValue(xdata, xindex)
		result = xdata[xindex]
		if result == nil {
			panic("DoSetValue Not Set <nil> Value")
		}
		arr.xsizes[zindex][yindex]++
		return result, false
	}
	return result, true
}

func (arr *Array3) Del(idx int) (interface{}, bool) {
	zindex := idx / arr.xyproduct
	nextsize := (idx % arr.xyproduct)
	yindex := nextsize / arr.xsize
	xindex := nextsize % arr.xsize

	ydata := arr.data[zindex]
	if ydata == nil {
		return nil, false
	}

	xdata := ydata[yindex]
	if xdata == nil {
		return nil, false
	}

	v := xdata[xindex]
	xdata[xindex] = nil

	isnotnil := v != nil

	if isnotnil {
		arr.xsizes[zindex][yindex]--
		if arr.xsizes[zindex][yindex] == 0 {
			arr.data[zindex][yindex] = nil

			arr.ysizes[zindex]--
			if arr.ysizes[zindex] == 0 {
				arr.data[zindex] = nil
			}
		}
	}

	return v, isnotnil
}
