package astar

func newParam(cur Point, graph []byte, paths []Point, weight int) *Param {
	p := &Param{cur, graph, paths, weight}
	return p
}

// Param 栈参数
type Param struct {
	cur    Point
	graph  []byte
	paths  []Point
	weight int
}

// SimpleWeight 默认估价函数
func SimpleWeight(nparam *Param) int {
	pw := len(nparam.paths)
	return -(nparam.cur.x*nparam.cur.x + nparam.cur.y*nparam.cur.y + pw)
}

// // WeightHeap 权重堆
// type WeightHeap []*Param

// func (ph *WeightHeap) Len() int { return len(*ph) }

// func (ph *WeightHeap) Less(i, j int) bool {
// 	return ph[i].weight > ph[j].weight
// }

// func (ph *WeightHeap) Swap(i, j int) {
// 	(*ph)[i], (*ph)[j] = (*ph)[j], ph[i]
// }

// func (wh *WeightHeap) Push(v interface{}) {
// 	*wh = append((*wh), v.(*Param))
// }

// func (wh *WeightHeap) Pop() interface{} {
// 	n := len(wh)
// 	v := (*wh)[n-1]
// 	*wh = (*wh)[0 : n-1]
// 	return v
// }
