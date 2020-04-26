package astar

import (
	"fmt"
	"log"

	"github.com/474420502/focus/tree/heap"
)

// Graph 图
type Graph struct {
	// direction   []*Graph
	// flag        int
	isDebug bool

	weight     func(nparam *Param) int
	weightHeap *heap.Tree

	step int

	infoMap []byte

	srart Point
	end   Point

	dimX  int
	dimY  int
	tsize int // dimX * dimY
	bsize int // bit size
}

// Point 点
type Point struct {
	x, y  int
	msize int
}

func weightCompare(x1, x2 interface{}) int {
	p1, p2 := x1.(*Param), x2.(*Param)
	if p1.weight > p2.weight {
		return 1
	}
	return -1
}

// New2D 一个graph. 必须指定维度数据
func New2D(dx, dy int) *Graph {
	g := &Graph{}
	g.setDimension(dx, dy)
	g.tsize = g.dimX * g.dimY
	g.bsize = (g.tsize + 1) / 8
	g.weightHeap = heap.New(weightCompare)
	g.weight = SimpleWeight
	g.infoMap = make([]byte, g.tsize)
	return g
}

// setDimension 初始化维度
func (graph *Graph) setDimension(dx, dy int) {
	graph.dimX = dx
	graph.dimY = dy
}

// SetBlock 设置起点 结束点
func (graph *Graph) SetBlock(x, y int) {

}

// SetTarget 设置起点 结束点
func (graph *Graph) SetTarget(sx, sy, ex, ey int) {
	graph.srart.x = sx
	graph.srart.y = sy

	graph.srart.msize = graph.srart.y*graph.dimX + graph.srart.x

	graph.end.x = ex
	graph.end.y = ey

	graph.end.msize = graph.end.y*graph.dimX + graph.end.x

	gdata := make([]byte, graph.bsize)
	paths := make([]Point, 0)

	param := newParam(graph.srart, gdata, paths, 0)
	graph.weightHeap.Put(param)
}

// GetStep 执行的步数
func (graph *Graph) GetStep() int {
	return graph.step
}

// Search 执行搜索
func (graph *Graph) Search() {
	graph.step = 0
	for !graph.weightHeap.Empty() {
		param, _ := graph.weightHeap.Pop()
		if graph.Traversing(param.(*Param)) {
			break
		}
		graph.step++
	}
}

func (graph *Graph) debugShow(param *Param) {
	param.cur.msize = param.cur.y*graph.dimX + param.cur.x
	param.graph[param.cur.msize] |= 0b10000000
	content := "\n"
	for y := 0; y < graph.dimY; y++ {
		for x := 0; x < graph.dimX; x++ {
			showmsize := y*graph.dimX + x
			content += fmt.Sprintf("%03d ", param.graph[showmsize])
		}
		content += "\n"
	}
	log.Println(content)
}

// Traversing 遍历结果
func (graph *Graph) Traversing(param *Param) bool {

	if param.cur.x == graph.end.x && param.cur.y == graph.end.y {
		if graph.isDebug {
			graph.debugShow(param)
		}
		log.Println("finish")

		return true
	}

	param.paths = append(param.paths, param.cur)
	// param.cur.msize = param.cur.y*graph.dimX + param.cur.x
	param.graph[param.cur.msize] |= 0b10000000

	graph.left(param)
	graph.right(param)
	graph.up(param)
	graph.down(param)

	return false
}

// SetWeight 设置估价函数
func (graph *Graph) SetWeight(weight func(nparam *Param) int) {
	graph.weight = weight
}

func (graph *Graph) evaluate(nparam *Param, param *Param) {
	nparam.graph = make([]byte, graph.bsize)
	copy(nparam.graph, param.graph)

	nparam.paths = make([]Point, len(param.paths))
	copy(nparam.paths, param.paths)

	nparam.weight = graph.weight(nparam)
	graph.weightHeap.Put(nparam)
}

func (graph *Graph) left(param *Param) {
	leftx := param.cur.x - 1
	if leftx < 0 {
		return
	}

	nparam := &Param{cur: Point{x: leftx, y: param.cur.y}}
	nparam.cur.msize = nparam.cur.y*graph.dimX + nparam.cur.x

	nb := nparam.cur.msize / 8
	mb := nparam.cur.msize % 8

	pinfo := graph.infoMap[nparam.cur.msize]
	if pinfo&0b10000000 > 0 {
		return
	}

	graph.evaluate(nparam, param)
}

func (graph *Graph) right(param *Param) {
	rightx := param.cur.x + 1
	if rightx >= graph.dimX {
		return
	}

	nparam := &Param{cur: Point{x: rightx, y: param.cur.y}}
	pinfo := param.graph[nparam.cur.y*graph.dimX+nparam.cur.x]
	if pinfo&0b11000000 > 0 {
		return
	}

	graph.evaluate(nparam, param)
}

func (graph *Graph) up(param *Param) {
	upy := param.cur.y + 1
	if upy >= graph.dimY {
		return
	}

	nparam := &Param{cur: Point{x: param.cur.x, y: upy}}
	pinfo := param.graph[nparam.cur.y*graph.dimX+nparam.cur.x]
	if pinfo&0b11000000 > 0 {
		return
	}

	graph.evaluate(nparam, param)
}

func (graph *Graph) down(param *Param) {
	downy := param.cur.y - 1
	if downy < 0 {
		return
	}

	nparam := &Param{cur: Point{x: param.cur.x, y: downy}}
	pinfo := param.graph[nparam.cur.y*graph.dimX+nparam.cur.x]

	if pinfo&0b11000000 > 0 {
		return
	}

	graph.evaluate(nparam, param)
}
