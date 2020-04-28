package gsearch

func newPather(pos *Point, bits *Bitmap2D, paths []*Point, weight int) *Pather {
	p := &Pather{pos, bits, paths, weight}
	return p
}

// Pather 栈参数
type Pather struct {
	pos    *Point
	bits   *Bitmap2D
	path   []*Point
	weight int
}

// GetPathLen 获取路径长度
func (p *Pather) GetPathLen() int {
	return len(p.path)
}

// SimpleWeight 默认估价函数
func SimpleWeight(nparam *Pather, graph *Graph) int {
	// mbit := 2
	// switch {
	// case graph.blockflag < 0.5:
	// 	mbit = 8
	// case graph.blockflag < 1.0:
	// 	mbit = 4
	// default:
	// }
	// pw := len(nparam.paths) >> mbit

	pw := len(nparam.path)

	vx := nparam.pos.X - graph.end.X
	if vx < 0 {
		vx = -vx
	}
	vy := nparam.pos.Y - graph.end.Y
	if vy < 0 {
		vy = -vy
	}
	return -(vx + vy + pw)
}
