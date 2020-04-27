package astar

func newParam(cur Point, bits *Bitmap2D, paths []Point, weight int) *Param {
	p := &Param{cur, bits, paths, weight}
	return p
}

// Param 栈参数
type Param struct {
	cur    Point
	bits   *Bitmap2D
	paths  []Point
	weight int
}

// SimpleWeight 默认估价函数
func SimpleWeight(nparam *Param, graph *Graph) int {
	mbit := 2
	switch {
	case graph.blockflag < 0.5:
		mbit = 8
	case graph.blockflag < 1.0:
		mbit = 4
	default:
	}
	pw := len(nparam.paths) >> mbit
	vx := nparam.cur.x - graph.end.x
	vy := nparam.cur.y - graph.end.y
	w := -((vx*vx + vy*vy) + pw*pw)

	return w
}
