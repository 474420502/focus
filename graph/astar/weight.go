package astar

import "math"

func newParam(pos *Point, bits *Bitmap2D, paths []*Point, weight int) *Param {
	p := &Param{pos, bits, paths, weight}
	return p
}

// Param 栈参数
type Param struct {
	pos    *Point
	bits   *Bitmap2D
	paths  []*Point
	weight int
}

// SimpleWeight 默认估价函数
func SimpleWeight(nparam *Param, graph *Graph) int {
	// mbit := 2
	// switch {
	// case graph.blockflag < 0.5:
	// 	mbit = 8
	// case graph.blockflag < 1.0:
	// 	mbit = 4
	// default:
	// }
	// pw := len(nparam.paths) >> mbit

	pw := len(nparam.paths)
	vx := int(math.Abs(float64(nparam.pos.X - graph.end.X)))
	vy := int(math.Abs(float64(nparam.pos.Y - graph.end.Y)))
	w := -(vx + vy + pw)

	return w
}
