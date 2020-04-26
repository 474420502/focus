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

	weight     func(nparam *Param) int
	weightHeap *heap.Tree

	srart Point
	end   Point

	dimX  int
	dimY  int
	tsize int // dimX * dimY
}

type Point struct {
	x, y  int
	msize int
}

func weightCompare(x1, x2 interface{}) int {
	p1, p2 := x1.(*Param), x2.(*Param)
	if p1.weight > p2.weight {
		return -1
	}
	return 1
}

// New 一个graph. 必须指定维度数据
func New(dx, dy int) *Graph {
	g := &Graph{}
	g.setDimension(dx, dy)
	g.tsize = g.dimX * g.dimY
	g.weightHeap = heap.New(weightCompare)
	g.weight = SimpleWeight
	return g
}

// setDimension 初始化维度
func (graph *Graph) setDimension(dx, dy int) {
	graph.dimX = dx
	graph.dimY = dy
}

// SetTarget 设置起点 结束点
func (graph *Graph) SetTarget(sx, sy, ex, ey int) {

	graph.srart.x = sx
	graph.srart.y = sy

	graph.end.x = ex
	graph.end.y = ey

	gdata := make([]byte, graph.tsize)
	paths := make([]Point, 0)

	param := newParam(graph.srart, gdata, paths, 0)

	graph.weightHeap.Put(param)
}

// Search 执行搜索
func (graph *Graph) Search() {
	for !graph.weightHeap.Empty() {
		param, _ := graph.weightHeap.Pop()
		graph.Traversing(param.(*Param))
	}
}

func (graph *Graph) debugShow(param *Param) {
	param.cur.msize = param.cur.y*graph.dimX + param.cur.x
	param.graph[param.cur.msize] |= 0b01000000
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
func (graph *Graph) Traversing(param *Param) {

	if param.cur.x == graph.end.x && param.cur.y == graph.end.y {

		graph.debugShow(param)

		log.Println("finish")
		return
	}

	param.paths = append(param.paths, param.cur)
	param.cur.msize = param.cur.y*graph.dimX + param.cur.x
	param.graph[param.cur.msize] |= 0b01000000

	graph.debugShow(param)

	graph.left(param)
	graph.right(param)
	graph.up(param)
	graph.down(param)
}

// SetWeight 设置估价函数
func (graph *Graph) SetWeight(weight func(nparam *Param) int) {
	graph.weight = weight
}

func (graph *Graph) evaluate(nparam *Param, param *Param) {
	nparam.graph = make([]byte, graph.tsize)
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
	if param.graph[nparam.cur.y*graph.dimX+nparam.cur.x]&0b01000000 > 0 {
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
	if param.graph[nparam.cur.y*graph.dimX+nparam.cur.x]&0b01000000 > 0 {
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
	if param.graph[nparam.cur.y*graph.dimX+nparam.cur.x]&0b01000000 > 0 {
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
	if param.graph[nparam.cur.y*graph.dimX+nparam.cur.x]&0b01000000 > 0 {
		return
	}

	graph.evaluate(nparam, param)
}
